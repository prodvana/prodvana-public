// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/release_channel/release_channel_manager.proto

package release_channel

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

// Validate checks the field values on ConfigureReleaseChannelReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConfigureReleaseChannelReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfigureReleaseChannelReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConfigureReleaseChannelReqMultiError, or nil if none found.
func (m *ConfigureReleaseChannelReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfigureReleaseChannelReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetReleaseChannel()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfigureReleaseChannelReqValidationError{
					field:  "ReleaseChannel",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfigureReleaseChannelReqValidationError{
					field:  "ReleaseChannel",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReleaseChannel()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigureReleaseChannelReqValidationError{
				field:  "ReleaseChannel",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ClusterId

	// no validation rules for ClusterName

	if len(errors) > 0 {
		return ConfigureReleaseChannelReqMultiError(errors)
	}

	return nil
}

// ConfigureReleaseChannelReqMultiError is an error wrapping multiple
// validation errors returned by ConfigureReleaseChannelReq.ValidateAll() if
// the designated constraints aren't met.
type ConfigureReleaseChannelReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigureReleaseChannelReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigureReleaseChannelReqMultiError) AllErrors() []error { return m }

// ConfigureReleaseChannelReqValidationError is the validation error returned
// by ConfigureReleaseChannelReq.Validate if the designated constraints aren't met.
type ConfigureReleaseChannelReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigureReleaseChannelReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigureReleaseChannelReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigureReleaseChannelReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigureReleaseChannelReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigureReleaseChannelReqValidationError) ErrorName() string {
	return "ConfigureReleaseChannelReqValidationError"
}

// Error satisfies the builtin error interface
func (e ConfigureReleaseChannelReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigureReleaseChannelReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigureReleaseChannelReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigureReleaseChannelReqValidationError{}

// Validate checks the field values on ConfigureReleaseChannelResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConfigureReleaseChannelResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfigureReleaseChannelResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConfigureReleaseChannelRespMultiError, or nil if none found.
func (m *ConfigureReleaseChannelResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfigureReleaseChannelResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Version

	if len(errors) > 0 {
		return ConfigureReleaseChannelRespMultiError(errors)
	}

	return nil
}

// ConfigureReleaseChannelRespMultiError is an error wrapping multiple
// validation errors returned by ConfigureReleaseChannelResp.ValidateAll() if
// the designated constraints aren't met.
type ConfigureReleaseChannelRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigureReleaseChannelRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigureReleaseChannelRespMultiError) AllErrors() []error { return m }

// ConfigureReleaseChannelRespValidationError is the validation error returned
// by ConfigureReleaseChannelResp.Validate if the designated constraints
// aren't met.
type ConfigureReleaseChannelRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigureReleaseChannelRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigureReleaseChannelRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigureReleaseChannelRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigureReleaseChannelRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigureReleaseChannelRespValidationError) ErrorName() string {
	return "ConfigureReleaseChannelRespValidationError"
}

// Error satisfies the builtin error interface
func (e ConfigureReleaseChannelRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigureReleaseChannelResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigureReleaseChannelRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigureReleaseChannelRespValidationError{}

// Validate checks the field values on DeleteReleaseChannelReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteReleaseChannelReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteReleaseChannelReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteReleaseChannelReqMultiError, or nil if none found.
func (m *DeleteReleaseChannelReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteReleaseChannelReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ReleaseChannel

	if len(errors) > 0 {
		return DeleteReleaseChannelReqMultiError(errors)
	}

	return nil
}

// DeleteReleaseChannelReqMultiError is an error wrapping multiple validation
// errors returned by DeleteReleaseChannelReq.ValidateAll() if the designated
// constraints aren't met.
type DeleteReleaseChannelReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteReleaseChannelReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteReleaseChannelReqMultiError) AllErrors() []error { return m }

// DeleteReleaseChannelReqValidationError is the validation error returned by
// DeleteReleaseChannelReq.Validate if the designated constraints aren't met.
type DeleteReleaseChannelReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteReleaseChannelReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteReleaseChannelReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteReleaseChannelReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteReleaseChannelReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteReleaseChannelReqValidationError) ErrorName() string {
	return "DeleteReleaseChannelReqValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteReleaseChannelReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteReleaseChannelReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteReleaseChannelReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteReleaseChannelReqValidationError{}

// Validate checks the field values on DeleteReleaseChannelResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteReleaseChannelResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteReleaseChannelResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteReleaseChannelRespMultiError, or nil if none found.
func (m *DeleteReleaseChannelResp) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteReleaseChannelResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteReleaseChannelRespMultiError(errors)
	}

	return nil
}

