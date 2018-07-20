package parsers

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
func ParseExampleLog(
	name string,
	data string,
) (
	conditions kombustionTypes.TemplateObject,
	metadata kombustionTypes.TemplateObject,
	mappings kombustionTypes.TemplateObject,
	outputs kombustionTypes.TemplateObject,
	parameters kombustionTypes.TemplateObject,
	resources kombustionTypes.TemplateObject,
	transform kombustionTypes.TemplateObject,
	errors []error,
) {
	// Parse the config data
	var config LogConfig

	err := yaml.Unmarshal([]byte(data), &config)

	// If there are errors unmarshalling, add them to the errors array and return
	if err != nil {
		errors = append(errors, err)
		return
	}

	// validate the config
	validateErrs := config.Validate()

	// If there are any validation errors add them to our errors array and return
	if validateErrs != nil {
		errors = append(errors, validateErrs...)
		return
	}

	// validate the config
	config.Validate()

	logGroupName := fmt.Sprintf("%s%s", name, config.Properties.LogGroupName)
	// Create a new resource and add it to cf
	resources = kombustionTypes.TemplateObject{
		// (name) lets us create a new resource with the same name as the input type
		(name): cfResources.NewLogsLogGroup(
			cfResources.LogsLogGroupProperties{
				LogGroupName:    fmt.Sprintf("!Join [ \"-\", [ !Ref %s, !Ref Environment ] ]", logGroupName),
				RetentionInDays: config.Properties.RetentionInDays,
			},
		),
	}

	parameters = kombustionTypes.TemplateObject{
		(logGroupName): map[string]string{
			"Type": "String",
		},
	}

	metadata = kombustionTypes.TemplateObject{
		(logGroupName): map[string]string{
			"Source": "kablamo-plugin-boilerplate",
		},
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
