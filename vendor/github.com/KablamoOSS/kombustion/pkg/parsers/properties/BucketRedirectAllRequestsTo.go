package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketRedirectAllRequestsTo Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-redirectallrequeststo.html
type BucketRedirectAllRequestsTo struct {
	HostName interface{} `yaml:"HostName"`
	Protocol interface{} `yaml:"Protocol,omitempty"`
}

// BucketRedirectAllRequestsTo validation
func (resource BucketRedirectAllRequestsTo) Validate() []error {
	errors := []error{}

	if resource.HostName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'HostName'"))
	}
	return errors
}
