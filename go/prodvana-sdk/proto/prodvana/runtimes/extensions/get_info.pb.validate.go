// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/runtimes/extensions/get_info.proto

package extensions

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

// Validate checks the field values on GetInfoOutput with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetInfoOutput) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetInfoOutput with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetInfoOutputMultiError, or
// nil if none found.
func (m *GetInfoOutput) ValidateAll() error {
	return m.validate(true)
}

func (m *GetInfoOutput) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetOutputs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetInfoOutputValidationError{
						field:  fmt.Sprintf("Outputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetInfoOutputValidationError{
						field:  fmt.Sprintf("Outputs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetInfoOutputValidationError{
					field:  fmt.Sprintf("Outputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetInfoOutputMultiError(errors)
	}

	return nil
}

// GetInfoOutputMultiError is an error wrapping multiple validation errors
// returned by GetInfoOutput.ValidateAll() if the designated constraints
// aren't met.
type GetInfoOutputMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetInfoOutputMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetInfoOutputMultiError) AllErrors() []error { return m }

// GetInfoOutputValidationError is the validation error returned by
// GetInfoOutput.Validate if the designated constraints aren't met.
type GetInfoOutputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoOutputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoOutputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoOutputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoOutputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoOutputValidationError) ErrorName() string { return "GetInfoOutputValidationError" }

// Error satisfies the builtin error interface
func (e GetInfoOutputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoOutput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoOutputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoOutputValidationError{}

// Validate checks the field values on OutputContent with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OutputContent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OutputContent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OutputContentMultiError, or
// nil if none found.
func (m *OutputContent) ValidateAll() error {
	return m.validate(true)
}

func (m *OutputContent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Text

	if len(errors) > 0 {
		return OutputContentMultiError(errors)
	}

	return nil
}

// OutputContentMultiError is an error wrapping multiple validation errors
// returned by OutputContent.ValidateAll() if the designated constraints
// aren't met.
type OutputContentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OutputContentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OutputContentMultiError) AllErrors() []error { return m }

// OutputContentValidationError is the validation error returned by
// OutputContent.Validate if the designated constraints aren't met.
type OutputContentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutputContentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutputContentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutputContentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutputContentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutputContentValidationError) ErrorName() string { return "OutputContentValidationError" }

// Error satisfies the builtin error interface
func (e OutputContentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutputContent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutputContentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutputContentValidationError{}
