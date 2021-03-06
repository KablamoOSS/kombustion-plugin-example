package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeliveryStreamSplunkDestinationConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-splunkdestinationconfiguration.html
type DeliveryStreamSplunkDestinationConfiguration struct {
	HECAcknowledgmentTimeoutInSeconds interface{}                               `yaml:"HECAcknowledgmentTimeoutInSeconds,omitempty"`
	HECEndpoint                       interface{}                               `yaml:"HECEndpoint"`
	HECEndpointType                   interface{}                               `yaml:"HECEndpointType"`
	HECToken                          interface{}                               `yaml:"HECToken"`
	S3BackupMode                      interface{}                               `yaml:"S3BackupMode,omitempty"`
	RetryOptions                      *DeliveryStreamSplunkRetryOptions         `yaml:"RetryOptions,omitempty"`
	S3Configuration                   *DeliveryStreamS3DestinationConfiguration `yaml:"S3Configuration"`
	ProcessingConfiguration           *DeliveryStreamProcessingConfiguration    `yaml:"ProcessingConfiguration,omitempty"`
	CloudWatchLoggingOptions          *DeliveryStreamCloudWatchLoggingOptions   `yaml:"CloudWatchLoggingOptions,omitempty"`
}

// DeliveryStreamSplunkDestinationConfiguration validation
func (resource DeliveryStreamSplunkDestinationConfiguration) Validate() []error {
	errors := []error{}

	if resource.HECEndpoint == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'HECEndpoint'"))
	}
	if resource.HECEndpointType == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'HECEndpointType'"))
	}
	if resource.HECToken == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'HECToken'"))
	}
	if resource.S3Configuration == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'S3Configuration'"))
	} else {
		errors = append(errors, resource.S3Configuration.Validate()...)
	}
	return errors
}
