package che

import (
	"k8s.io/api/authentication/v1"
	"net/http"
	"os"

	brokerapi "github.com/integr8ly/managed-service-broker/pkg/broker"
	"github.com/integr8ly/managed-service-broker/pkg/clients/openshift"
	glog "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

type CheDeployer struct {
	id string
}

func NewDeployer(id string) *CheDeployer {
	return &CheDeployer{id: id}
}

func (cd *CheDeployer) IsForService(serviceID string) bool {
	return serviceID == "che-service-id"
}

func (cd *CheDeployer) GetCatalogEntries() []*brokerapi.Service {
	glog.Infof("Getting che catalog entries")
	return getCatalogServicesObj()
}

func (cd *CheDeployer) GetID() string {
	return cd.id
}

func (cd *CheDeployer) Deploy(instanceID, managedNamespace string, contextProfile brokerapi.ContextProfile, parameters map[string]interface{}, userInfo v1.UserInfo, k8sclient kubernetes.Interface, osClientFactory *openshift.ClientFactory) (*brokerapi.CreateServiceInstanceResponse, error) {
	glog.Infof("Deploying che from deployer, id: %s", instanceID)

	dashboardUrl := os.Getenv("CHE_DASHBOARD_URL")

	return &brokerapi.CreateServiceInstanceResponse{
		Code:         http.StatusAccepted,
		DashboardURL: dashboardUrl,
	}, nil
}

func (cd *CheDeployer) RemoveDeploy(serviceInstanceId string, namespace string, k8sclient kubernetes.Interface) error {
	return nil
}

func (cd *CheDeployer) LastOperation(instanceID, namespace string, k8sclient kubernetes.Interface, osclient *openshift.ClientFactory) (*brokerapi.LastOperationResponse, error) {
	glog.Infof("Getting last operation for %s", instanceID)

	return &brokerapi.LastOperationResponse{
		State:       brokerapi.StateSucceeded,
		Description: "che deployed successfully",
	}, nil
}
