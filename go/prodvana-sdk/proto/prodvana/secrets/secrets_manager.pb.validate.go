// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/secrets/secrets_manager.proto

package secrets

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

// Validate checks the field values on ListSecretsReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListSecretsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSecretsReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListSecretsReqMultiError,
// or nil if none found.
func (m *ListSecretsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSecretsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListSecretsReqMultiError(errors)
	}

	return nil
}

// ListSecretsReqMultiError is an error wrapping multiple validation errors
// returned by ListSecretsReq.ValidateAll() if the designated constraints
// aren't met.
type ListSecretsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSecretsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSecretsReqMultiError) AllErrors() []error { return m }

// ListSecretsReqValidationError is the validation error returned by
// ListSecretsReq.Validate if the designated constraints aren't met.
type ListSecretsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSecretsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSecretsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSecretsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSecretsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSecretsReqValidationError) ErrorName() string { return "ListSecretsReqValidationError" }

// Error satisfies the builtin error interface
func (e ListSecretsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSecretsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSecretsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSecretsReqValidationError{}

// Validate checks the field values on ListSecretsResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListSecretsResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSecretsResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSecretsRespMultiError, or nil if none found.
func (m *ListSecretsResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSecretsResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSecrets() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListSecretsRespValidationError{
						field:  fmt.Sprintf("Secrets[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListSecretsRespValidationError{
						field:  fmt.Sprintf("Secrets[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSecretsRespValidationError{
					field:  fmt.Sprintf("Secrets[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListSecretsRespMultiError(errors)
	}

	return nil
}

// ListSecretsRespMultiError is an error wrapping multiple validation errors
// returned by ListSecretsResp.ValidateAll() if the designated constraints
// aren't met.
type ListSecretsRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSecretsRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSecretsRespMultiError) AllErrors() []error { return m }

// ListSecretsRespValidationError is the validation error returned by
// ListSecretsResp.Validate if the designated constraints aren't met.
type ListSecretsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSecretsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSecretsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSecretsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSecretsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSecretsRespValidationError) ErrorName() string { return "ListSecretsRespValidationError" }

// Error satisfies the builtin error interface
func (e ListSecretsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSecretsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSecretsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSecretsRespValidationError{}

// Validate checks the field values on ListSecretVersionsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListSecretVersionsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSecretVersionsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSecretVersionsReqMultiError, or nil if none found.
func (m *ListSecretVersionsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSecretVersionsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := ListSecretVersionsReqValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListSecretVersionsReqMultiError(errors)
	}

	return nil
}

// ListSecretVersionsReqMultiError is an error wrapping multiple validation
// errors returned by ListSecretVersionsReq.ValidateAll() if the designated
// constraints aren't met.
type ListSecretVersionsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSecretVersionsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSecretVersionsReqMultiError) AllErrors() []error { return m }

// ListSecretVersionsReqValidationError is the validation error returned by
// ListSecretVersionsReq.Validate if the designated constraints aren't met.
type ListSecretVersionsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSecretVersionsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSecretVersionsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSecretVersionsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSecretVersionsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSecretVersionsReqValidationError) ErrorName() string {
	return "ListSecretVersionsReqValidationError"
}

// Error satisfies the builtin error interface
func (e ListSecretVersionsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSecretVersionsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSecretVersionsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSecretVersionsReqValidationError{}

// Validate checks the field values on ListSecretVersionsResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListSecretVersionsResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSecretVersionsResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSecretVersionsRespMultiError, or nil if none found.
func (m *ListSecretVersionsResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSecretVersionsResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListSecretVersionsRespMultiError(errors)
	}

	return nil
}

// ListSecretVersionsRespMultiError is an error wrapping multiple validation
// errors returned by ListSecretVersionsResp.ValidateAll() if the designated
// constraints aren't met.
type ListSecretVersionsRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSecretVersionsRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSecretVersionsRespMultiError) AllErrors() []error { return m }

// ListSecretVersionsRespValidationError is the validation error returned by
// ListSecretVersionsResp.Validate if the designated constraints aren't met.
type ListSecretVersionsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSecretVersionsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSecretVersionsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSecretVersionsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSecretVersionsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSecretVersionsRespValidationError) ErrorName() string {
	return "ListSecretVersionsRespValidationError"
}

// Error satisfies the builtin error interface
func (e ListSecretVersionsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSecretVersionsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSecretVersionsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSecretVersionsRespValidationError{}

// Validate checks the field values on SetSecretReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SetSecretReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetSecretReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SetSecretReqMultiError, or
// nil if none found.
func (m *SetSecretReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SetSecretReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := SetSecretReqValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetValue()) < 1 {
		err := SetSecretReqValidationError{
			field:  "Value",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SetSecretReqMultiError(errors)
	}

	return nil
}

// SetSecretReqMultiError is an error wrapping multiple validation errors
// returned by SetSecretReq.ValidateAll() if the designated constraints aren't met.
type SetSecretReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetSecretReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetSecretReqMultiError) AllErrors() []error { return m }

// SetSecretReqValidationError is the validation error returned by
// SetSecretReq.Validate if the designated constraints aren't met.
type SetSecretReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetSecretReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetSecretReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetSecretReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetSecretReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetSecretReqValidationError) ErrorName() string { return "SetSecretReqValidationError" }

