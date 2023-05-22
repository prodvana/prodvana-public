// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/delivery_extension/manager.proto

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

	version "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/version"
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

	_ = version.Source(0)
)

// Validate checks the field values on ConfigureDeliveryExtensionReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConfigureDeliveryExtensionReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfigureDeliveryExtensionReq with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ConfigureDeliveryExtensionReqMultiError, or nil if none found.
func (m *ConfigureDeliveryExtensionReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfigureDeliveryExtensionReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetDeliveryExtensionConfig() == nil {
		err := ConfigureDeliveryExtensionReqValidationError{
			field:  "DeliveryExtensionConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDeliveryExtensionConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfigureDeliveryExtensionReqValidationError{
					field:  "DeliveryExtensionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfigureDeliveryExtensionReqValidationError{
					field:  "DeliveryExtensionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeliveryExtensionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigureDeliveryExtensionReqValidationError{
				field:  "DeliveryExtensionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Source

	if all {
		switch v := interface{}(m.GetSourceMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfigureDeliveryExtensionReqValidationError{
					field:  "SourceMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfigureDeliveryExtensionReqValidationError{
					field:  "SourceMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSourceMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigureDeliveryExtensionReqValidationError{
				field:  "SourceMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ConfigureDeliveryExtensionReqMultiError(errors)
	}

	return nil
}

// ConfigureDeliveryExtensionReqMultiError is an error wrapping multiple
// validation errors returned by ConfigureDeliveryExtensionReq.ValidateAll()
// if the designated constraints aren't met.
type ConfigureDeliveryExtensionReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigureDeliveryExtensionReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigureDeliveryExtensionReqMultiError) AllErrors() []error { return m }

// ConfigureDeliveryExtensionReqValidationError is the validation error
// returned by ConfigureDeliveryExtensionReq.Validate if the designated
// constraints aren't met.
type ConfigureDeliveryExtensionReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigureDeliveryExtensionReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigureDeliveryExtensionReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigureDeliveryExtensionReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigureDeliveryExtensionReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigureDeliveryExtensionReqValidationError) ErrorName() string {
	return "ConfigureDeliveryExtensionReqValidationError"
}

// Error satisfies the builtin error interface
func (e ConfigureDeliveryExtensionReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigureDeliveryExtensionReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigureDeliveryExtensionReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigureDeliveryExtensionReqValidationError{}

// Validate checks the field values on ConfigureDeliveryExtensionResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConfigureDeliveryExtensionResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfigureDeliveryExtensionResp with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ConfigureDeliveryExtensionRespMultiError, or nil if none found.
func (m *ConfigureDeliveryExtensionResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfigureDeliveryExtensionResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DeliveryExtensionId

	// no validation rules for Version

	if len(errors) > 0 {
		return ConfigureDeliveryExtensionRespMultiError(errors)
	}

	return nil
}

// ConfigureDeliveryExtensionRespMultiError is an error wrapping multiple
// validation errors returned by ConfigureDeliveryExtensionResp.ValidateAll()
// if the designated constraints aren't met.
type ConfigureDeliveryExtensionRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigureDeliveryExtensionRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigureDeliveryExtensionRespMultiError) AllErrors() []error { return m }

// ConfigureDeliveryExtensionRespValidationError is the validation error
// returned by ConfigureDeliveryExtensionResp.Validate if the designated
// constraints aren't met.
type ConfigureDeliveryExtensionRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigureDeliveryExtensionRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigureDeliveryExtensionRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigureDeliveryExtensionRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigureDeliveryExtensionRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigureDeliveryExtensionRespValidationError) ErrorName() string {
	return "ConfigureDeliveryExtensionRespValidationError"
}

// Error satisfies the builtin error interface
func (e ConfigureDeliveryExtensionRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigureDeliveryExtensionResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigureDeliveryExtensionRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigureDeliveryExtensionRespValidationError{}

// Validate checks the field values on ValidateConfigureDeliveryExtensionResp
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *ValidateConfigureDeliveryExtensionResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// ValidateConfigureDeliveryExtensionResp with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// ValidateConfigureDeliveryExtensionRespMultiError, or nil if none found.
func (m *ValidateConfigureDeliveryExtensionResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ValidateConfigureDeliveryExtensionResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ValidateConfigureDeliveryExtensionRespMultiError(errors)
	}

	return nil
}

// ValidateConfigureDeliveryExtensionRespMultiError is an error wrapping
// multiple validation errors returned by
// ValidateConfigureDeliveryExtensionResp.ValidateAll() if the designated
// constraints aren't met.
type ValidateConfigureDeliveryExtensionRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValidateConfigureDeliveryExtensionRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValidateConfigureDeliveryExtensionRespMultiError) AllErrors() []error { return m }

