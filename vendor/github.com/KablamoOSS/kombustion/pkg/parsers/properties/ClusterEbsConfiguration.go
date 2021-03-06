package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClusterEbsConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-ebsconfiguration.html
type ClusterEbsConfiguration struct {
	EbsOptimized          interface{} `yaml:"EbsOptimized,omitempty"`
	EbsBlockDeviceConfigs interface{} `yaml:"EbsBlockDeviceConfigs,omitempty"`
}

// ClusterEbsConfiguration validation
func (resource ClusterEbsConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
