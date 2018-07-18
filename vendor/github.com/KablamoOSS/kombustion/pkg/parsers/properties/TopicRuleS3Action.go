package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TopicRuleS3Action Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-s3action.html
type TopicRuleS3Action struct {
	BucketName interface{} `yaml:"BucketName"`
	Key        interface{} `yaml:"Key"`
	RoleArn    interface{} `yaml:"RoleArn"`
}

// TopicRuleS3Action validation
func (resource TopicRuleS3Action) Validate() []error {
	errs := []error{}

	if resource.BucketName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketName'"))
	}
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	return errs
}