// ValidateConfigureDeliveryExtensionRespValidationError is the validation
// error returned by ValidateConfigureDeliveryExtensionResp.Validate if the
// designated constraints aren't met.
type ValidateConfigureDeliveryExtensionRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateConfigureDeliveryExtensionRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateConfigureDeliveryExtensionRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateConfigureDeliveryExtensionRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateConfigureDeliveryExtensionRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateConfigureDeliveryExtensionRespValidationError) ErrorName() string {
	return "ValidateConfigureDeliveryExtensionRespValidationError"
}

// Error satisfies the builtin error interface
func (e ValidateConfigureDeliveryExtensionRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateConfigureDeliveryExtensionResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateConfigureDeliveryExtensionRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateConfigureDeliveryExtensionRespValidationError{}

// Validate checks the field values on ListDeliveryExtensionsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListDeliveryExtensionsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListDeliveryExtensionsReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListDeliveryExtensionsReqMultiError, or nil if none found.
func (m *ListDeliveryExtensionsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListDeliveryExtensionsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageToken

	// no validation rules for PageSize

	if len(errors) > 0 {
		return ListDeliveryExtensionsReqMultiError(errors)
	}

	return nil
}

// ListDeliveryExtensionsReqMultiError is an error wrapping multiple validation
// errors returned by ListDeliveryExtensionsReq.ValidateAll() if the
// designated constraints aren't met.
type ListDeliveryExtensionsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListDeliveryExtensionsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListDeliveryExtensionsReqMultiError) AllErrors() []error { return m }

// ListDeliveryExtensionsReqValidationError is the validation error returned by
// ListDeliveryExtensionsReq.Validate if the designated constraints aren't met.
type ListDeliveryExtensionsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDeliveryExtensionsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDeliveryExtensionsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDeliveryExtensionsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDeliveryExtensionsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDeliveryExtensionsReqValidationError) ErrorName() string {
	return "ListDeliveryExtensionsReqValidationError"
}

