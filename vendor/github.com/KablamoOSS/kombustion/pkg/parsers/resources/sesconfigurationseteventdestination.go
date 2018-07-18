package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// SESConfigurationSetEventDestination Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationseteventdestination.html
type SESConfigurationSetEventDestination struct {
	Type       string                                        `yaml:"Type"`
	Properties SESConfigurationSetEventDestinationProperties `yaml:"Properties"`
	Condition  interface{}                                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                   `yaml:"DependsOn,omitempty"`
}

// SESConfigurationSetEventDestination Properties
type SESConfigurationSetEventDestinationProperties struct {
	ConfigurationSetName interface{}                                                  `yaml:"ConfigurationSetName"`
	EventDestination     *properties.ConfigurationSetEventDestinationEventDestination `yaml:"EventDestination"`
}

// NewSESConfigurationSetEventDestination constructor creates a new SESConfigurationSetEventDestination
func NewSESConfigurationSetEventDestination(properties SESConfigurationSetEventDestinationProperties, deps ...interface{}) SESConfigurationSetEventDestination {
	return SESConfigurationSetEventDestination{
		Type:       "AWS::SES::ConfigurationSetEventDestination",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSESConfigurationSetEventDestination parses SESConfigurationSetEventDestination
func ParseSESConfigurationSetEventDestination(name string, data string) (cf types.TemplateObject, err error) {
	var resource SESConfigurationSetEventDestination
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SESConfigurationSetEventDestination - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseSESConfigurationSetEventDestination validator
func (resource SESConfigurationSetEventDestination) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSESConfigurationSetEventDestinationProperties validator
func (resource SESConfigurationSetEventDestinationProperties) Validate() []error {
	errs := []error{}
	if resource.ConfigurationSetName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ConfigurationSetName'"))
	}
	if resource.EventDestination == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EventDestination'"))
	} else {
		errs = append(errs, resource.EventDestination.Validate()...)
	}
	return errs
}
