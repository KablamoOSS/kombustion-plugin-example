package resources

import (
	"fmt"

	cfResources "github.com/KablamoOSS/kombustion/pkg/parsers/resources"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// LogConfig - This is the shape of the YAML users of your plugin will write
type LogConfig struct {
	Properties struct {
		LogGroupName    interface{} `yaml:"LogGroupName,omitempty"`
		RetentionInDays interface{} `yaml:"RetentionInDays,omitempty"`
	} `yaml:"Properties"`
}

// ParseExampleLog - Parse the YAML from
func ParseExampleLog(name string, data string) (cf kombustionTypes.TemplateObject, errs []error) {
	// Parse the config data
	var config LogConfig

	err := yaml.Unmarshal([]byte(data), &config)

	// If there are errors unmarshalling, add them to the errs array and return
	if err != nil {
		errs = append(errs, err)
		return
	}

	// validate the config
	validateErrs := config.Validate()

	// If there are any validation errors add them to our errs array and return
	if validateErrs != nil {
		errs = append(errs, validateErrs...)
		return
	}

	// validate the config
	config.Validate()

	// Create a new resource and add it to cf
	cf = kombustionTypes.TemplateObject{
		// (name) lets us create a new resource with the same name as the input type
		(name): cfResources.NewLogsLogGroup(
			cfResources.LogsLogGroupProperties{
				LogGroupName:    config.Properties.LogGroupName,
				RetentionInDays: config.Properties.RetentionInDays,
			},
		),
	}

	return
}

// Validate - input Config validation
func (config LogConfig) Validate() (errors []error) {
	props := config.Properties

	// Ensure LogGrouName has a value
	if props.LogGroupName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'LogGroupName'"))
	}

	// Ensure RetentionInDays has a value
	if props.RetentionInDays == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RetentionInDays'"))
	}

	return
}
