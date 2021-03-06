package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ClusterScalingConstraints Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingconstraints.html
type ClusterScalingConstraints struct {
	MaxCapacity interface{} `yaml:"MaxCapacity"`
	MinCapacity interface{} `yaml:"MinCapacity"`
}

// ClusterScalingConstraints validation
func (resource ClusterScalingConstraints) Validate() []error {
	errors := []error{}

	if resource.MaxCapacity == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'MaxCapacity'"))
	}
	if resource.MinCapacity == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'MinCapacity'"))
	}
	return errors
}
