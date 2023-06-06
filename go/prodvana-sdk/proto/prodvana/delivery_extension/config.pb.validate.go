// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/delivery_extension/config.proto

package delivery_extension

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

// Validate checks the field values on DeliveryExtensionConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeliveryExtensionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeliveryExtensionConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeliveryExtensionConfigMultiError, or nil if none found.
func (m *DeliveryExtensionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DeliveryExtensionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 0 || l > 63 {
		err := DeliveryExtensionConfigValidationError{
			field:  "Name",
			reason: "value length must be between 0 and 63 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_DeliveryExtensionConfig_Name_Pattern.MatchString(m.GetName()) {
		err := DeliveryExtensionConfigValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetParameters() {
		_, _ = idx, item

		if item == nil {
			err := DeliveryExtensionConfigValidationError{
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
					errors = append(errors, DeliveryExtensionConfigValidationError{
						field:  fmt.Sprintf("Parameters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DeliveryExtensionConfigValidationError{
						field:  fmt.Sprintf("Parameters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeliveryExtensionConfigValidationError{
					field:  fmt.Sprintf("Parameters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	oneofExecConfigPresent := false
	switch v := m.ExecConfig.(type) {
	case *DeliveryExtensionConfig_TaskConfig:
		if v == nil {
			err := DeliveryExtensionConfigValidationError{
				field:  "ExecConfig",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofExecConfigPresent = true

		if all {
			switch v := interface{}(m.GetTaskConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DeliveryExtensionConfigValidationError{
						field:  "TaskConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DeliveryExtensionConfigValidationError{
						field:  "TaskConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTaskConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeliveryExtensionConfigValidationError{
					field:  "TaskConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DeliveryExtensionConfig_KubernetesConfig:
		if v == nil {
			err := DeliveryExtensionConfigValidationError{
				field:  "ExecConfig",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofExecConfigPresent = true

		if all {
			switch v := interface{}(m.GetKubernetesConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DeliveryExtensionConfigValidationError{
						field:  "KubernetesConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DeliveryExtensionConfigValidationError{
						field:  "KubernetesConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetKubernetesConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeliveryExtensionConfigValidationError{
					field:  "KubernetesConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofExecConfigPresent {
		err := DeliveryExtensionConfigValidationError{
			field:  "ExecConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeliveryExtensionConfigMultiError(errors)
	}

	return nil
}

// DeliveryExtensionConfigMultiError is an error wrapping multiple validation
// errors returned by DeliveryExtensionConfig.ValidateAll() if the designated
// constraints aren't met.
type DeliveryExtensionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeliveryExtensionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeliveryExtensionConfigMultiError) AllErrors() []error { return m }

// DeliveryExtensionConfigValidationError is the validation error returned by
// DeliveryExtensionConfig.Validate if the designated constraints aren't met.
type DeliveryExtensionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeliveryExtensionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeliveryExtensionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeliveryExtensionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeliveryExtensionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeliveryExtensionConfigValidationError) ErrorName() string {
	return "DeliveryExtensionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e DeliveryExtensionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeliveryExtensionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeliveryExtensionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeliveryExtensionConfigValidationError{}

var _DeliveryExtensionConfig_Name_Pattern = regexp.MustCompile("^[a-z]?([a-z0-9-]*[a-z0-9]){0,1}$")

// Validate checks the field values on DeliveryExtensionInstanceConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeliveryExtensionInstanceConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeliveryExtensionInstanceConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DeliveryExtensionInstanceConfigMultiError, or nil if none found.
func (m *DeliveryExtensionInstanceConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DeliveryExtensionInstanceConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 63 {
		err := DeliveryExtensionInstanceConfigValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 63 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_DeliveryExtensionInstanceConfig_Name_Pattern.MatchString(m.GetName()) {
		err := DeliveryExtensionInstanceConfigValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetParameters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DeliveryExtensionInstanceConfigValidationError{
						field:  fmt.Sprintf("Parameters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DeliveryExtensionInstanceConfigValidationError{
						field:  fmt.Sprintf("Parameters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeliveryExtensionInstanceConfigValidationError{
					field:  fmt.Sprintf("Parameters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DeliveryExtensionInstanceConfigMultiError(errors)
	}

	return nil
}

// DeliveryExtensionInstanceConfigMultiError is an error wrapping multiple
// validation errors returned by DeliveryExtensionInstanceConfig.ValidateAll()
// if the designated constraints aren't met.
type DeliveryExtensionInstanceConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeliveryExtensionInstanceConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeliveryExtensionInstanceConfigMultiError) AllErrors() []error { return m }

// DeliveryExtensionInstanceConfigValidationError is the validation error
// returned by DeliveryExtensionInstanceConfig.Validate if the designated
// constraints aren't met.
type DeliveryExtensionInstanceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeliveryExtensionInstanceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeliveryExtensionInstanceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeliveryExtensionInstanceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeliveryExtensionInstanceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeliveryExtensionInstanceConfigValidationError) ErrorName() string {
	return "DeliveryExtensionInstanceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e DeliveryExtensionInstanceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeliveryExtensionInstanceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeliveryExtensionInstanceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeliveryExtensionInstanceConfigValidationError{}

var _DeliveryExtensionInstanceConfig_Name_Pattern = regexp.MustCompile("^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$")

// Validate checks the field values on CompiledDeliveryExtensionInstanceConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *CompiledDeliveryExtensionInstanceConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CompiledDeliveryExtensionInstanceConfig with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// CompiledDeliveryExtensionInstanceConfigMultiError, or nil if none found.
func (m *CompiledDeliveryExtensionInstanceConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *CompiledDeliveryExtensionInstanceConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDefinition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompiledDeliveryExtensionInstanceConfigValidationError{
					field:  "Definition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompiledDeliveryExtensionInstanceConfigValidationError{
					field:  "Definition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDefinition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompiledDeliveryExtensionInstanceConfigValidationError{
				field:  "Definition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRuntimeExecution()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompiledDeliveryExtensionInstanceConfigValidationError{
					field:  "RuntimeExecution",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompiledDeliveryExtensionInstanceConfigValidationError{
					field:  "RuntimeExecution",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRuntimeExecution()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompiledDeliveryExtensionInstanceConfigValidationError{
				field:  "RuntimeExecution",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	{
		sorted_keys := make([]string, len(m.GetEnv()))
		i := 0
		for key := range m.GetEnv() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetEnv()[key]
			_ = val

			if val == nil {
				err := CompiledDeliveryExtensionInstanceConfigValidationError{
					field:  fmt.Sprintf("Env[%v]", key),
					reason: "value cannot be sparse, all pairs must be non-nil",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if !_CompiledDeliveryExtensionInstanceConfig_Env_Pattern.MatchString(key) {
				err := CompiledDeliveryExtensionInstanceConfigValidationError{
					field:  fmt.Sprintf("Env[%v]", key),
					reason: "value does not match regex pattern \"^[a-zA-Z_]+[a-zA-Z0-9_]*$\"",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, CompiledDeliveryExtensionInstanceConfigValidationError{
							field:  fmt.Sprintf("Env[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, CompiledDeliveryExtensionInstanceConfigValidationError{
							field:  fmt.Sprintf("Env[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return CompiledDeliveryExtensionInstanceConfigValidationError{
						field:  fmt.Sprintf("Env[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	for idx, item := range m.GetParameterValues() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CompiledDeliveryExtensionInstanceConfigValidationError{
						field:  fmt.Sprintf("ParameterValues[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CompiledDeliveryExtensionInstanceConfigValidationError{
						field:  fmt.Sprintf("ParameterValues[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CompiledDeliveryExtensionInstanceConfigValidationError{
					field:  fmt.Sprintf("ParameterValues[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CompiledDeliveryExtensionInstanceConfigMultiError(errors)
	}

	return nil
}

// CompiledDeliveryExtensionInstanceConfigMultiError is an error wrapping
// multiple validation errors returned by
// CompiledDeliveryExtensionInstanceConfig.ValidateAll() if the designated
// constraints aren't met.
type CompiledDeliveryExtensionInstanceConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompiledDeliveryExtensionInstanceConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompiledDeliveryExtensionInstanceConfigMultiError) AllErrors() []error { return m }

// CompiledDeliveryExtensionInstanceConfigValidationError is the validation
// error returned by CompiledDeliveryExtensionInstanceConfig.Validate if the
// designated constraints aren't met.
type CompiledDeliveryExtensionInstanceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompiledDeliveryExtensionInstanceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompiledDeliveryExtensionInstanceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompiledDeliveryExtensionInstanceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompiledDeliveryExtensionInstanceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompiledDeliveryExtensionInstanceConfigValidationError) ErrorName() string {
	return "CompiledDeliveryExtensionInstanceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e CompiledDeliveryExtensionInstanceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompiledDeliveryExtensionInstanceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompiledDeliveryExtensionInstanceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompiledDeliveryExtensionInstanceConfigValidationError{}

var _CompiledDeliveryExtensionInstanceConfig_Env_Pattern = regexp.MustCompile("^[a-zA-Z_]+[a-zA-Z0-9_]*$")
