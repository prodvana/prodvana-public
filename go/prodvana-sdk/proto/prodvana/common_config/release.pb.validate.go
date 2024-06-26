// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/common_config/release.proto

package common_config

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

// Validate checks the field values on ReleaseSettings with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ReleaseSettings) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReleaseSettings with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReleaseSettingsMultiError, or nil if none found.
func (m *ReleaseSettings) ValidateAll() error {
	return m.validate(true)
}

func (m *ReleaseSettings) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBypassSettings()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReleaseSettingsValidationError{
					field:  "BypassSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReleaseSettingsValidationError{
					field:  "BypassSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBypassSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReleaseSettingsValidationError{
				field:  "BypassSettings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ReleaseSettingsMultiError(errors)
	}

	return nil
}

// ReleaseSettingsMultiError is an error wrapping multiple validation errors
// returned by ReleaseSettings.ValidateAll() if the designated constraints
// aren't met.
type ReleaseSettingsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReleaseSettingsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReleaseSettingsMultiError) AllErrors() []error { return m }

// ReleaseSettingsValidationError is the validation error returned by
// ReleaseSettings.Validate if the designated constraints aren't met.
type ReleaseSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReleaseSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReleaseSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReleaseSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReleaseSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReleaseSettingsValidationError) ErrorName() string { return "ReleaseSettingsValidationError" }

// Error satisfies the builtin error interface
func (e ReleaseSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReleaseSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReleaseSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReleaseSettingsValidationError{}

// Validate checks the field values on ReleaseSettings_BypassSettings with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReleaseSettings_BypassSettings) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReleaseSettings_BypassSettings with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ReleaseSettings_BypassSettingsMultiError, or nil if none found.
func (m *ReleaseSettings_BypassSettings) ValidateAll() error {
	return m.validate(true)
}

func (m *ReleaseSettings_BypassSettings) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NoBypassProtections

	// no validation rules for NoBypassConvergenceExtensions

	// no validation rules for NoBypassApprovals

	// no validation rules for NoBypassReleaseChannelOrdering

	if len(errors) > 0 {
		return ReleaseSettings_BypassSettingsMultiError(errors)
	}

	return nil
}

// ReleaseSettings_BypassSettingsMultiError is an error wrapping multiple
// validation errors returned by ReleaseSettings_BypassSettings.ValidateAll()
// if the designated constraints aren't met.
type ReleaseSettings_BypassSettingsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReleaseSettings_BypassSettingsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReleaseSettings_BypassSettingsMultiError) AllErrors() []error { return m }

// ReleaseSettings_BypassSettingsValidationError is the validation error
// returned by ReleaseSettings_BypassSettings.Validate if the designated
// constraints aren't met.
type ReleaseSettings_BypassSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReleaseSettings_BypassSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReleaseSettings_BypassSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReleaseSettings_BypassSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReleaseSettings_BypassSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReleaseSettings_BypassSettingsValidationError) ErrorName() string {
	return "ReleaseSettings_BypassSettingsValidationError"
}

// Error satisfies the builtin error interface
func (e ReleaseSettings_BypassSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReleaseSettings_BypassSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReleaseSettings_BypassSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReleaseSettings_BypassSettingsValidationError{}
