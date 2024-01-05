// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/protection/protection_config.proto

package protection

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

// Validate checks the field values on ProtectionConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProtectionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtectionConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtectionConfigMultiError, or nil if none found.
func (m *ProtectionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtectionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 63 {
		err := ProtectionConfigValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 63 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_ProtectionConfig_Name_Pattern.MatchString(m.GetName()) {
		err := ProtectionConfigValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPollInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProtectionConfigValidationError{
					field:  "PollInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProtectionConfigValidationError{
					field:  "PollInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPollInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProtectionConfigValidationError{
				field:  "PollInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProtectionConfigValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProtectionConfigValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProtectionConfigValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetParameters() {
		_, _ = idx, item

		if item == nil {
			err := ProtectionConfigValidationError{
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
					errors = append(errors, ProtectionConfigValidationError{
						field:  fmt.Sprintf("Parameters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProtectionConfigValidationError{
						field:  fmt.Sprintf("Parameters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtectionConfigValidationError{
					field:  fmt.Sprintf("Parameters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
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
				err := ProtectionConfigValidationError{
					field:  fmt.Sprintf("Env[%v]", key),
					reason: "value cannot be sparse, all pairs must be non-nil",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if !_ProtectionConfig_Env_Pattern.MatchString(key) {
				err := ProtectionConfigValidationError{
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
						errors = append(errors, ProtectionConfigValidationError{
							field:  fmt.Sprintf("Env[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, ProtectionConfigValidationError{
							field:  fmt.Sprintf("Env[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return ProtectionConfigValidationError{
						field:  fmt.Sprintf("Env[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	oneofExecConfigPresent := false
	switch v := m.ExecConfig.(type) {
	case *ProtectionConfig_TaskConfig:
		if v == nil {
			err := ProtectionConfigValidationError{
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
					errors = append(errors, ProtectionConfigValidationError{
						field:  "TaskConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProtectionConfigValidationError{
						field:  "TaskConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTaskConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtectionConfigValidationError{
					field:  "TaskConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProtectionConfig_KubernetesConfig:
		if v == nil {
			err := ProtectionConfigValidationError{
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
					errors = append(errors, ProtectionConfigValidationError{
						field:  "KubernetesConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProtectionConfigValidationError{
						field:  "KubernetesConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetKubernetesConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtectionConfigValidationError{
					field:  "KubernetesConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProtectionConfig_CommitDenylist:
		if v == nil {
			err := ProtectionConfigValidationError{
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
			switch v := interface{}(m.GetCommitDenylist()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProtectionConfigValidationError{
						field:  "CommitDenylist",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProtectionConfigValidationError{
						field:  "CommitDenylist",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCommitDenylist()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtectionConfigValidationError{
					field:  "CommitDenylist",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofExecConfigPresent {
		err := ProtectionConfigValidationError{
			field:  "ExecConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ProtectionConfigMultiError(errors)
	}

	return nil
}

// ProtectionConfigMultiError is an error wrapping multiple validation errors
// returned by ProtectionConfig.ValidateAll() if the designated constraints
// aren't met.
type ProtectionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtectionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtectionConfigMultiError) AllErrors() []error { return m }

// ProtectionConfigValidationError is the validation error returned by
// ProtectionConfig.Validate if the designated constraints aren't met.
type ProtectionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtectionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtectionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtectionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtectionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtectionConfigValidationError) ErrorName() string { return "ProtectionConfigValidationError" }

// Error satisfies the builtin error interface
func (e ProtectionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtectionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtectionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtectionConfigValidationError{}

var _ProtectionConfig_Name_Pattern = regexp.MustCompile("^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$")

var _ProtectionConfig_Env_Pattern = regexp.MustCompile("^[a-zA-Z_]+[a-zA-Z0-9_]*$")

// Validate checks the field values on CompiledProtectionAttachmentConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CompiledProtectionAttachmentConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompiledProtectionAttachmentConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CompiledProtectionAttachmentConfigMultiError, or nil if none found.
func (m *CompiledProtectionAttachmentConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *CompiledProtectionAttachmentConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompiledProtectionAttachmentConfigValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompiledProtectionAttachmentConfigValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompiledProtectionAttachmentConfigValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAttachment()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompiledProtectionAttachmentConfigValidationError{
					field:  "Attachment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompiledProtectionAttachmentConfigValidationError{
					field:  "Attachment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAttachment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompiledProtectionAttachmentConfigValidationError{
				field:  "Attachment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRuntimeExecution()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompiledProtectionAttachmentConfigValidationError{
					field:  "RuntimeExecution",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompiledProtectionAttachmentConfigValidationError{
					field:  "RuntimeExecution",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRuntimeExecution()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompiledProtectionAttachmentConfigValidationError{
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
				err := CompiledProtectionAttachmentConfigValidationError{
					field:  fmt.Sprintf("Env[%v]", key),
					reason: "value cannot be sparse, all pairs must be non-nil",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if !_CompiledProtectionAttachmentConfig_Env_Pattern.MatchString(key) {
				err := CompiledProtectionAttachmentConfigValidationError{
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
						errors = append(errors, CompiledProtectionAttachmentConfigValidationError{
							field:  fmt.Sprintf("Env[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, CompiledProtectionAttachmentConfigValidationError{
							field:  fmt.Sprintf("Env[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return CompiledProtectionAttachmentConfigValidationError{
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
					errors = append(errors, CompiledProtectionAttachmentConfigValidationError{
						field:  fmt.Sprintf("ParameterValues[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CompiledProtectionAttachmentConfigValidationError{
						field:  fmt.Sprintf("ParameterValues[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CompiledProtectionAttachmentConfigValidationError{
					field:  fmt.Sprintf("ParameterValues[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CompiledProtectionAttachmentConfigMultiError(errors)
	}

	return nil
}

// CompiledProtectionAttachmentConfigMultiError is an error wrapping multiple
// validation errors returned by
// CompiledProtectionAttachmentConfig.ValidateAll() if the designated
// constraints aren't met.
type CompiledProtectionAttachmentConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompiledProtectionAttachmentConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompiledProtectionAttachmentConfigMultiError) AllErrors() []error { return m }

// CompiledProtectionAttachmentConfigValidationError is the validation error
// returned by CompiledProtectionAttachmentConfig.Validate if the designated
// constraints aren't met.
type CompiledProtectionAttachmentConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompiledProtectionAttachmentConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompiledProtectionAttachmentConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompiledProtectionAttachmentConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompiledProtectionAttachmentConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompiledProtectionAttachmentConfigValidationError) ErrorName() string {
	return "CompiledProtectionAttachmentConfigValidationError"
}

// Error satisfies the builtin error interface
func (e CompiledProtectionAttachmentConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompiledProtectionAttachmentConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompiledProtectionAttachmentConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompiledProtectionAttachmentConfigValidationError{}

var _CompiledProtectionAttachmentConfig_Env_Pattern = regexp.MustCompile("^[a-zA-Z_]+[a-zA-Z0-9_]*$")

// Validate checks the field values on ServiceInstanceAttachment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ServiceInstanceAttachment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceInstanceAttachment with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServiceInstanceAttachmentMultiError, or nil if none found.
func (m *ServiceInstanceAttachment) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceInstanceAttachment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Service

	// no validation rules for ReleaseChannel

	// no validation rules for Application

	if len(errors) > 0 {
		return ServiceInstanceAttachmentMultiError(errors)
	}

	return nil
}

// ServiceInstanceAttachmentMultiError is an error wrapping multiple validation
// errors returned by ServiceInstanceAttachment.ValidateAll() if the
// designated constraints aren't met.
type ServiceInstanceAttachmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceInstanceAttachmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceInstanceAttachmentMultiError) AllErrors() []error { return m }

// ServiceInstanceAttachmentValidationError is the validation error returned by
// ServiceInstanceAttachment.Validate if the designated constraints aren't met.
type ServiceInstanceAttachmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceInstanceAttachmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceInstanceAttachmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceInstanceAttachmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceInstanceAttachmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceInstanceAttachmentValidationError) ErrorName() string {
	return "ServiceInstanceAttachmentValidationError"
}

// Error satisfies the builtin error interface
func (e ServiceInstanceAttachmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceInstanceAttachment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceInstanceAttachmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceInstanceAttachmentValidationError{}

// Validate checks the field values on ReleaseChannelAttachment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReleaseChannelAttachment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReleaseChannelAttachment with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReleaseChannelAttachmentMultiError, or nil if none found.
func (m *ReleaseChannelAttachment) ValidateAll() error {
	return m.validate(true)
}

func (m *ReleaseChannelAttachment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Application

	// no validation rules for ReleaseChannel

	if len(errors) > 0 {
		return ReleaseChannelAttachmentMultiError(errors)
	}

	return nil
}

// ReleaseChannelAttachmentMultiError is an error wrapping multiple validation
// errors returned by ReleaseChannelAttachment.ValidateAll() if the designated
// constraints aren't met.
type ReleaseChannelAttachmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReleaseChannelAttachmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReleaseChannelAttachmentMultiError) AllErrors() []error { return m }

// ReleaseChannelAttachmentValidationError is the validation error returned by
// ReleaseChannelAttachment.Validate if the designated constraints aren't met.
type ReleaseChannelAttachmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReleaseChannelAttachmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReleaseChannelAttachmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReleaseChannelAttachmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReleaseChannelAttachmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReleaseChannelAttachmentValidationError) ErrorName() string {
	return "ReleaseChannelAttachmentValidationError"
}

// Error satisfies the builtin error interface
func (e ReleaseChannelAttachmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReleaseChannelAttachment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReleaseChannelAttachmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReleaseChannelAttachmentValidationError{}

// Validate checks the field values on ConvergenceAttachment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConvergenceAttachment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConvergenceAttachment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConvergenceAttachmentMultiError, or nil if none found.
func (m *ConvergenceAttachment) ValidateAll() error {
	return m.validate(true)
}

func (m *ConvergenceAttachment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DesiredStateId

	if len(errors) > 0 {
		return ConvergenceAttachmentMultiError(errors)
	}

	return nil
}

// ConvergenceAttachmentMultiError is an error wrapping multiple validation
// errors returned by ConvergenceAttachment.ValidateAll() if the designated
// constraints aren't met.
type ConvergenceAttachmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConvergenceAttachmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConvergenceAttachmentMultiError) AllErrors() []error { return m }

// ConvergenceAttachmentValidationError is the validation error returned by
// ConvergenceAttachment.Validate if the designated constraints aren't met.
type ConvergenceAttachmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConvergenceAttachmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConvergenceAttachmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConvergenceAttachmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConvergenceAttachmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConvergenceAttachmentValidationError) ErrorName() string {
	return "ConvergenceAttachmentValidationError"
}

// Error satisfies the builtin error interface
func (e ConvergenceAttachmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConvergenceAttachment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConvergenceAttachmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConvergenceAttachmentValidationError{}

// Validate checks the field values on ProtectionAttachment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProtectionAttachment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtectionAttachment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtectionAttachmentMultiError, or nil if none found.
func (m *ProtectionAttachment) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtectionAttachment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofAttachmentPresent := false
	switch v := m.Attachment.(type) {
	case *ProtectionAttachment_ServiceInstance:
		if v == nil {
			err := ProtectionAttachmentValidationError{
				field:  "Attachment",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofAttachmentPresent = true

		if all {
			switch v := interface{}(m.GetServiceInstance()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProtectionAttachmentValidationError{
						field:  "ServiceInstance",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProtectionAttachmentValidationError{
						field:  "ServiceInstance",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetServiceInstance()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtectionAttachmentValidationError{
					field:  "ServiceInstance",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProtectionAttachment_ReleaseChannel:
		if v == nil {
			err := ProtectionAttachmentValidationError{
				field:  "Attachment",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofAttachmentPresent = true

		if all {
			switch v := interface{}(m.GetReleaseChannel()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProtectionAttachmentValidationError{
						field:  "ReleaseChannel",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProtectionAttachmentValidationError{
						field:  "ReleaseChannel",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetReleaseChannel()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtectionAttachmentValidationError{
					field:  "ReleaseChannel",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProtectionAttachment_Convergence:
		if v == nil {
			err := ProtectionAttachmentValidationError{
				field:  "Attachment",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofAttachmentPresent = true

		if all {
			switch v := interface{}(m.GetConvergence()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProtectionAttachmentValidationError{
						field:  "Convergence",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProtectionAttachmentValidationError{
						field:  "Convergence",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetConvergence()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtectionAttachmentValidationError{
					field:  "Convergence",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofAttachmentPresent {
		err := ProtectionAttachmentValidationError{
			field:  "Attachment",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ProtectionAttachmentMultiError(errors)
	}

	return nil
}

// ProtectionAttachmentMultiError is an error wrapping multiple validation
// errors returned by ProtectionAttachment.ValidateAll() if the designated
// constraints aren't met.
type ProtectionAttachmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtectionAttachmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtectionAttachmentMultiError) AllErrors() []error { return m }

// ProtectionAttachmentValidationError is the validation error returned by
// ProtectionAttachment.Validate if the designated constraints aren't met.
type ProtectionAttachmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtectionAttachmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtectionAttachmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtectionAttachmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtectionAttachmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtectionAttachmentValidationError) ErrorName() string {
	return "ProtectionAttachmentValidationError"
}

// Error satisfies the builtin error interface
func (e ProtectionAttachmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtectionAttachment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtectionAttachmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtectionAttachmentValidationError{}
