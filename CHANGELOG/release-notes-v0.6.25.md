## v0.6.25



## Bugs
- fix:Not able to change Authentication Type for Prometheus. (#4284)
- fix:cve search  (#4347)
- fix: added missing clusterConfig in releaseIdentifier request (#4359)
- fix: Inconsistent behaviour on deploying application from application group v/s devtron cli tool after changing deploymentType  (#4353)
- fix:rbac is not in sync with Actual application status  (#4237)
- fix: cm-env override (#4350)
- fix: launch ephemeral containers targetting containers running with non root access (#4288)
- fix: not able to add env in global configuration (#4330)
- fix: Message not updating for Ci (#4323)
- fix: app-group page breaking for inactive users (#4315)
- fix: limit ci build message length to 250 (#4300)
- fix: Job trigger throws error (#4296)
- fix: Get Artifacts list API is throwing pg no rows error  (#4292)
- fix: updated copy container image plugin for multi arc build support (#4282)
- fix: old custom tag migration (#4273)
- fix: No rows for rest config in case of Default namespace (#4269)
- fix: DevtronApp Deployment Status Fixed (#4267)
- fix: deleted service showing port as missing fix (#4240)
- fix: kubernetes events fix for resources (#4247)
- fix:post-cd plugin not getting added again, after removing and adding again (#4219)
- fix: update docker config overridden flag in linked ci's (#4220)
- fix: setting cipipelineId to parent ci-pipelines id for linked ci's (#4215)
- fix: ImageScanDeployInfo not initialized while fetching from db (#4195)
- fix: Stable dt19 v3 + v4 bugathon bug fixes (#4183)
- fix: inducing overridden DockerRegistryId in case docker registry is overridden (#4178)
- fix: cm cs handling for inheriting type in get sample workflow api (core app routers) (#4189)
## Enhancements
- feat: Enable Notification for Protect Configuration Approval request (#4358)
- feat: Role based access control in JOBS (#4198)
- feat: Bulk Hibernate (#4229)
- feat: Helm async deploy Devtron Apps (#4045)
- feat: copy container images (#4209)
- feat: add plugin via api (#3937)
- feat: Run jobs using system executor (#4161)
- feat: Description for cluster and apps (#4154)
- feat: Plugin for image scanning in Pre/Post step (#4021)
- feat: scoped variable CMCS support and manager layer refactorings (#4174)
- feat: pre-postcd trigger with plugin (#4176)
- feat: upload and download logs/artifact from blob storage configured in external cluster (#4138)
## Documentation
- doc: Added Application Groups Doc (#4275)
- doc: Update gitops.md for Azure DevOps Integration (#4328)
- docs: img size fix (#4301)
- docs: App Configuration Corrections + Structuring (#4235)
- docs: Added Filter Doc in Index (#4253)
- docs: Filter Condition Doc (#4224)
- docs: Updated chart parameters in doc for deployment and rollout deployment chart (#4218)
- doc: ENT+OSS Bifurcation + Descriptions for System Variables (#4230)
- doc: GC Index Additions + Fixes (#4214)
- docs: Scoped Variables Draft + Other Fixes (Ephemeral Doc Alignment) (#3982)
- docs: Added Videos + Fixed Typos and Navigation (#4169)
- doc: added a Config.md to expose all the env variables use in this microservice (#4173)
## Others
- chore: Integrate Terraform CLI, Ansible, and SonarQube Plugins (#4363)
- chore: updated schema for resources (#4266)
- chore: change-ci and app-group fixes (#4332)
- chore: change CI backend infra  (#4251)
- chore: updating dependency tracker maven plugin name  (#4265)
- chore: Update pr-issue-validator.yaml (#4207)
- chore: update maven dependency tracker plugin (#4254)
- chore: artifact api refactoring (#4137)
- chore: resource sql migration script (#4175)
- chore: refactored few type objects (#4171)
- chore: struct type refactored (#4180)



