package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ConfigurationSetEventDestinationDimensionConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-dimensionconfiguration.html
type ConfigurationSetEventDestinationDimensionConfiguration struct {
	DefaultDimensionValue interface{} `yaml:"DefaultDimensionValue"`
	DimensionName         interface{} `yaml:"DimensionName"`
	DimensionValueSource  interface{} `yaml:"DimensionValueSource"`
}

// ConfigurationSetEventDestinationDimensionConfiguration validation
func (resource ConfigurationSetEventDestinationDimensionConfiguration) Validate() []error {
	errs := []error{}

	if resource.DefaultDimensionValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DefaultDimensionValue'"))
	}
	if resource.DimensionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DimensionName'"))
	}
	if resource.DimensionValueSource == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DimensionValueSource'"))
	}
	return errs
}