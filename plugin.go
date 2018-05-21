package main

import (
	"bitbucket.org/kablamo-dev/kombustion/plugins/kombustion-plugin-boilerplate/resources"
	"bitbucket.org/kablamo-dev/kombustion/types"
)

var Resources = map[string]types.ParserFunc{
// TODO: fill me in
}

var Outputs = map[string]types.ParserFunc{}

var Mappings = map[string]types.ParserFunc{}

var Help = types.PluginHelp{
	Description: "My plugin",
	TypeMappings: []types.TypeMapping{
		{
			Name:        "",
			Description: "",
			Config:      resources.MyConfig{},
		},
	},
}

func main() {}
