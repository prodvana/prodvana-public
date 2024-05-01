// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/common_config/env.proto

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

// Validate checks the field values on EnvValue with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EnvValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EnvValue with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EnvValueMultiError, or nil
// if none found.
func (m *EnvValue) ValidateAll() error {
	return m.validate(true)
}

func (m *EnvValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofValueOneofPresent := false
	switch v := m.ValueOneof.(type) {
	case *EnvValue_Value:
		if v == nil {
			err := EnvValueValidationError{
				field:  "ValueOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValueOneofPresent = true
		// no validation rules for Value
	case *EnvValue_RawSecret:
		if v == nil {
			err := EnvValueValidationError{
				field:  "ValueOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValueOneofPresent = true
		// no validation rules for RawSecret
	case *EnvValue_Secret:
		if v == nil {
			err := EnvValueValidationError{
				field:  "ValueOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValueOneofPresent = true

		if all {
			switch v := interface{}(m.GetSecret()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnvValueValidationError{
						field:  "Secret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnvValueValidationError{
						field:  "Secret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSecret()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnvValueValidationError{
					field:  "Secret",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *EnvValue_KubernetesSecret:
		if v == nil {
			err := EnvValueValidationError{
				field:  "ValueOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValueOneofPresent = true

		if all {
			switch v := interface{}(m.GetKubernetesSecret()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnvValueValidationError{
						field:  "KubernetesSecret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnvValueValidationError{
						field:  "KubernetesSecret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetKubernetesSecret()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnvValueValidationError{
					field:  "KubernetesSecret",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofValueOneofPresent {
		err := EnvValueValidationError{
			field:  "ValueOneof",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return EnvValueMultiError(errors)
	}

	return nil
}

// EnvValueMultiError is an error wrapping multiple validation errors returned
// by EnvValue.ValidateAll() if the designated constraints aren't met.
type EnvValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnvValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnvValueMultiError) AllErrors() []error { return m }

// EnvValueValidationError is the validation error returned by
// EnvValue.Validate if the designated constraints aren't met.
type EnvValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnvValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnvValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnvValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnvValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnvValueValidationError) ErrorName() string { return "EnvValueValidationError" }

// Error satisfies the builtin error interface
func (e EnvValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnvValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnvValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnvValueValidationError{}

// Validate checks the field values on SecretReferenceValue with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SecretReferenceValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SecretReferenceValue with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SecretReferenceValueMultiError, or nil if none found.
func (m *SecretReferenceValue) ValidateAll() error {
	return m.validate(true)
}

func (m *SecretReferenceValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofValueOneofPresent := false
	switch v := m.ValueOneof.(type) {
	case *SecretReferenceValue_Secret:
		if v == nil {
			err := SecretReferenceValueValidationError{
				field:  "ValueOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValueOneofPresent = true

		if m.GetSecret() == nil {
			err := SecretReferenceValueValidationError{
				field:  "Secret",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetSecret()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SecretReferenceValueValidationError{
						field:  "Secret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SecretReferenceValueValidationError{
						field:  "Secret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSecret()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecretReferenceValueValidationError{
					field:  "Secret",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *SecretReferenceValue_KubernetesSecret:
		if v == nil {
			err := SecretReferenceValueValidationError{
				field:  "ValueOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValueOneofPresent = true

		if m.GetKubernetesSecret() == nil {
			err := SecretReferenceValueValidationError{
				field:  "KubernetesSecret",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetKubernetesSecret()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SecretReferenceValueValidationError{
						field:  "KubernetesSecret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SecretReferenceValueValidationError{
						field:  "KubernetesSecret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetKubernetesSecret()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SecretReferenceValueValidationError{
					field:  "KubernetesSecret",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofValueOneofPresent {
		err := SecretReferenceValueValidationError{
			field:  "ValueOneof",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SecretReferenceValueMultiError(errors)
	}

	return nil
}

// SecretReferenceValueMultiError is an error wrapping multiple validation
// errors returned by SecretReferenceValue.ValidateAll() if the designated
// constraints aren't met.
type SecretReferenceValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SecretReferenceValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SecretReferenceValueMultiError) AllErrors() []error { return m }

// SecretReferenceValueValidationError is the validation error returned by
// SecretReferenceValue.Validate if the designated constraints aren't met.
type SecretReferenceValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecretReferenceValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecretReferenceValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecretReferenceValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecretReferenceValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecretReferenceValueValidationError) ErrorName() string {
	return "SecretReferenceValueValidationError"
}

// Error satisfies the builtin error interface
func (e SecretReferenceValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecretReferenceValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecretReferenceValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecretReferenceValueValidationError{}

// Validate checks the field values on NativeSecretReferenceValue with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NativeSecretReferenceValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NativeSecretReferenceValue with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NativeSecretReferenceValueMultiError, or nil if none found.
func (m *NativeSecretReferenceValue) ValidateAll() error {
	return m.validate(true)
}

func (m *NativeSecretReferenceValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofValueOneofPresent := false
	switch v := m.ValueOneof.(type) {
	case *NativeSecretReferenceValue_Secret:
		if v == nil {
			err := NativeSecretReferenceValueValidationError{
				field:  "ValueOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValueOneofPresent = true

		if m.GetSecret() == nil {
			err := NativeSecretReferenceValueValidationError{
				field:  "Secret",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetSecret()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NativeSecretReferenceValueValidationError{
						field:  "Secret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NativeSecretReferenceValueValidationError{
						field:  "Secret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSecret()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NativeSecretReferenceValueValidationError{
					field:  "Secret",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofValueOneofPresent {
		err := NativeSecretReferenceValueValidationError{
			field:  "ValueOneof",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return NativeSecretReferenceValueMultiError(errors)
	}

	return nil
}

// NativeSecretReferenceValueMultiError is an error wrapping multiple
// validation errors returned by NativeSecretReferenceValue.ValidateAll() if
// the designated constraints aren't met.
type NativeSecretReferenceValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NativeSecretReferenceValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NativeSecretReferenceValueMultiError) AllErrors() []error { return m }

// NativeSecretReferenceValueValidationError is the validation error returned
// by NativeSecretReferenceValue.Validate if the designated constraints aren't met.
type NativeSecretReferenceValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NativeSecretReferenceValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NativeSecretReferenceValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NativeSecretReferenceValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NativeSecretReferenceValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NativeSecretReferenceValueValidationError) ErrorName() string {
	return "NativeSecretReferenceValueValidationError"
}

// Error satisfies the builtin error interface
func (e NativeSecretReferenceValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNativeSecretReferenceValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NativeSecretReferenceValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NativeSecretReferenceValueValidationError{}

// Validate checks the field values on Secret with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Secret) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Secret with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SecretMultiError, or nil if none found.
func (m *Secret) ValidateAll() error {
	return m.validate(true)
}

func (m *Secret) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := SecretValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetVersion()) < 1 {
		err := SecretValidationError{
			field:  "Version",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SecretMultiError(errors)
	}

	return nil
}

// SecretMultiError is an error wrapping multiple validation errors returned by
// Secret.ValidateAll() if the designated constraints aren't met.
type SecretMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SecretMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SecretMultiError) AllErrors() []error { return m }

// SecretValidationError is the validation error returned by Secret.Validate if
// the designated constraints aren't met.
type SecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecretValidationError) ErrorName() string { return "SecretValidationError" }

// Error satisfies the builtin error interface
func (e SecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecretValidationError{}

// Validate checks the field values on KubernetesSecret with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *KubernetesSecret) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KubernetesSecret with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// KubernetesSecretMultiError, or nil if none found.
func (m *KubernetesSecret) ValidateAll() error {
	return m.validate(true)
}

func (m *KubernetesSecret) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetSecretName()) < 1 {
		err := KubernetesSecretValidationError{
			field:  "SecretName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := KubernetesSecretValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return KubernetesSecretMultiError(errors)
	}

	return nil
}

// KubernetesSecretMultiError is an error wrapping multiple validation errors
// returned by KubernetesSecret.ValidateAll() if the designated constraints
// aren't met.
type KubernetesSecretMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KubernetesSecretMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KubernetesSecretMultiError) AllErrors() []error { return m }

// KubernetesSecretValidationError is the validation error returned by
// KubernetesSecret.Validate if the designated constraints aren't met.
type KubernetesSecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KubernetesSecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KubernetesSecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KubernetesSecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KubernetesSecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KubernetesSecretValidationError) ErrorName() string { return "KubernetesSecretValidationError" }

// Error satisfies the builtin error interface
func (e KubernetesSecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKubernetesSecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KubernetesSecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KubernetesSecretValidationError{}
