package variables

import (
	"github.com/caarlos0/env"

	"github.com/devtron-labs/devtron/pkg/sql"
	"github.com/devtron-labs/devtron/pkg/variables/cache"
	"github.com/devtron-labs/devtron/pkg/variables/helper"
	"github.com/devtron-labs/devtron/pkg/variables/models"
	repository2 "github.com/devtron-labs/devtron/pkg/variables/repository"
	"github.com/devtron-labs/devtron/pkg/variables/utils"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"sync"
	"time"
)

type ScopedVariableService interface {
	CreateVariables(payload models.Payload) error
	GetScopedVariables(scope models.Scope, varNames []string, maskSensitiveData bool) (scopedVariableDataObj []*models.ScopedVariableData, err error)
	GetJsonForVariables() (*models.Payload, error)
}

type ScopedVariableServiceImpl struct {
	logger                   *zap.SugaredLogger
	scopedVariableRepository repository2.ScopedVariableRepository
	VariableNameConfig       *VariableConfig
	VariableCache            *cache.VariableCacheObj
}

func NewScopedVariableServiceImpl(logger *zap.SugaredLogger, scopedVariableRepository repository2.ScopedVariableRepository) (*ScopedVariableServiceImpl, error) {
	scopedVariableService := &ScopedVariableServiceImpl{
		logger:                   logger,
		scopedVariableRepository: scopedVariableRepository,
		VariableCache:            &cache.VariableCacheObj{CacheLock: &sync.Mutex{}},
	}
	cfg, err := GetVariableNameConfig()
	if err != nil {
		return nil, err
	}
	loadVariableCache(cfg, scopedVariableService)
	scopedVariableService.VariableNameConfig = cfg
	return scopedVariableService, nil
}

type VariableConfig struct {
	VariableNameRegex    string `env:"SCOPED_VARIABLE_NAME_REGEX" envDefault:"^[a-zA-Z][a-zA-Z0-9_-]{0,62}[a-zA-Z0-9]$"`
	VariableCacheEnabled bool   `env:"VARIABLE_CACHE_ENABLED" envDefault:"true"`
}

func loadVariableCache(cfg *VariableConfig, service *ScopedVariableServiceImpl) {
	if cfg.VariableCacheEnabled {
		go service.loadVarCache()
	}
}
func GetVariableNameConfig() (*VariableConfig, error) {
	cfg := &VariableConfig{}
	err := env.Parse(cfg)
	return cfg, err
}
func (impl *ScopedVariableServiceImpl) SetVariableCache(cache *cache.VariableCacheObj) {
	impl.VariableCache = cache
}

func (impl *ScopedVariableServiceImpl) loadVarCache() {
	variableCache := impl.VariableCache
	variableCache.ResetCache()
	variableCache.TakeLock()
	defer variableCache.ReleaseLock()
	variableMetadata, err := impl.scopedVariableRepository.GetAllVariableMetadata()
	if err != nil {
		impl.logger.Errorw("error occurred while fetching variable metadata", "err", err)
		return
	}
	variableCache.SetData(variableMetadata)
	impl.logger.Info("variable cache loaded successfully")
}

func (impl *ScopedVariableServiceImpl) CreateVariables(payload models.Payload) error {
	err, _ := impl.isValidPayload(payload)
	if err != nil {
		impl.logger.Errorw("error in variable payload validation", "err", err)
		return err
	}

	auditLog := getAuditLog(payload)
	// Begin Transaction
	tx, err := impl.scopedVariableRepository.StartTx()
	if err != nil {
		impl.logger.Errorw("error in starting transaction of variable creation", "err", err)
		return err
	}
	// Rollback Transaction in case of any error
	defer func(scopedVariableRepository repository2.ScopedVariableRepository, tx *pg.Tx) {
		err = scopedVariableRepository.RollbackTx(tx)
		if err != nil {
			impl.logger.Errorw("error in rollback transaction of variable creation", "err", err)
			return
		}
	}(impl.scopedVariableRepository, tx)

	err = impl.scopedVariableRepository.DeleteVariables(auditLog, tx)
	if err != nil {
		impl.logger.Errorw("error in deleting variables", "err", err)
		return err
	}
	if len(payload.Variables) != 0 {
		varNameIdMap, err := impl.storeVariableDefinitions(payload, auditLog, tx)
		if err != nil {
			return err
		}

		scopeIdToVarData, err := impl.createVariableScopes(payload, varNameIdMap, auditLog, tx)
		if err != nil {
			return err
		}
		err = impl.storeVariableData(scopeIdToVarData, auditLog, tx)
		if err != nil {
			return err
		}

	}
	err = impl.scopedVariableRepository.CommitTx(tx)
	if err != nil {
		impl.logger.Errorw("error in committing transaction of variable creation", "err", err)
		return err
	}
	loadVariableCache(impl.VariableNameConfig, impl)
	return nil
}

