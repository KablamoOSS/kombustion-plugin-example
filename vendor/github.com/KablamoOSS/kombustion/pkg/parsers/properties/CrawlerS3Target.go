package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// CrawlerS3Target Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-s3target.html
type CrawlerS3Target struct {
	Path       interface{} `yaml:"Path,omitempty"`
	Exclusions interface{} `yaml:"Exclusions,omitempty"`
}

// CrawlerS3Target validation
func (resource CrawlerS3Target) Validate() []error {
	errors := []error{}

	return errors
}
