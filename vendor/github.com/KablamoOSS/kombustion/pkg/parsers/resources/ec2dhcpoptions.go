package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2DHCPOptions Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-dhcp-options.html
type EC2DHCPOptions struct {
	Type       string                   `yaml:"Type"`
	Properties EC2DHCPOptionsProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// EC2DHCPOptions Properties
type EC2DHCPOptionsProperties struct {
	DomainName         interface{} `yaml:"DomainName,omitempty"`
	NetbiosNodeType    interface{} `yaml:"NetbiosNodeType,omitempty"`
	DomainNameServers  interface{} `yaml:"DomainNameServers,omitempty"`
	NetbiosNameServers interface{} `yaml:"NetbiosNameServers,omitempty"`
	NtpServers         interface{} `yaml:"NtpServers,omitempty"`
	Tags               interface{} `yaml:"Tags,omitempty"`
}

// NewEC2DHCPOptions constructor creates a new EC2DHCPOptions
func NewEC2DHCPOptions(properties EC2DHCPOptionsProperties, deps ...interface{}) EC2DHCPOptions {
	return EC2DHCPOptions{
		Type:       "AWS::EC2::DHCPOptions",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2DHCPOptions parses EC2DHCPOptions
func ParseEC2DHCPOptions(name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2DHCPOptions
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2DHCPOptions - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEC2DHCPOptions validator
func (resource EC2DHCPOptions) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2DHCPOptionsProperties validator
func (resource EC2DHCPOptionsProperties) Validate() []error {
	errs := []error{}
	return errs
}
