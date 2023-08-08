// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/labels/labels.proto

package labels

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on LabelDefinition with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LabelDefinition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LabelDefinition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LabelDefinitionMultiError, or nil if none found.
func (m *LabelDefinition) ValidateAll() error {
	return m.validate(true)
}

func (m *LabelDefinition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetLabel()) < 1 {
		err := LabelDefinitionValidationError{
			field:  "Label",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_LabelDefinition_Label_Pattern.MatchString(m.GetLabel()) {
		err := LabelDefinitionValidationError{
			field:  "Label",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9.\\\\-_@+]$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_LabelDefinition_Value_Pattern.MatchString(m.GetValue()) {
		err := LabelDefinitionValidationError{
			field:  "Value",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9.\\\\-_@+]$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LabelDefinitionMultiError(errors)
	}

	return nil
}

// LabelDefinitionMultiError is an error wrapping multiple validation errors
// returned by LabelDefinition.ValidateAll() if the designated constraints
// aren't met.
type LabelDefinitionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LabelDefinitionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LabelDefinitionMultiError) AllErrors() []error { return m }

// LabelDefinitionValidationError is the validation error returned by
// LabelDefinition.Validate if the designated constraints aren't met.
type LabelDefinitionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LabelDefinitionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LabelDefinitionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LabelDefinitionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LabelDefinitionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LabelDefinitionValidationError) ErrorName() string { return "LabelDefinitionValidationError" }

// Error satisfies the builtin error interface
func (e LabelDefinitionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLabelDefinition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LabelDefinitionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LabelDefinitionValidationError{}

var _LabelDefinition_Label_Pattern = regexp.MustCompile("^[a-zA-Z0-9.\\-_@+]$")

var _LabelDefinition_Value_Pattern = regexp.MustCompile("^[a-zA-Z0-9.\\-_@+]$")
