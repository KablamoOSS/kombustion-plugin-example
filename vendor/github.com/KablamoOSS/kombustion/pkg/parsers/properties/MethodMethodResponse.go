package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// MethodMethodResponse Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html
type MethodMethodResponse struct {
	StatusCode         interface{} `yaml:"StatusCode"`
	ResponseModels     interface{} `yaml:"ResponseModels,omitempty"`
	ResponseParameters interface{} `yaml:"ResponseParameters,omitempty"`
}

// MethodMethodResponse validation
func (resource MethodMethodResponse) Validate() []error {
	errors := []error{}

	if resource.StatusCode == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'StatusCode'"))
	}
	return errors
}
