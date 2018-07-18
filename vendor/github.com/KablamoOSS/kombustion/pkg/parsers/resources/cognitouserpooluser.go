package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CognitoUserPoolUser Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html
type CognitoUserPoolUser struct {
	Type       string                        `yaml:"Type"`
	Properties CognitoUserPoolUserProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

// CognitoUserPoolUser Properties
type CognitoUserPoolUserProperties struct {
	ForceAliasCreation     interface{} `yaml:"ForceAliasCreation,omitempty"`
	MessageAction          interface{} `yaml:"MessageAction,omitempty"`
	UserPoolId             interface{} `yaml:"UserPoolId"`
	Username               interface{} `yaml:"Username,omitempty"`
	DesiredDeliveryMediums interface{} `yaml:"DesiredDeliveryMediums,omitempty"`
	UserAttributes         interface{} `yaml:"UserAttributes,omitempty"`
	ValidationData         interface{} `yaml:"ValidationData,omitempty"`
}

// NewCognitoUserPoolUser constructor creates a new CognitoUserPoolUser
func NewCognitoUserPoolUser(properties CognitoUserPoolUserProperties, deps ...interface{}) CognitoUserPoolUser {
	return CognitoUserPoolUser{
		Type:       "AWS::Cognito::UserPoolUser",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCognitoUserPoolUser parses CognitoUserPoolUser
func ParseCognitoUserPoolUser(name string, data string) (cf types.TemplateObject, err error) {
	var resource CognitoUserPoolUser
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CognitoUserPoolUser - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseCognitoUserPoolUser validator
func (resource CognitoUserPoolUser) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCognitoUserPoolUserProperties validator
func (resource CognitoUserPoolUserProperties) Validate() []error {
	errs := []error{}
	if resource.UserPoolId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UserPoolId'"))
	}
	return errs
}
