package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CodeDeployDeploymentConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html
type CodeDeployDeploymentConfig struct {
	Type       string                               `yaml:"Type"`
	Properties CodeDeployDeploymentConfigProperties `yaml:"Properties"`
	Condition  interface{}                          `yaml:"Condition,omitempty"`
	Metadata   interface{}                          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                          `yaml:"DependsOn,omitempty"`
}

// CodeDeployDeploymentConfig Properties
type CodeDeployDeploymentConfigProperties struct {
	DeploymentConfigName interface{}                                     `yaml:"DeploymentConfigName,omitempty"`
	MinimumHealthyHosts  *properties.DeploymentConfigMinimumHealthyHosts `yaml:"MinimumHealthyHosts,omitempty"`
}

// NewCodeDeployDeploymentConfig constructor creates a new CodeDeployDeploymentConfig
func NewCodeDeployDeploymentConfig(properties CodeDeployDeploymentConfigProperties, deps ...interface{}) CodeDeployDeploymentConfig {
	return CodeDeployDeploymentConfig{
		Type:       "AWS::CodeDeploy::DeploymentConfig",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCodeDeployDeploymentConfig parses CodeDeployDeploymentConfig
func ParseCodeDeployDeploymentConfig(name string, data string) (cf types.TemplateObject, err error) {
	var resource CodeDeployDeploymentConfig
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CodeDeployDeploymentConfig - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseCodeDeployDeploymentConfig validator
func (resource CodeDeployDeploymentConfig) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCodeDeployDeploymentConfigProperties validator
func (resource CodeDeployDeploymentConfigProperties) Validate() []error {
	errs := []error{}
	return errs
}