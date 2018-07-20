package main

import (
	"github.com/KablamoOSS/kombustion-plugin-example/parsers"
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
		name = "kombustion-plugin-example"
	}
}

// Register plugin
func Register() []byte {
	return api.RegisterPlugin(types.Config{
		Name:    name,
		Version: version,
		Prefix:  "Kablamo::Example",
		Help: types.Help{
			Description: "An example plugin to create a log group",
			Types: []types.TypeMapping{
				{
					Name:        "Log",
					Description: "Creates a Log Group.",
					Config:      parsers.LogConfig{},
				},
			},
		},
	})
}

// Parsers for this plugin
var Parsers = map[string]func(
	name string,
	data string,
) []byte{
	"Log": api.RegisterParser(parsers.ParseExampleLog),
}

// This function is never called, but it's required by the compiler.
func main() {}
