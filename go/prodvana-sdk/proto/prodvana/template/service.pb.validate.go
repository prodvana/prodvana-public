// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/template/service.proto

package template

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

// Validate checks the field values on ServiceTemplate with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ServiceTemplate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceTemplate with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServiceTemplateMultiError, or nil if none found.
func (m *ServiceTemplate) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceTemplate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceTemplateValidationError{
					field:  "Meta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceTemplateValidationError{
					field:  "Meta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceTemplateValidationError{
				field:  "Meta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLastUpdateTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceTemplateValidationError{
					field:  "LastUpdateTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceTemplateValidationError{
					field:  "LastUpdateTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastUpdateTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceTemplateValidationError{
				field:  "LastUpdateTimestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// skipping validation for service_config

	if len(errors) > 0 {
		return ServiceTemplateMultiError(errors)
	}

	return nil
}

// ServiceTemplateMultiError is an error wrapping multiple validation errors
// returned by ServiceTemplate.ValidateAll() if the designated constraints
// aren't met.
type ServiceTemplateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceTemplateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceTemplateMultiError) AllErrors() []error { return m }

// ServiceTemplateValidationError is the validation error returned by
// ServiceTemplate.Validate if the designated constraints aren't met.
type ServiceTemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceTemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceTemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceTemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceTemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceTemplateValidationError) ErrorName() string { return "ServiceTemplateValidationError" }

// Error satisfies the builtin error interface
func (e ServiceTemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceTemplate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceTemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceTemplateValidationError{}
