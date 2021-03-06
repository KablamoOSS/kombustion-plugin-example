package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayDeployment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html
type ApiGatewayDeployment struct {
	Type       string                         `yaml:"Type"`
	Properties ApiGatewayDeploymentProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// ApiGatewayDeployment Properties
type ApiGatewayDeploymentProperties struct {
	Description      interface{}                            `yaml:"Description,omitempty"`
	RestApiId        interface{}                            `yaml:"RestApiId"`
	StageName        interface{}                            `yaml:"StageName,omitempty"`
	StageDescription *properties.DeploymentStageDescription `yaml:"StageDescription,omitempty"`
}

// NewApiGatewayDeployment constructor creates a new ApiGatewayDeployment
func NewApiGatewayDeployment(properties ApiGatewayDeploymentProperties, deps ...interface{}) ApiGatewayDeployment {
	return ApiGatewayDeployment{
		Type:       "AWS::ApiGateway::Deployment",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayDeployment parses ApiGatewayDeployment
func ParseApiGatewayDeployment(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource ApiGatewayDeployment
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseApiGatewayDeployment validator
func (resource ApiGatewayDeployment) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayDeploymentProperties validator
func (resource ApiGatewayDeploymentProperties) Validate() []error {
	errors := []error{}
	if resource.RestApiId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errors
}