func (impl *ScopedVariableServiceImpl) storeVariableData(scopeIdToVarData map[int]string, auditLog sql.AuditLog, tx *pg.Tx) error {
	VariableDataList := make([]*repository2.VariableData, 0)
	for scopeId, data := range scopeIdToVarData {
		varData := &repository2.VariableData{
			VariableScopeId: scopeId,
			Data:            data,
			AuditLog:        auditLog,
		}
		VariableDataList = append(VariableDataList, varData)
	}
	if len(VariableDataList) > 0 {
		err := impl.scopedVariableRepository.CreateVariableData(VariableDataList, tx)
		if err != nil {
			impl.logger.Errorw("error in saving variable data", "err", err)
			return err
		}
	}
	return nil
}

func (impl *ScopedVariableServiceImpl) storeVariableDefinitions(payload models.Payload, auditLog sql.AuditLog, tx *pg.Tx) (map[string]int, error) {
	variableDefinitions := make([]*repository2.VariableDefinition, 0, len(payload.Variables))
	for _, variable := range payload.Variables {
		variableDefinition := repository2.CreateFromDefinition(variable.Definition, auditLog)
		variableDefinitions = append(variableDefinitions, variableDefinition)
	}
	varDef, err := impl.scopedVariableRepository.CreateVariableDefinition(variableDefinitions, tx)
	if err != nil {
		impl.logger.Errorw("error occurred while saving variable definition", "err", err)
		return nil, err
	}
	variableNameToId := make(map[string]int)
	for _, variable := range varDef {
		variableNameToId[variable.Name] = variable.Id
	}
	return variableNameToId, nil
}

func (impl *ScopedVariableServiceImpl) createVariableScopes(payload models.Payload, variableNameToId map[string]int, auditLog sql.AuditLog, tx *pg.Tx) (map[int]string, error) {

	variableScopes := make([]*repository2.VariableScope, 0)
	for _, variable := range payload.Variables {

		variableId := variableNameToId[variable.Definition.VarName]
		for _, value := range variable.AttributeValues {
			var varValue string
			varValue, err := utils.StringifyValue(value.VariableValue.Value)
			if err != nil {
				return nil, err
			}
			if value.AttributeType == models.Global {
				scope := &repository2.VariableScope{
					VariableDefinitionId: variableId,
					QualifierId:          int(helper.GetQualifierId(value.AttributeType)),
					Active:               true,
					Data:                 varValue,
					AuditLog:             auditLog,
				}
				variableScopes = append(variableScopes, scope)
			}
		}
	}
	parentVariableScope := make([]*repository2.VariableScope, 0)
	childrenVariableScope := make([]*repository2.VariableScope, 0)
	parentScopesMap := make(map[string]*repository2.VariableScope)

	var parentVarScope []*repository2.VariableScope

	for _, scope := range variableScopes {
		parentVariableScope = append(parentVariableScope, scope)
	}
	var err error
	if len(parentVariableScope) > 0 {
		parentVarScope, err = impl.scopedVariableRepository.CreateVariableScope(parentVariableScope, tx)
		if err != nil {
			impl.logger.Errorw("error in getting parentVarScope", parentVarScope)
			return nil, err
		}
	}
	scopeIdToVarData := make(map[int]string)
	for _, parentVar := range parentVarScope {
		scopeIdToVarData[parentVar.Id] = parentVar.Data
	}
	for _, childScope := range childrenVariableScope {
		parentScope, exists := parentScopesMap[childScope.CompositeKey]
		if exists {
			childScope.ParentIdentifier = parentScope.Id
		}
	}
	if len(childrenVariableScope) > 0 {
		childVarScope, err := impl.scopedVariableRepository.CreateVariableScope(childrenVariableScope, tx)
		if err != nil {
			impl.logger.Errorw("error in getting childVarScope", err, childVarScope)
			return nil, err
		}
	}
	return scopeIdToVarData, nil
}

