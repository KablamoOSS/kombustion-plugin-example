package main

import (
	"github.com/KablamoOSS/kombustion-plugin-boilerplate/resources"
	"github.com/KablamoOSS/kombustion/pkg/plugins/api"
	"github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
)

var (
	version string
	name    string
)

func init() {
	if version == "" {
		version = "BUILT_FROM_SOURCE"
	}
	if name == "" {
		// Update this to the name of your plugin (you repo name)
		name = "kombustion-plugin-boilerplate"
	}
}

// Register plugin
func Register() []byte {
	return api.RegisterPlugin(types.Config{
		Name:    name,
		Version: version,
		Prefix:  "Kablamo",
		Help: types.Help{
			Description: "An example plugin to create a log group",
			TypeMappings: []types.TypeMapping{
				{
					Name:        "Example::Log",
					Description: "Creates a Log Group.",
					Config:      resources.LogConfig{},
				},
			},
		},
	})
}

// Resources for this plugin
var Resources = map[string]func(
	name string,
	data string,
) []byte{
	"Example::Log": api.RegisterResource(resources.ParseExampleLog),
}

// Outputs for this plugin
var Outputs = map[string]func(
	name string,
	data string,
) []byte{}

// Mappings for this plugin
var Mappings = map[string]func(
	name string,
	data string,
) []byte{}

// This function is never called, but it's required by the compiler.
func main() {}