// DeleteReleaseChannelRespMultiError is an error wrapping multiple validation
// errors returned by DeleteReleaseChannelResp.ValidateAll() if the designated
// constraints aren't met.
type DeleteReleaseChannelRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteReleaseChannelRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteReleaseChannelRespMultiError) AllErrors() []error { return m }

// DeleteReleaseChannelRespValidationError is the validation error returned by
// DeleteReleaseChannelResp.Validate if the designated constraints aren't met.
type DeleteReleaseChannelRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteReleaseChannelRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteReleaseChannelRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteReleaseChannelRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteReleaseChannelRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteReleaseChannelRespValidationError) ErrorName() string {
	return "DeleteReleaseChannelRespValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteReleaseChannelRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteReleaseChannelResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteReleaseChannelRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteReleaseChannelRespValidationError{}

// Validate checks the field values on ListReleaseChannelsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListReleaseChannelsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReleaseChannelsReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListReleaseChannelsReqMultiError, or nil if none found.
func (m *ListReleaseChannelsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReleaseChannelsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Application

	if len(errors) > 0 {
		return ListReleaseChannelsReqMultiError(errors)
	}

	return nil
}

// ListReleaseChannelsReqMultiError is an error wrapping multiple validation
// errors returned by ListReleaseChannelsReq.ValidateAll() if the designated
// constraints aren't met.
type ListReleaseChannelsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReleaseChannelsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReleaseChannelsReqMultiError) AllErrors() []error { return m }

// ListReleaseChannelsReqValidationError is the validation error returned by
// ListReleaseChannelsReq.Validate if the designated constraints aren't met.
type ListReleaseChannelsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReleaseChannelsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReleaseChannelsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReleaseChannelsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReleaseChannelsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReleaseChannelsReqValidationError) ErrorName() string {
	return "ListReleaseChannelsReqValidationError"
}

