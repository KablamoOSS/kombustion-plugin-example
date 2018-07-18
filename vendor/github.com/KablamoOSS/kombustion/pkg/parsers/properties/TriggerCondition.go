package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TriggerCondition Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-condition.html
type TriggerCondition struct {
	JobName         interface{} `yaml:"JobName,omitempty"`
	LogicalOperator interface{} `yaml:"LogicalOperator,omitempty"`
	State           interface{} `yaml:"State,omitempty"`
}

// TriggerCondition validation
func (resource TriggerCondition) Validate() []error {
	errs := []error{}

	return errs
}
