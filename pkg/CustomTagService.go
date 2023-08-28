package pkg

import (
	"fmt"
	"github.com/devtron-labs/devtron/api/bean"
	"github.com/devtron-labs/devtron/internal/sql/repository"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

const (
	EntityTypeCiPipelineId = iota
)

const (
	imagePathPattern           = "%s/%s:%s" // dockerReg/dockerRepo:Tag
	ImageTagUnavailableMessage = "Desired image tag already exists"
)

var (
	ErrImagePathInUse = fmt.Errorf("image path is already being used by someone")
)

type CustomTagService interface {
	CreateOrUpdateCustomTag(tag *bean.CustomTag) error
	GetCustomTagByEntityKeyAndValue(entityKey int, entityValue string) (*repository.CustomTag, error)
	GenerateImagePath(entityKey int, entityValue string, dockerRegistryURL string, dockerRepo string) (*repository.ImagePathReservation, error)
	DeleteCustomTagIfExists(tag bean.CustomTag) error
}

type CustomTagServiceImpl struct {
	Logger              *zap.SugaredLogger
	customTagRepository repository.ImageTagRepository
}

func NewCustomTagService(logger *zap.SugaredLogger, customTagRepo repository.ImageTagRepository) *CustomTagServiceImpl {
	return &CustomTagServiceImpl{
		Logger:              logger,
		customTagRepository: customTagRepo,
	}
}

func (impl *CustomTagServiceImpl) CreateOrUpdateCustomTag(tag *bean.CustomTag) error {
	customTagData := repository.CustomTag{
		EntityKey:            tag.EntityKey,
		EntityValue:          tag.EntityValue,
		TagPattern:           strings.ReplaceAll(tag.TagPattern, "{X}", "{x}"),
		AutoIncreasingNumber: tag.AutoIncreasingNumber,
		Metadata:             tag.Metadata,
	}
	oldTagObject, err := impl.customTagRepository.FetchCustomTagData(customTagData.EntityKey, customTagData.EntityValue)
	if err != nil && err != pg.ErrNoRows {
		return err
	}
	if oldTagObject.Id == 0 {
		return impl.customTagRepository.CreateImageTag(&customTagData)
	} else {
		customTagData.Id = oldTagObject.Id
		return impl.customTagRepository.UpdateImageTag(&customTagData)
	}
}

func (impl *CustomTagServiceImpl) DeleteCustomTagIfExists(tag bean.CustomTag) error {
	return impl.customTagRepository.DeleteByEntityKeyAndValue(tag.EntityKey, tag.EntityValue)
}

func (impl *CustomTagServiceImpl) GetCustomTagByEntityKeyAndValue(entityKey int, entityValue string) (*repository.CustomTag, error) {
	return impl.customTagRepository.FetchCustomTagData(entityKey, entityValue)
}

func (impl *CustomTagServiceImpl) GenerateImagePath(entityKey int, entityValue string, dockerRegistryURL string, dockerRepo string) (*repository.ImagePathReservation, error) {
	connection := impl.customTagRepository.GetConnection()
	tx, err := connection.Begin()
	if err != nil {
		return nil, nil
	}
	defer tx.Rollback()
	customTagData, err := impl.customTagRepository.IncrementAndFetchByEntityKeyAndValue(tx, entityKey, entityValue)
	if err != nil {
		return nil, err
	}
	tag, err := validateAndConstructTag(customTagData)
	if err != nil {
		return nil, err
	}
	imagePath := fmt.Sprintf(imagePathPattern, dockerRegistryURL, dockerRepo, tag)
	imagePathReservations, err := impl.customTagRepository.FindByImagePath(tx, imagePath)
	if err != nil && err != pg.ErrNoRows {
		return nil, err
	}
	if len(imagePathReservations) > 0 {
		return nil, ErrImagePathInUse
	}
	imagePathReservation := repository.ImagePathReservation{
		ImagePath:   imagePath,
		CustomTagId: customTagData.Id,
	}
	err = impl.customTagRepository.InsertImagePath(tx, imagePathReservation)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &imagePathReservation, nil
}

func validateAndConstructTag(customTagData *repository.CustomTag) (string, error) {
	err := validateTagPattern(customTagData.TagPattern)
	if err != nil {
		return "", err
	}
	dockerImageTag := strings.ReplaceAll(customTagData.TagPattern, "{x}", strconv.Itoa(customTagData.AutoIncreasingNumber-1)) //-1 because number is already incremented, current value will be used next time
	err = validateTag(dockerImageTag)
	if err != nil {
		return "", err
	}
	return dockerImageTag, nil
}

func validateTagPattern(customTagPattern string) error {
	allowedVariables := []string{"{x}", "{X}"}
	totalX := 0
	for _, variable := range allowedVariables {
		totalX += strings.Count(customTagPattern, variable)
	}
	if totalX != 1 {
		return fmt.Errorf("variable {x} is allowed exactly once")
	}
	return nil
}

func validateTag(imageTag string) error {
	if len(imageTag) == 0 || len(imageTag) > 128 {
		return fmt.Errorf("image tag should be of len 1-128 only, imageTag: %s", imageTag)
	}
	allowedSymbols := ".abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVWXYZ-0987654321"
	allowedCharSet := make(map[int32]struct{})
	for _, c := range allowedSymbols {
		allowedCharSet[c] = struct{}{}
	}
	firstChar := imageTag[0:1]
	if firstChar == "." || firstChar == "-" {
		return fmt.Errorf("image tag can not start with a period or a hyphen, imageTag: %s", imageTag)
	}
	return nil
}