// Error satisfies the builtin error interface
func (e ListReleaseChannelsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReleaseChannelsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReleaseChannelsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReleaseChannelsReqValidationError{}

// Validate checks the field values on ListReleaseChannelsResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListReleaseChannelsResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReleaseChannelsResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListReleaseChannelsRespMultiError, or nil if none found.
func (m *ListReleaseChannelsResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReleaseChannelsResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetReleaseChannels() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListReleaseChannelsRespValidationError{
						field:  fmt.Sprintf("ReleaseChannels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListReleaseChannelsRespValidationError{
						field:  fmt.Sprintf("ReleaseChannels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListReleaseChannelsRespValidationError{
					field:  fmt.Sprintf("ReleaseChannels[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListReleaseChannelsRespMultiError(errors)
	}

	return nil
}

// ListReleaseChannelsRespMultiError is an error wrapping multiple validation
// errors returned by ListReleaseChannelsResp.ValidateAll() if the designated
// constraints aren't met.
type ListReleaseChannelsRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReleaseChannelsRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReleaseChannelsRespMultiError) AllErrors() []error { return m }

// ListReleaseChannelsRespValidationError is the validation error returned by
// ListReleaseChannelsResp.Validate if the designated constraints aren't met.
type ListReleaseChannelsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReleaseChannelsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReleaseChannelsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReleaseChannelsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReleaseChannelsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReleaseChannelsRespValidationError) ErrorName() string {
	return "ListReleaseChannelsRespValidationError"
}

// Error satisfies the builtin error interface
func (e ListReleaseChannelsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReleaseChannelsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReleaseChannelsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReleaseChannelsRespValidationError{}

// Validate checks the field values on GetReleaseChannelReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetReleaseChannelReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReleaseChannelReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetReleaseChannelReqMultiError, or nil if none found.
func (m *GetReleaseChannelReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReleaseChannelReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Application

	// no validation rules for ReleaseChannel

	if len(errors) > 0 {
		return GetReleaseChannelReqMultiError(errors)
	}

	return nil
}

// GetReleaseChannelReqMultiError is an error wrapping multiple validation
// errors returned by GetReleaseChannelReq.ValidateAll() if the designated
// constraints aren't met.
type GetReleaseChannelReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReleaseChannelReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReleaseChannelReqMultiError) AllErrors() []error { return m }

// GetReleaseChannelReqValidationError is the validation error returned by
// GetReleaseChannelReq.Validate if the designated constraints aren't met.
type GetReleaseChannelReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReleaseChannelReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReleaseChannelReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReleaseChannelReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReleaseChannelReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReleaseChannelReqValidationError) ErrorName() string {
	return "GetReleaseChannelReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetReleaseChannelReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReleaseChannelReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReleaseChannelReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReleaseChannelReqValidationError{}

// Validate checks the field values on GetReleaseChannelResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetReleaseChannelResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReleaseChannelResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetReleaseChannelRespMultiError, or nil if none found.
func (m *GetReleaseChannelResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReleaseChannelResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetReleaseChannel()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetReleaseChannelRespValidationError{
					field:  "ReleaseChannel",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetReleaseChannelRespValidationError{
					field:  "ReleaseChannel",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReleaseChannel()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetReleaseChannelRespValidationError{
				field:  "ReleaseChannel",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetReleaseChannelRespMultiError(errors)
	}

	return nil
}

// GetReleaseChannelRespMultiError is an error wrapping multiple validation
// errors returned by GetReleaseChannelResp.ValidateAll() if the designated
// constraints aren't met.
type GetReleaseChannelRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReleaseChannelRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReleaseChannelRespMultiError) AllErrors() []error { return m }

// GetReleaseChannelRespValidationError is the validation error returned by
// GetReleaseChannelResp.Validate if the designated constraints aren't met.
type GetReleaseChannelRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReleaseChannelRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReleaseChannelRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReleaseChannelRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReleaseChannelRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReleaseChannelRespValidationError) ErrorName() string {
	return "GetReleaseChannelRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetReleaseChannelRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReleaseChannelResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReleaseChannelRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReleaseChannelRespValidationError{}

// Validate checks the field values on GetReleaseChannelEventsReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetReleaseChannelEventsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReleaseChannelEventsReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetReleaseChannelEventsReqMultiError, or nil if none found.
func (m *GetReleaseChannelEventsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReleaseChannelEventsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetApplication()) < 1 {
		err := GetReleaseChannelEventsReqValidationError{
			field:  "Application",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetReleaseChannel()) < 1 {
		err := GetReleaseChannelEventsReqValidationError{
			field:  "ReleaseChannel",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PageToken

	// no validation rules for PageSize

	// no validation rules for OrderByDescTimestamp

	if len(errors) > 0 {
		return GetReleaseChannelEventsReqMultiError(errors)
	}

	return nil
}

// GetReleaseChannelEventsReqMultiError is an error wrapping multiple
// validation errors returned by GetReleaseChannelEventsReq.ValidateAll() if
// the designated constraints aren't met.
type GetReleaseChannelEventsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReleaseChannelEventsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReleaseChannelEventsReqMultiError) AllErrors() []error { return m }

// GetReleaseChannelEventsReqValidationError is the validation error returned
// by GetReleaseChannelEventsReq.Validate if the designated constraints aren't met.
type GetReleaseChannelEventsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReleaseChannelEventsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReleaseChannelEventsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReleaseChannelEventsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReleaseChannelEventsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReleaseChannelEventsReqValidationError) ErrorName() string {
	return "GetReleaseChannelEventsReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetReleaseChannelEventsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReleaseChannelEventsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReleaseChannelEventsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReleaseChannelEventsReqValidationError{}

// Validate checks the field values on GetReleaseChannelEventsResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetReleaseChannelEventsResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReleaseChannelEventsResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetReleaseChannelEventsRespMultiError, or nil if none found.
func (m *GetReleaseChannelEventsResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReleaseChannelEventsResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetReleaseChannelEventsRespValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetReleaseChannelEventsRespValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetReleaseChannelEventsRespValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return GetReleaseChannelEventsRespMultiError(errors)
	}

	return nil
}

// GetReleaseChannelEventsRespMultiError is an error wrapping multiple
// validation errors returned by GetReleaseChannelEventsResp.ValidateAll() if
// the designated constraints aren't met.
type GetReleaseChannelEventsRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReleaseChannelEventsRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReleaseChannelEventsRespMultiError) AllErrors() []error { return m }

// GetReleaseChannelEventsRespValidationError is the validation error returned
// by GetReleaseChannelEventsResp.Validate if the designated constraints
// aren't met.
type GetReleaseChannelEventsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReleaseChannelEventsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReleaseChannelEventsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReleaseChannelEventsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReleaseChannelEventsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReleaseChannelEventsRespValidationError) ErrorName() string {
	return "GetReleaseChannelEventsRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetReleaseChannelEventsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReleaseChannelEventsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReleaseChannelEventsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReleaseChannelEventsRespValidationError{}
