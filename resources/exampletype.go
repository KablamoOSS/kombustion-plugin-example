package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/pluginParsers/resources"
	"github.com/KablamoOSS/kombustion/types"
)

type MyConfig struct {
	Properties struct {
		// TODO: fill me in
	} `yaml:"Properties"`
}

func ParseMyType(name string, data string) (cf types.ValueMap, err error) {
	// Parse the config data
	var config MyConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// validate the config
	config.Validate()

	// create a group of objects (each to be validated)
	cf = make(types.ValueMap)

	// defaults
	myproperty := 1
	if config.Properties.MyProperty != nil {
		myproperty = *config.Properties.MyProperty
	}

	// create the resource
	cf[name] = resources.NewAWSType(
		resources.AWSTypeProperties{},
	)

	return
}

// Validate - input Config validation
func (this MyConfig) Validate() {

}
