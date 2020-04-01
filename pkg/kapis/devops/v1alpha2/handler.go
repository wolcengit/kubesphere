package v1alpha2

import (
	"kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	"kubesphere.io/kubesphere/pkg/client/informers/externalversions"
	"kubesphere.io/kubesphere/pkg/models/devops"
	devopsClient "kubesphere.io/kubesphere/pkg/simple/client/devops"
	"kubesphere.io/kubesphere/pkg/simple/client/mysql"
	"kubesphere.io/kubesphere/pkg/simple/client/s3"
	"kubesphere.io/kubesphere/pkg/simple/client/sonarqube"
)

type ProjectPipelineHandler struct {
	projectMemberOperator   devops.ProjectMemberOperator
	devopsOperator          devops.DevopsOperator
	projectOperator         devops.ProjectOperator
	projectCredentialGetter devops.ProjectCredentialGetter
}

type PipelineSonarHandler struct {
	pipelineSonarGetter devops.PipelineSonarGetter
	projectOperator     devops.ProjectOperator
}

func NewProjectPipelineHandler(devopsClient devopsClient.Interface, dbClient *mysql.Database) ProjectPipelineHandler {
	return ProjectPipelineHandler{
		projectMemberOperator:   devops.NewProjectMemberOperator(devopsClient, dbClient),
		devopsOperator:          devops.NewDevopsOperator(devopsClient),
		projectOperator:         devops.NewProjectOperator(dbClient),
		projectCredentialGetter: devops.NewProjectCredentialOperator(devopsClient),
	}
}

func NewPipelineSonarHandler(devopsClient devopsClient.Interface, sonarClient sonarqube.SonarInterface) PipelineSonarHandler {
	return PipelineSonarHandler{
		pipelineSonarGetter: devops.NewPipelineSonarGetter(devopsClient, sonarClient),
	}
}

func NewS2iBinaryHandler(client versioned.Interface, informers externalversions.SharedInformerFactory, s3Client s3.Interface) S2iBinaryHandler {
	return S2iBinaryHandler{devops.NewS2iBinaryUploader(client, informers, s3Client)}
}
