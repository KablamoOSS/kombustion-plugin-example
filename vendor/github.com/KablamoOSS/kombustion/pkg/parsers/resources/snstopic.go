package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// SNSTopic Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html
type SNSTopic struct {
	Type       string             `yaml:"Type"`
	Properties SNSTopicProperties `yaml:"Properties"`
	Condition  interface{}        `yaml:"Condition,omitempty"`
	Metadata   interface{}        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}        `yaml:"DependsOn,omitempty"`
}

// SNSTopic Properties
type SNSTopicProperties struct {
	DisplayName  interface{} `yaml:"DisplayName,omitempty"`
	TopicName    interface{} `yaml:"TopicName,omitempty"`
	Subscription interface{} `yaml:"Subscription,omitempty"`
}

// NewSNSTopic constructor creates a new SNSTopic
func NewSNSTopic(properties SNSTopicProperties, deps ...interface{}) SNSTopic {
	return SNSTopic{
		Type:       "AWS::SNS::Topic",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSNSTopic parses SNSTopic
func ParseSNSTopic(name string, data string) (cf types.TemplateObject, err error) {
	var resource SNSTopic
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SNSTopic - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseSNSTopic validator
func (resource SNSTopic) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSNSTopicProperties validator
func (resource SNSTopicProperties) Validate() []error {
	errs := []error{}
	return errs
}
