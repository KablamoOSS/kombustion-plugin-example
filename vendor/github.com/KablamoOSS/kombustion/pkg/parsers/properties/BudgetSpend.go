package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BudgetSpend Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-spend.html
type BudgetSpend struct {
	Amount interface{} `yaml:"Amount"`
	Unit   interface{} `yaml:"Unit"`
}

// BudgetSpend validation
func (resource BudgetSpend) Validate() []error {
	errors := []error{}

	if resource.Amount == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Amount'"))
	}
	if resource.Unit == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Unit'"))
	}
	return errors
}
