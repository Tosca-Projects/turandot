package client

import (
	"strings"

	"github.com/google/uuid"
	spoolerpkg "github.com/tliron/kubernetes-registry-spooler/client"
	"github.com/tliron/puccini/common/format"
	urlpkg "github.com/tliron/puccini/url"
	"github.com/tliron/turandot/common"
	resources "github.com/tliron/turandot/resources/turandot.puccini.cloud/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (self *Client) DeployServiceFromTemplate(serviceName string, serviceTemplateName string, inputs map[string]interface{}) error {
	if url, err := self.GetInventoryServiceTemplateURL(serviceTemplateName); err == nil {
		defer url.Release()
		_, err := self.createService(serviceName, url, inputs)
		return err
	} else {
		return err
	}
}

func (self *Client) DeployServiceFromURL(serviceName string, url string, inputs map[string]interface{}) error {
	if url_, err := urlpkg.NewURL(url); err == nil {
		defer url_.Release()
		_, err = self.createService(serviceName, url_, inputs)
		return err
	} else {
		return err
	}
}

func (self *Client) DeployServiceFromContent(serviceName string, spooler *spoolerpkg.Client, url urlpkg.URL, inputs map[string]interface{}) error {
	serviceTemplateName := uuid.New().String()
	imageName := GetInventoryImageName(serviceTemplateName)
	if err := common.PushToRegistry(imageName, url, spooler); err == nil {
		return self.DeployServiceFromTemplate(serviceName, serviceTemplateName, inputs)
	} else {
		return err
	}
}

func (self *Client) DeleteService(serviceName string) error {
	return self.turandot.TurandotV1alpha1().Services(self.namespace).Delete(self.context, serviceName, meta.DeleteOptions{})
}

func (self *Client) ListServices() (*resources.ServiceList, error) {
	return self.turandot.TurandotV1alpha1().Services(self.namespace).List(self.context, meta.ListOptions{})
}

func (self *Client) createService(name string, url urlpkg.URL, inputs map[string]interface{}) (*resources.Service, error) {
	// Encode inputs
	inputs_ := make(map[string]string)
	for key, input := range inputs {
		var err error
		if inputs_[key], err = format.EncodeYAML(input, " ", false); err == nil {
			inputs_[key] = strings.TrimRight(inputs_[key], "\n")
		} else {
			return nil, err
		}
	}

	service := &resources.Service{
		ObjectMeta: meta.ObjectMeta{
			Name:      name,
			Namespace: self.namespace,
		},
		Spec: resources.ServiceSpec{
			ServiceTemplateURL: url.String(),
			Inputs:             inputs_,
		},
	}

	if service, err := self.turandot.TurandotV1alpha1().Services(self.namespace).Create(self.context, service, meta.CreateOptions{}); err == nil {
		return service, nil
	} else if errors.IsAlreadyExists(err) {
		return self.turandot.TurandotV1alpha1().Services(self.namespace).Get(self.context, name, meta.GetOptions{})
	} else {
		return nil, err
	}
}