// Error satisfies the builtin error interface
func (e SetSecretReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetSecretReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetSecretReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetSecretReqValidationError{}

// Validate checks the field values on SetSecretResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SetSecretResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetSecretResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SetSecretRespMultiError, or
// nil if none found.
func (m *SetSecretResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SetSecretResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Version

	if len(errors) > 0 {
		return SetSecretRespMultiError(errors)
	}

	return nil
}

// SetSecretRespMultiError is an error wrapping multiple validation errors
// returned by SetSecretResp.ValidateAll() if the designated constraints
// aren't met.
type SetSecretRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetSecretRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetSecretRespMultiError) AllErrors() []error { return m }

// SetSecretRespValidationError is the validation error returned by
// SetSecretResp.Validate if the designated constraints aren't met.
type SetSecretRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetSecretRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetSecretRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetSecretRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetSecretRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetSecretRespValidationError) ErrorName() string { return "SetSecretRespValidationError" }

// Error satisfies the builtin error interface
func (e SetSecretRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetSecretResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetSecretRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetSecretRespValidationError{}

// Validate checks the field values on DeleteSecretReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteSecretReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSecretReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSecretReqMultiError, or nil if none found.
func (m *DeleteSecretReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSecretReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := DeleteSecretReqValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteSecretReqMultiError(errors)
	}

	return nil
}

// DeleteSecretReqMultiError is an error wrapping multiple validation errors
// returned by DeleteSecretReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteSecretReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSecretReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSecretReqMultiError) AllErrors() []error { return m }

// DeleteSecretReqValidationError is the validation error returned by
// DeleteSecretReq.Validate if the designated constraints aren't met.
type DeleteSecretReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSecretReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSecretReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSecretReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSecretReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSecretReqValidationError) ErrorName() string { return "DeleteSecretReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteSecretReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSecretReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSecretReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSecretReqValidationError{}

// Validate checks the field values on DeleteSecretResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteSecretResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSecretResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSecretRespMultiError, or nil if none found.
func (m *DeleteSecretResp) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSecretResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteSecretRespMultiError(errors)
	}

	return nil
}

// DeleteSecretRespMultiError is an error wrapping multiple validation errors
// returned by DeleteSecretResp.ValidateAll() if the designated constraints
// aren't met.
type DeleteSecretRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSecretRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSecretRespMultiError) AllErrors() []error { return m }

// DeleteSecretRespValidationError is the validation error returned by
// DeleteSecretResp.Validate if the designated constraints aren't met.
type DeleteSecretRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSecretRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSecretRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSecretRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSecretRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSecretRespValidationError) ErrorName() string { return "DeleteSecretRespValidationError" }

// Error satisfies the builtin error interface
func (e DeleteSecretRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSecretResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSecretRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSecretRespValidationError{}

// Validate checks the field values on DeleteSecretVersionReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteSecretVersionReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSecretVersionReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSecretVersionReqMultiError, or nil if none found.
func (m *DeleteSecretVersionReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSecretVersionReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := DeleteSecretVersionReqValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetVersion()) < 1 {
		err := DeleteSecretVersionReqValidationError{
			field:  "Version",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteSecretVersionReqMultiError(errors)
	}

	return nil
}

// DeleteSecretVersionReqMultiError is an error wrapping multiple validation
// errors returned by DeleteSecretVersionReq.ValidateAll() if the designated
// constraints aren't met.
type DeleteSecretVersionReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSecretVersionReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSecretVersionReqMultiError) AllErrors() []error { return m }

// DeleteSecretVersionReqValidationError is the validation error returned by
// DeleteSecretVersionReq.Validate if the designated constraints aren't met.
type DeleteSecretVersionReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSecretVersionReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSecretVersionReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSecretVersionReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSecretVersionReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSecretVersionReqValidationError) ErrorName() string {
	return "DeleteSecretVersionReqValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSecretVersionReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSecretVersionReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSecretVersionReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSecretVersionReqValidationError{}

// Validate checks the field values on DeleteSecretVersionResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteSecretVersionResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSecretVersionResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSecretVersionRespMultiError, or nil if none found.
func (m *DeleteSecretVersionResp) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSecretVersionResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteSecretVersionRespMultiError(errors)
	}

	return nil
}

// DeleteSecretVersionRespMultiError is an error wrapping multiple validation
// errors returned by DeleteSecretVersionResp.ValidateAll() if the designated
// constraints aren't met.
type DeleteSecretVersionRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSecretVersionRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSecretVersionRespMultiError) AllErrors() []error { return m }

// DeleteSecretVersionRespValidationError is the validation error returned by
// DeleteSecretVersionResp.Validate if the designated constraints aren't met.
type DeleteSecretVersionRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSecretVersionRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSecretVersionRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSecretVersionRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSecretVersionRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSecretVersionRespValidationError) ErrorName() string {
	return "DeleteSecretVersionRespValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSecretVersionRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSecretVersionResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSecretVersionRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSecretVersionRespValidationError{}