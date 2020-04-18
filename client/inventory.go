package client

import (
	"fmt"
	neturlpkg "net/url"
	"strings"

	urlpkg "github.com/tliron/puccini/url"
)

const serviceTemplateCategory = "service-template"

var serviceTemplateImageNamePrefix = fmt.Sprintf("%s/", serviceTemplateCategory)

func GetInventoryImageName(serviceTemplateName string) string {
	return fmt.Sprintf("%s/%s", serviceTemplateCategory, serviceTemplateName)
}

func ServiceTemplateNameFromInventoryImageName(imageName string) (string, bool) {
	if strings.HasPrefix(imageName, serviceTemplateImageNamePrefix) {
		return imageName[len(serviceTemplateImageNamePrefix):], true
	} else {
		return "", false
	}
}

func (self *Client) GetInventoryServiceTemplateURL(serviceTemplateName string) (*urlpkg.DockerURL, error) {
	appName := fmt.Sprintf("%s-inventory", self.namePrefix)
	if ip, err := self.getFirstPodIp(appName); err == nil {
		imageName := GetInventoryImageName(serviceTemplateName)
		url := fmt.Sprintf("docker://%s:5000/%s?format=csar", ip, imageName)
		if url_, err := neturlpkg.ParseRequestURI(url); err == nil {
			return urlpkg.NewDockerURL(url_), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}