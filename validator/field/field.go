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
package field

import (
	"fmt"
	"net/mail"
)

type RuleChecker = func(value string) (ok bool)

type Rule struct {
	Error   string
	Checker RuleChecker
}

type Field struct {
	Key   string
	rules []Rule
}

func (f *Field) appendRule(defaultMessage string, messages []string, checker RuleChecker) {
	message := defaultMessage
	if len(messages) > 0 {
		message = messages[0]
	}
	f.rules = append(f.rules, Rule{
		Error:   message,
		Checker: checker,
	})
}

func New(key string) *Field {
	return &Field{
		Key: key,
	}
}

// Email validates that the field value is a valid email address
func (f *Field) Email(messages ...string) *Field {
	f.appendRule("should be a valid email address", messages, func(value string) (ok bool) {
		_, err := mail.ParseAddress(value)
		return err == nil
	})
	return f
}

// Unique checks if the field value is unique according to the provided checker function
func (f *Field) Unique(checker func(value string) (ok bool), messages ...string) *Field {
	defaultMessage := "should be unique"
	f.appendRule(defaultMessage, messages, checker)
	return f
}

// MinLength validates that the field value has at least the specified minimum length
func (f *Field) MinLength(minLength int, messages ...string) *Field {
	defaultMessage := fmt.Sprintf("should have a minimum length of %d", minLength)
	f.appendRule(defaultMessage, messages, func(value string) (ok bool) {
		return len(value) >= minLength
	})
	return f
}

// MaxLength validates that the field value doesn't exceed the specified maximum length
func (f *Field) MaxLength(maxLength int, messages ...string) *Field {
	defaultMessage := fmt.Sprintf("should have a maximum length of %d", maxLength)
	f.appendRule(defaultMessage, messages, func(value string) (ok bool) {
		return len(value) <= maxLength
	})
	return f
}

// Validate runs all the validation rules associated with the field
func (f *Field) Validate(value string) (bool, []string) {
	errors := []string{}

	for _, rule := range f.rules {
		if !rule.Checker(value) {
			errors = append(errors, rule.Error)
		}
	}

	return len(errors) == 0, errors
}
