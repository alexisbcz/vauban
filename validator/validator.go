/*
Copyright 2025 Alexis Bouchez <alexbcz@proton.me>

This file is part of Vauban.

Vauban is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Vauban is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with Vauban. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing options, please contact Alexis Bouchez at alexbcz@proton.me
*/
package validator

import (
	"net/url"

	"github.com/alexisbcz/vauban/validator/field"
)

// Validator represents a collection of fields to be validated
type Validator struct {
	fields []*field.Field
}

// New creates a new validator with the given fields
func New(fields ...*field.Field) *Validator {
	return &Validator{
		fields: fields,
	}
}

// AddField adds a field to the validator
func (v *Validator) AddField(field *field.Field) *Validator {
	v.fields = append(v.fields, field)
	return v
}

// AddFields adds multiple fields to the validator
func (v *Validator) AddFields(fields ...*field.Field) *Validator {
	v.fields = append(v.fields, fields...)
	return v
}

// Fields returns all fields in the validator
func (v *Validator) Fields() []*field.Field {
	return v.fields
}

type Errors = map[string][]string

// Validate validates all fields and returns if the validation succeeded and any error messages
func (v *Validator) Validate() (bool, Errors) {
	return v.ValidateValues(make(map[string]string))
}

// ValidateValues validates all fields against the provided values map and returns if validation succeeded and any error messages
func (v *Validator) ValidateValues(values map[string]string) (bool, Errors) {
	// Initialize errors map
	errors := make(Errors)
	valid := true

	// Iterate through all fields and validate with the corresponding values
	for _, f := range v.fields {
		value, _ := values[f.Key]
		fieldValid, fieldErrors := f.Validate(value)
		if !fieldValid {
			valid = false
			errors[f.Key] = fieldErrors
		}
	}

	return valid, errors
}

// ValidateForm validates form data from url.Values against the validator's fields
func (v *Validator) ValidateForm(form url.Values) (valid bool, errors Errors) {
	// Convert form values to a simple map for ValidateValues
	values := make(map[string]string)
	for _, f := range v.fields {
		formValues := form[f.Key]
		if len(formValues) > 0 {
			values[f.Key] = formValues[0]
		} else {
			values[f.Key] = ""
		}
	}

	return v.ValidateValues(values)
}

// HasErrors checks if the validation has any errors
func (v *Validator) HasErrors(errors Errors) bool {
	return len(errors) > 0
}

// GetFirstError returns the first error message for a specific field
func (v *Validator) GetFirstError(errors Errors, fieldKey string) string {
	if fieldErrors, ok := errors[fieldKey]; ok && len(fieldErrors) > 0 {
		return fieldErrors[0]
	}
	return ""
}

// GetAllErrors returns all error messages across all fields as a flat slice
func (v *Validator) GetAllErrors(errors Errors) []string {
	var allErrors []string
	for _, fieldErrors := range errors {
		allErrors = append(allErrors, fieldErrors...)
	}
	return allErrors
}