func (impl *ScopedVariableServiceImpl) getMatchedScopedVariables(varScope []*repository2.VariableScope) map[int]int {
	variableIdToVariableScopes := make(map[int][]*repository2.VariableScope)
	variableIdToSelectedScopeId := make(map[int]int)
	for _, vScope := range varScope {
		variableId := vScope.VariableDefinitionId
		variableIdToVariableScopes[variableId] = append(variableIdToVariableScopes[variableId], vScope)
	}
	// Filter out the unneeded scoped which were fetched from DB for the same variable and qualifier
	for variableId, scopes := range variableIdToVariableScopes {

		selectedScopes := make([]*repository2.VariableScope, 0)
		compoundQualifierToScopes := make(map[repository2.Qualifier][]*repository2.VariableScope)

		for _, variableScope := range scopes {
			qualifier := repository2.Qualifier(variableScope.QualifierId)
			if slices.Contains(repository2.CompoundQualifiers, qualifier) {
				compoundQualifierToScopes[qualifier] = append(compoundQualifierToScopes[qualifier], variableScope)
			} else {
				selectedScopes = append(selectedScopes, variableScope)
			}
		}

		for _, qualifier := range repository2.CompoundQualifiers {
			selectedScope := impl.selectScopeForCompoundQualifier(compoundQualifierToScopes[qualifier], repository2.GetNumOfChildQualifiers(qualifier))
			if selectedScope != nil {
				selectedScopes = append(selectedScopes, selectedScope)
			}
		}
		variableIdToVariableScopes[variableId] = selectedScopes
	}

	var minScope *repository2.VariableScope
	for variableId, scopes := range variableIdToVariableScopes {
		minScope = helper.FindMinWithComparator(scopes, helper.QualifierComparator)
		if minScope != nil {
			variableIdToSelectedScopeId[variableId] = minScope.Id
		}
	}
	return variableIdToSelectedScopeId
}

func (impl *ScopedVariableServiceImpl) selectScopeForCompoundQualifier(scopes []*repository2.VariableScope, numQualifiers int) *repository2.VariableScope {
	parentIdToChildScopes := make(map[int][]*repository2.VariableScope)
	parentScopeIdToScope := make(map[int]*repository2.VariableScope, 0)
	parentScopeIds := make([]int, 0)
	for _, scope := range scopes {
		// is not parent so append it to the list in the map with key as its parent scopeID
		if scope.ParentIdentifier > 0 {
			parentIdToChildScopes[scope.ParentIdentifier] = append(parentIdToChildScopes[scope.ParentIdentifier], scope)
		} else {
			//is parent so collect IDs and put it in a map for easy retrieval
			parentScopeIds = append(parentScopeIds, scope.Id)
			parentScopeIdToScope[scope.Id] = scope
		}
	}

	for parentScopeId, _ := range parentIdToChildScopes {
		// this deletes the keys in the map where the key does not exist in the collected IDs for parent
		if !slices.Contains(parentScopeIds, parentScopeId) {
			delete(parentIdToChildScopes, parentScopeId)
		}
	}

	// Now in the map only those will exist with all child matched or partial matches.
	// Because only one will entry exist with all matched we'll return that scope.
	var selectedParentScope *repository2.VariableScope
	for parentScopeId, childScopes := range parentIdToChildScopes {
		if len(childScopes) == numQualifiers {
			selectedParentScope = parentScopeIdToScope[parentScopeId]
		}
	}
	return selectedParentScope
}

