package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DatabaseDatabaseInput Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html
type DatabaseDatabaseInput struct {
	Description interface{} `yaml:"Description,omitempty"`
	LocationUri interface{} `yaml:"LocationUri,omitempty"`
	Name        interface{} `yaml:"Name,omitempty"`
	Parameters  interface{} `yaml:"Parameters,omitempty"`
}

// DatabaseDatabaseInput validation
func (resource DatabaseDatabaseInput) Validate() []error {
	errors := []error{}

	return errors
}
