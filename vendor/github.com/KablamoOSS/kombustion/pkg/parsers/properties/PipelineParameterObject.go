package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// PipelineParameterObject Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects.html
type PipelineParameterObject struct {
	Id         interface{} `yaml:"Id"`
	Attributes interface{} `yaml:"Attributes"`
}

// PipelineParameterObject validation
func (resource PipelineParameterObject) Validate() []error {
	errors := []error{}

	if resource.Id == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Id'"))
	}
	if resource.Attributes == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Attributes'"))
	}
	return errors
}
