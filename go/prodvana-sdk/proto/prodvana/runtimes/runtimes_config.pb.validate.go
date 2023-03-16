// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/runtimes/runtimes_config.proto

package runtimes

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

// Validate checks the field values on K8SRuntimeInitializationConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *K8SRuntimeInitializationConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on K8SRuntimeInitializationConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// K8SRuntimeInitializationConfigMultiError, or nil if none found.
func (m *K8SRuntimeInitializationConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *K8SRuntimeInitializationConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AgentImage

	// no validation rules for AuthToken

	// no validation rules for UseResourceDefault

	// no validation rules for InteractionServerAddress

	if len(errors) > 0 {
		return K8SRuntimeInitializationConfigMultiError(errors)
	}

	return nil
}

// K8SRuntimeInitializationConfigMultiError is an error wrapping multiple
// validation errors returned by K8SRuntimeInitializationConfig.ValidateAll()
// if the designated constraints aren't met.
type K8SRuntimeInitializationConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m K8SRuntimeInitializationConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m K8SRuntimeInitializationConfigMultiError) AllErrors() []error { return m }

// K8SRuntimeInitializationConfigValidationError is the validation error
// returned by K8SRuntimeInitializationConfig.Validate if the designated
// constraints aren't met.
type K8SRuntimeInitializationConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e K8SRuntimeInitializationConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e K8SRuntimeInitializationConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e K8SRuntimeInitializationConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e K8SRuntimeInitializationConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e K8SRuntimeInitializationConfigValidationError) ErrorName() string {
	return "K8SRuntimeInitializationConfigValidationError"
}

// Error satisfies the builtin error interface
func (e K8SRuntimeInitializationConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sK8SRuntimeInitializationConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = K8SRuntimeInitializationConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = K8SRuntimeInitializationConfigValidationError{}

// Validate checks the field values on RuntimeInitializationConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RuntimeInitializationConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RuntimeInitializationConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RuntimeInitializationConfigMultiError, or nil if none found.
func (m *RuntimeInitializationConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *RuntimeInitializationConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.RuntimeSpecific.(type) {
	case *RuntimeInitializationConfig_K8S:
		if v == nil {
			err := RuntimeInitializationConfigValidationError{
				field:  "RuntimeSpecific",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetK8S()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RuntimeInitializationConfigValidationError{
						field:  "K8S",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RuntimeInitializationConfigValidationError{
						field:  "K8S",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetK8S()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RuntimeInitializationConfigValidationError{
					field:  "K8S",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return RuntimeInitializationConfigMultiError(errors)
	}

	return nil
}

// RuntimeInitializationConfigMultiError is an error wrapping multiple
// validation errors returned by RuntimeInitializationConfig.ValidateAll() if
// the designated constraints aren't met.
type RuntimeInitializationConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RuntimeInitializationConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RuntimeInitializationConfigMultiError) AllErrors() []error { return m }

// RuntimeInitializationConfigValidationError is the validation error returned
// by RuntimeInitializationConfig.Validate if the designated constraints
// aren't met.
type RuntimeInitializationConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeInitializationConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeInitializationConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeInitializationConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeInitializationConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeInitializationConfigValidationError) ErrorName() string {
	return "RuntimeInitializationConfigValidationError"
}

// Error satisfies the builtin error interface
func (e RuntimeInitializationConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntimeInitializationConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeInitializationConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeInitializationConfigValidationError{}