func (impl *ScopedVariableServiceImpl) GetScopedVariables(scope models.Scope, varNames []string, maskSensitiveData bool) (scopedVariableDataObj []*models.ScopedVariableData, err error) {

	// getting all variables from cache
	allVariableDefinitions := impl.VariableCache.GetData()

	// cache is loaded and no active variables exist. Returns empty
	if allVariableDefinitions != nil && len(allVariableDefinitions) == 0 {
		return nil, nil
	}

	// Need to get from repo for isSensitive even if cache is loaded since cache only contains metadata
	if allVariableDefinitions == nil {
		allVariableDefinitions, err = impl.scopedVariableRepository.GetAllVariables()

		//Cache was not loaded and no active variables found
		if len(allVariableDefinitions) == 0 {
			return nil, nil
		}
	}

	// filtering out variables whose name is not present in varNames
	var variableDefinitions []*repository2.VariableDefinition
	for _, definition := range allVariableDefinitions {
		// we don't apply filter logic when provided varNames is nil
		if varNames == nil || slices.Contains(varNames, definition.Name) {
			variableDefinitions = append(variableDefinitions, definition)
		}
	}

	variableIds := make([]int, 0)
	variableIdToDefinition := make(map[int]*repository2.VariableDefinition)
	for _, definition := range variableDefinitions {
		variableIds = append(variableIds, definition.Id)
		variableIdToDefinition[definition.Id] = definition
	}

	// This to prevent corner case where no variables were found for the provided names
	if len(varNames) > 0 && len(variableIds) == 0 {
		return make([]*models.ScopedVariableData, 0), nil
	}

	varScope, err := impl.scopedVariableRepository.GetScopedVariableData(scope, variableIds)
	if err != nil {
		impl.logger.Errorw("error in getting varScope", "err", err)
		return nil, err
	}

	variableIdToSelectedScopeId := impl.getMatchedScopedVariables(varScope)

	scopeIds := make([]int, 0)
	foundVarIds := make([]int, 0) // the variable IDs which have data
	for varId, scopeId := range variableIdToSelectedScopeId {
		scopeIds = append(scopeIds, scopeId)
		foundVarIds = append(foundVarIds, varId)
	}
	var variableData []*repository2.VariableData
	if len(scopeIds) != 0 {
		variableData, err = impl.scopedVariableRepository.GetDataForScopeIds(scopeIds)
		if err != nil {
			impl.logger.Errorw("error in getting variable data", "err", err)
			return nil, err
		}
	}

	scopeIdToVarData := make(map[int]*repository2.VariableData)
	for _, varData := range variableData {
		scopeIdToVarData[varData.VariableScopeId] = varData
	}

	for varId, scopeId := range variableIdToSelectedScopeId {
		var value interface{}
		value, err = utils.DestringifyValue(scopeIdToVarData[scopeId].Data)
		if err != nil {
			impl.logger.Errorw("error in validating value", "err", err)
			return nil, err
		}

		var varValue *models.VariableValue
		var isRedacted bool
		if !maskSensitiveData && variableIdToDefinition[varId].VarType == models.PRIVATE {
			varValue = &models.VariableValue{Value: ""}
			isRedacted = true
		} else {
			varValue = &models.VariableValue{Value: value}
		}
		scopedVariableData := &models.ScopedVariableData{
			VariableName:     variableIdToDefinition[varId].Name,
			ShortDescription: variableIdToDefinition[varId].ShortDescription,
			VariableValue:    varValue,
			IsRedacted:       isRedacted}

		scopedVariableDataObj = append(scopedVariableDataObj, scopedVariableData)
	}

	//adding variable def for variables which don't have any scoped data defined
	// This only happens when passed var names is null (called from UI to get all variables with or without data)
	if varNames == nil {
		for _, definition := range allVariableDefinitions {
			if !slices.Contains(foundVarIds, definition.Id) {
				scopedVariableDataObj = append(scopedVariableDataObj, &models.ScopedVariableData{
					VariableName:     definition.Name,
					ShortDescription: definition.ShortDescription,
				})
			}
		}
	}

	return scopedVariableDataObj, err
}

func (impl *ScopedVariableServiceImpl) GetJsonForVariables() (*models.Payload, error) {

	// getting all variables from cache, if empty then no variables exist
	allVariableDefinitions := impl.VariableCache.GetData()
	if allVariableDefinitions != nil && len(allVariableDefinitions) == 0 {
		return nil, nil
	}
	dataForJson, err := impl.scopedVariableRepository.GetAllVariableScopeAndDefinition()
	if err != nil {
		impl.logger.Errorw("error in getting data for json", "err", err)
		return nil, err
	}

	payload := &models.Payload{
		Variables: make([]*models.Variables, 0),
	}
	variables := make([]*models.Variables, 0)

	for _, data := range dataForJson {
		definition := models.Definition{
			VarName:          data.Name,
			DataType:         data.DataType,
			VarType:          data.VarType,
			Description:      data.Description,
			ShortDescription: data.ShortDescription,
		}
		attributes := make([]models.AttributeValue, 0)

		scopeIdToVarScopes := make(map[int][]*repository2.VariableScope)
		for _, scope := range data.VariableScope {
			if scope.ParentIdentifier != 0 {
				scopeIdToVarScopes[scope.ParentIdentifier] = append(scopeIdToVarScopes[scope.ParentIdentifier], scope)
			} else {
				scopeIdToVarScopes[scope.Id] = []*repository2.VariableScope{scope}
			}
		}
		for parentScopeId, scopes := range scopeIdToVarScopes {
			attribute := models.AttributeValue{
				AttributeParams: make(map[models.IdentifierType]string),
			}
			for _, scope := range scopes {
				if parentScopeId == scope.Id {
					var value interface{}
					value, err = utils.DestringifyValue(scope.VariableData.Data)
					if err != nil {
						return nil, err
					}
					attribute.VariableValue = models.VariableValue{
						Value: value,
					}
					attribute.AttributeType = helper.GetAttributeType(repository2.Qualifier(scope.QualifierId))
				}
			}
			if len(attribute.AttributeParams) == 0 {
				attribute.AttributeParams = nil
			}
			attributes = append(attributes, attribute)
		}

		variable := &models.Variables{
			Definition:      definition,
			AttributeValues: attributes,
		}
		variables = append(variables, variable)
	}

	payload.Variables = variables
	if len(payload.Variables) == 0 {
		return nil, nil
	}
	return payload, nil
}

func getAuditLog(payload models.Payload) sql.AuditLog {
	auditLog := sql.AuditLog{
		CreatedOn: time.Now(),
		CreatedBy: payload.UserId,
		UpdatedOn: time.Now(),
		UpdatedBy: payload.UserId,
	}
	return auditLog
}