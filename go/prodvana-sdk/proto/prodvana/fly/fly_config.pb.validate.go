// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/fly/fly_config.proto

package fly

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

// Validate checks the field values on FlyConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FlyConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FlyConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FlyConfigMultiError, or nil
// if none found.
func (m *FlyConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *FlyConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.TomlOneof.(type) {
	case *FlyConfig_Inlined:
		if v == nil {
			err := FlyConfigValidationError{
				field:  "TomlOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if utf8.RuneCountInString(m.GetInlined()) < 1 {
			err := FlyConfigValidationError{
				field:  "Inlined",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *FlyConfig_Local:
		if v == nil {
			err := FlyConfigValidationError{
				field:  "TomlOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if m.GetLocal() == nil {
			err := FlyConfigValidationError{
				field:  "Local",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetLocal()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FlyConfigValidationError{
						field:  "Local",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FlyConfigValidationError{
						field:  "Local",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLocal()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FlyConfigValidationError{
					field:  "Local",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *FlyConfig_Remote:
		if v == nil {
			err := FlyConfigValidationError{
				field:  "TomlOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if m.GetRemote() == nil {
			err := FlyConfigValidationError{
				field:  "Remote",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetRemote()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FlyConfigValidationError{
						field:  "Remote",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FlyConfigValidationError{
						field:  "Remote",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRemote()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FlyConfigValidationError{
					field:  "Remote",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return FlyConfigMultiError(errors)
	}

	return nil
}

// FlyConfigMultiError is an error wrapping multiple validation errors returned
// by FlyConfig.ValidateAll() if the designated constraints aren't met.
type FlyConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FlyConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FlyConfigMultiError) AllErrors() []error { return m }

// FlyConfigValidationError is the validation error returned by
// FlyConfig.Validate if the designated constraints aren't met.
type FlyConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FlyConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FlyConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FlyConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FlyConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FlyConfigValidationError) ErrorName() string { return "FlyConfigValidationError" }

// Error satisfies the builtin error interface
func (e FlyConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFlyConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FlyConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FlyConfigValidationError{}
