// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/common_config/parameters.proto

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

// Validate checks the field values on StringParameterDefinition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StringParameterDefinition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StringParameterDefinition with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StringParameterDefinitionMultiError, or nil if none found.
func (m *StringParameterDefinition) ValidateAll() error {
	return m.validate(true)
}

func (m *StringParameterDefinition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for InitialValue

	if len(errors) > 0 {
		return StringParameterDefinitionMultiError(errors)
	}

	return nil
}

// StringParameterDefinitionMultiError is an error wrapping multiple validation
// errors returned by StringParameterDefinition.ValidateAll() if the
// designated constraints aren't met.
type StringParameterDefinitionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StringParameterDefinitionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StringParameterDefinitionMultiError) AllErrors() []error { return m }

// StringParameterDefinitionValidationError is the validation error returned by
// StringParameterDefinition.Validate if the designated constraints aren't met.
type StringParameterDefinitionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringParameterDefinitionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringParameterDefinitionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringParameterDefinitionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringParameterDefinitionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringParameterDefinitionValidationError) ErrorName() string {
	return "StringParameterDefinitionValidationError"
}

// Error satisfies the builtin error interface
func (e StringParameterDefinitionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringParameterDefinition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringParameterDefinitionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringParameterDefinitionValidationError{}

// Validate checks the field values on DockerImageParameterDefinition with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DockerImageParameterDefinition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DockerImageParameterDefinition with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DockerImageParameterDefinitionMultiError, or nil if none found.
func (m *DockerImageParameterDefinition) ValidateAll() error {
	return m.validate(true)
}

func (m *DockerImageParameterDefinition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetInitialValue()) < 1 {
		err := DockerImageParameterDefinitionValidationError{
			field:  "InitialValue",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetImageRegistryInfo() == nil {
		err := DockerImageParameterDefinitionValidationError{
			field:  "ImageRegistryInfo",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetImageRegistryInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DockerImageParameterDefinitionValidationError{
					field:  "ImageRegistryInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DockerImageParameterDefinitionValidationError{
					field:  "ImageRegistryInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetImageRegistryInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DockerImageParameterDefinitionValidationError{
				field:  "ImageRegistryInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DockerImageParameterDefinitionMultiError(errors)
	}

	return nil
}

// DockerImageParameterDefinitionMultiError is an error wrapping multiple
// validation errors returned by DockerImageParameterDefinition.ValidateAll()
// if the designated constraints aren't met.
type DockerImageParameterDefinitionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DockerImageParameterDefinitionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DockerImageParameterDefinitionMultiError) AllErrors() []error { return m }

// DockerImageParameterDefinitionValidationError is the validation error
// returned by DockerImageParameterDefinition.Validate if the designated
// constraints aren't met.
type DockerImageParameterDefinitionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DockerImageParameterDefinitionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DockerImageParameterDefinitionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DockerImageParameterDefinitionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DockerImageParameterDefinitionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DockerImageParameterDefinitionValidationError) ErrorName() string {
	return "DockerImageParameterDefinitionValidationError"
}

// Error satisfies the builtin error interface
func (e DockerImageParameterDefinitionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDockerImageParameterDefinition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DockerImageParameterDefinitionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DockerImageParameterDefinitionValidationError{}

// Validate checks the field values on IntParameterDefinition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntParameterDefinition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntParameterDefinition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntParameterDefinitionMultiError, or nil if none found.
func (m *IntParameterDefinition) ValidateAll() error {
	return m.validate(true)
}

func (m *IntParameterDefinition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for InitialValue

	if len(errors) > 0 {
		return IntParameterDefinitionMultiError(errors)
	}

	return nil
}

// IntParameterDefinitionMultiError is an error wrapping multiple validation
// errors returned by IntParameterDefinition.ValidateAll() if the designated
// constraints aren't met.
type IntParameterDefinitionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntParameterDefinitionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntParameterDefinitionMultiError) AllErrors() []error { return m }

// IntParameterDefinitionValidationError is the validation error returned by
// IntParameterDefinition.Validate if the designated constraints aren't met.
type IntParameterDefinitionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntParameterDefinitionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntParameterDefinitionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntParameterDefinitionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntParameterDefinitionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntParameterDefinitionValidationError) ErrorName() string {
	return "IntParameterDefinitionValidationError"
}

// Error satisfies the builtin error interface
func (e IntParameterDefinitionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntParameterDefinition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntParameterDefinitionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntParameterDefinitionValidationError{}

// Validate checks the field values on ParameterDefinition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ParameterDefinition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ParameterDefinition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ParameterDefinitionMultiError, or nil if none found.
func (m *ParameterDefinition) ValidateAll() error {
	return m.validate(true)
}

func (m *ParameterDefinition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := ParameterDefinitionValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	oneofConfigOneofPresent := false
	switch v := m.ConfigOneof.(type) {
	case *ParameterDefinition_String_:
		if v == nil {
			err := ParameterDefinitionValidationError{
				field:  "ConfigOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConfigOneofPresent = true

		if all {
			switch v := interface{}(m.GetString_()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ParameterDefinitionValidationError{
						field:  "String_",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ParameterDefinitionValidationError{
						field:  "String_",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetString_()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ParameterDefinitionValidationError{
					field:  "String_",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ParameterDefinition_DockerImage:
		if v == nil {
			err := ParameterDefinitionValidationError{
				field:  "ConfigOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConfigOneofPresent = true

		if all {
			switch v := interface{}(m.GetDockerImage()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ParameterDefinitionValidationError{
						field:  "DockerImage",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ParameterDefinitionValidationError{
						field:  "DockerImage",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDockerImage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ParameterDefinitionValidationError{
					field:  "DockerImage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ParameterDefinition_Int:
		if v == nil {
			err := ParameterDefinitionValidationError{
				field:  "ConfigOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConfigOneofPresent = true

		if all {
			switch v := interface{}(m.GetInt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ParameterDefinitionValidationError{
						field:  "Int",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ParameterDefinitionValidationError{
						field:  "Int",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ParameterDefinitionValidationError{
					field:  "Int",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofConfigOneofPresent {
		err := ParameterDefinitionValidationError{
			field:  "ConfigOneof",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ParameterDefinitionMultiError(errors)
	}

	return nil
}

// ParameterDefinitionMultiError is an error wrapping multiple validation
// errors returned by ParameterDefinition.ValidateAll() if the designated
// constraints aren't met.
type ParameterDefinitionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ParameterDefinitionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ParameterDefinitionMultiError) AllErrors() []error { return m }

// ParameterDefinitionValidationError is the validation error returned by
// ParameterDefinition.Validate if the designated constraints aren't met.
type ParameterDefinitionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParameterDefinitionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParameterDefinitionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParameterDefinitionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParameterDefinitionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParameterDefinitionValidationError) ErrorName() string {
	return "ParameterDefinitionValidationError"
}

// Error satisfies the builtin error interface
func (e ParameterDefinitionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParameterDefinition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParameterDefinitionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParameterDefinitionValidationError{}

// Validate checks the field values on ParameterValue with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ParameterValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ParameterValue with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ParameterValueMultiError,
// or nil if none found.
func (m *ParameterValue) ValidateAll() error {
	return m.validate(true)
}

func (m *ParameterValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := ParameterValueValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	oneofValueOneofPresent := false
	switch v := m.ValueOneof.(type) {
	case *ParameterValue_String_:
		if v == nil {
			err := ParameterValueValidationError{
				field:  "ValueOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValueOneofPresent = true
		// no validation rules for String_
	case *ParameterValue_DockerImage:
		if v == nil {
			err := ParameterValueValidationError{
				field:  "ValueOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValueOneofPresent = true
		// no validation rules for DockerImage
	case *ParameterValue_Int:
		if v == nil {
			err := ParameterValueValidationError{
				field:  "ValueOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValueOneofPresent = true
		// no validation rules for Int
	default:
		_ = v // ensures v is used
	}
	if !oneofValueOneofPresent {
		err := ParameterValueValidationError{
			field:  "ValueOneof",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ParameterValueMultiError(errors)
	}

	return nil
}

// ParameterValueMultiError is an error wrapping multiple validation errors
// returned by ParameterValue.ValidateAll() if the designated constraints
// aren't met.
type ParameterValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ParameterValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ParameterValueMultiError) AllErrors() []error { return m }

// ParameterValueValidationError is the validation error returned by
// ParameterValue.Validate if the designated constraints aren't met.
type ParameterValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParameterValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParameterValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParameterValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParameterValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParameterValueValidationError) ErrorName() string { return "ParameterValueValidationError" }

// Error satisfies the builtin error interface
func (e ParameterValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParameterValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParameterValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParameterValueValidationError{}

// Validate checks the field values on ParametersConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ParametersConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ParametersConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ParametersConfigMultiError, or nil if none found.
func (m *ParametersConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ParametersConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetParameters() {
		_, _ = idx, item

		if item == nil {
			err := ParametersConfigValidationError{
				field:  fmt.Sprintf("Parameters[%v]", idx),
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ParametersConfigValidationError{
						field:  fmt.Sprintf("Parameters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ParametersConfigValidationError{
						field:  fmt.Sprintf("Parameters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ParametersConfigValidationError{
					field:  fmt.Sprintf("Parameters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ParametersConfigMultiError(errors)
	}

	return nil
}

// ParametersConfigMultiError is an error wrapping multiple validation errors
// returned by ParametersConfig.ValidateAll() if the designated constraints
// aren't met.
type ParametersConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ParametersConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ParametersConfigMultiError) AllErrors() []error { return m }

// ParametersConfigValidationError is the validation error returned by
// ParametersConfig.Validate if the designated constraints aren't met.
type ParametersConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParametersConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParametersConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParametersConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParametersConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParametersConfigValidationError) ErrorName() string { return "ParametersConfigValidationError" }

// Error satisfies the builtin error interface
func (e ParametersConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParametersConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParametersConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParametersConfigValidationError{}

// Validate checks the field values on Template with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Template) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Template with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TemplateMultiError, or nil
// if none found.
func (m *Template) ValidateAll() error {
	return m.validate(true)
}

func (m *Template) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTemplate()) < 1 {
		err := TemplateValidationError{
			field:  "Template",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetParameters() {
		_, _ = idx, item

		if item == nil {
			err := TemplateValidationError{
				field:  fmt.Sprintf("Parameters[%v]", idx),
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TemplateValidationError{
						field:  fmt.Sprintf("Parameters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TemplateValidationError{
						field:  fmt.Sprintf("Parameters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TemplateValidationError{
					field:  fmt.Sprintf("Parameters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TemplateMultiError(errors)
	}

	return nil
}

// TemplateMultiError is an error wrapping multiple validation errors returned
// by Template.ValidateAll() if the designated constraints aren't met.
type TemplateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TemplateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TemplateMultiError) AllErrors() []error { return m }

// TemplateValidationError is the validation error returned by
// Template.Validate if the designated constraints aren't met.
type TemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TemplateValidationError) ErrorName() string { return "TemplateValidationError" }

// Error satisfies the builtin error interface
func (e TemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTemplate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TemplateValidationError{}