// Error satisfies the builtin error interface
func (e ListDeliveryExtensionsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDeliveryExtensionsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDeliveryExtensionsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDeliveryExtensionsReqValidationError{}

// Validate checks the field values on ListDeliveryExtensionsResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListDeliveryExtensionsResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListDeliveryExtensionsResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListDeliveryExtensionsRespMultiError, or nil if none found.
func (m *ListDeliveryExtensionsResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ListDeliveryExtensionsResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDeliveryExtensions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListDeliveryExtensionsRespValidationError{
						field:  fmt.Sprintf("DeliveryExtensions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListDeliveryExtensionsRespValidationError{
						field:  fmt.Sprintf("DeliveryExtensions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListDeliveryExtensionsRespValidationError{
					field:  fmt.Sprintf("DeliveryExtensions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListDeliveryExtensionsRespMultiError(errors)
	}

	return nil
}

// ListDeliveryExtensionsRespMultiError is an error wrapping multiple
// validation errors returned by ListDeliveryExtensionsResp.ValidateAll() if
// the designated constraints aren't met.
type ListDeliveryExtensionsRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListDeliveryExtensionsRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListDeliveryExtensionsRespMultiError) AllErrors() []error { return m }

// ListDeliveryExtensionsRespValidationError is the validation error returned
// by ListDeliveryExtensionsResp.Validate if the designated constraints aren't met.
type ListDeliveryExtensionsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDeliveryExtensionsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDeliveryExtensionsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDeliveryExtensionsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDeliveryExtensionsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDeliveryExtensionsRespValidationError) ErrorName() string {
	return "ListDeliveryExtensionsRespValidationError"
}

// Error satisfies the builtin error interface
func (e ListDeliveryExtensionsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDeliveryExtensionsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDeliveryExtensionsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDeliveryExtensionsRespValidationError{}

// Validate checks the field values on GetDeliveryExtensionReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeliveryExtensionReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeliveryExtensionReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeliveryExtensionReqMultiError, or nil if none found.
func (m *GetDeliveryExtensionReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeliveryExtensionReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDeliveryExtension()) < 1 {
		err := GetDeliveryExtensionReqValidationError{
			field:  "DeliveryExtension",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetDeliveryExtensionReqMultiError(errors)
	}

	return nil
}

// GetDeliveryExtensionReqMultiError is an error wrapping multiple validation
// errors returned by GetDeliveryExtensionReq.ValidateAll() if the designated
// constraints aren't met.
type GetDeliveryExtensionReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeliveryExtensionReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeliveryExtensionReqMultiError) AllErrors() []error { return m }

// GetDeliveryExtensionReqValidationError is the validation error returned by
// GetDeliveryExtensionReq.Validate if the designated constraints aren't met.
type GetDeliveryExtensionReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeliveryExtensionReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeliveryExtensionReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeliveryExtensionReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeliveryExtensionReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeliveryExtensionReqValidationError) ErrorName() string {
	return "GetDeliveryExtensionReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeliveryExtensionReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeliveryExtensionReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeliveryExtensionReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeliveryExtensionReqValidationError{}

// Validate checks the field values on GetDeliveryExtensionResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeliveryExtensionResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeliveryExtensionResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeliveryExtensionRespMultiError, or nil if none found.
func (m *GetDeliveryExtensionResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeliveryExtensionResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDeliveryExtension()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetDeliveryExtensionRespValidationError{
					field:  "DeliveryExtension",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetDeliveryExtensionRespValidationError{
					field:  "DeliveryExtension",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeliveryExtension()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDeliveryExtensionRespValidationError{
				field:  "DeliveryExtension",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetDeliveryExtensionRespMultiError(errors)
	}

	return nil
}

// GetDeliveryExtensionRespMultiError is an error wrapping multiple validation
// errors returned by GetDeliveryExtensionResp.ValidateAll() if the designated
// constraints aren't met.
type GetDeliveryExtensionRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeliveryExtensionRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeliveryExtensionRespMultiError) AllErrors() []error { return m }

// GetDeliveryExtensionRespValidationError is the validation error returned by
// GetDeliveryExtensionResp.Validate if the designated constraints aren't met.
type GetDeliveryExtensionRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeliveryExtensionRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeliveryExtensionRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeliveryExtensionRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeliveryExtensionRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeliveryExtensionRespValidationError) ErrorName() string {
	return "GetDeliveryExtensionRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeliveryExtensionRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeliveryExtensionResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeliveryExtensionRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeliveryExtensionRespValidationError{}

// Validate checks the field values on GetDeliveryExtensionConfigReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeliveryExtensionConfigReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeliveryExtensionConfigReq with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetDeliveryExtensionConfigReqMultiError, or nil if none found.
func (m *GetDeliveryExtensionConfigReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeliveryExtensionConfigReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDeliveryExtension()) < 1 {
		err := GetDeliveryExtensionConfigReqValidationError{
			field:  "DeliveryExtension",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Version

	if len(errors) > 0 {
		return GetDeliveryExtensionConfigReqMultiError(errors)
	}

	return nil
}

// GetDeliveryExtensionConfigReqMultiError is an error wrapping multiple
// validation errors returned by GetDeliveryExtensionConfigReq.ValidateAll()
// if the designated constraints aren't met.
type GetDeliveryExtensionConfigReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeliveryExtensionConfigReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeliveryExtensionConfigReqMultiError) AllErrors() []error { return m }

// GetDeliveryExtensionConfigReqValidationError is the validation error
// returned by GetDeliveryExtensionConfigReq.Validate if the designated
// constraints aren't met.
type GetDeliveryExtensionConfigReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeliveryExtensionConfigReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeliveryExtensionConfigReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeliveryExtensionConfigReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeliveryExtensionConfigReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeliveryExtensionConfigReqValidationError) ErrorName() string {
	return "GetDeliveryExtensionConfigReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeliveryExtensionConfigReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeliveryExtensionConfigReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeliveryExtensionConfigReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeliveryExtensionConfigReqValidationError{}

// Validate checks the field values on GetDeliveryExtensionConfigResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeliveryExtensionConfigResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeliveryExtensionConfigResp with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetDeliveryExtensionConfigRespMultiError, or nil if none found.
func (m *GetDeliveryExtensionConfigResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeliveryExtensionConfigResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetDeliveryExtensionConfigRespValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetDeliveryExtensionConfigRespValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDeliveryExtensionConfigRespValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Version

	if len(errors) > 0 {
		return GetDeliveryExtensionConfigRespMultiError(errors)
	}

	return nil
}

// GetDeliveryExtensionConfigRespMultiError is an error wrapping multiple
// validation errors returned by GetDeliveryExtensionConfigResp.ValidateAll()
// if the designated constraints aren't met.
type GetDeliveryExtensionConfigRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeliveryExtensionConfigRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeliveryExtensionConfigRespMultiError) AllErrors() []error { return m }

// GetDeliveryExtensionConfigRespValidationError is the validation error
// returned by GetDeliveryExtensionConfigResp.Validate if the designated
// constraints aren't met.
type GetDeliveryExtensionConfigRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeliveryExtensionConfigRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeliveryExtensionConfigRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeliveryExtensionConfigRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeliveryExtensionConfigRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeliveryExtensionConfigRespValidationError) ErrorName() string {
	return "GetDeliveryExtensionConfigRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeliveryExtensionConfigRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeliveryExtensionConfigResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeliveryExtensionConfigRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeliveryExtensionConfigRespValidationError{}