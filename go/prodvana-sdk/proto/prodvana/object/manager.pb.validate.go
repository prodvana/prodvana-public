// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/object/manager.proto

package object

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

// Validate checks the field values on QueryObjectReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QueryObjectReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryObjectReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QueryObjectReqMultiError,
// or nil if none found.
func (m *QueryObjectReq) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryObjectReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetQuery()) < 1 {
		err := QueryObjectReqValidationError{
			field:  "Query",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return QueryObjectReqMultiError(errors)
	}

	return nil
}

// QueryObjectReqMultiError is an error wrapping multiple validation errors
// returned by QueryObjectReq.ValidateAll() if the designated constraints
// aren't met.
type QueryObjectReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryObjectReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryObjectReqMultiError) AllErrors() []error { return m }

// QueryObjectReqValidationError is the validation error returned by
// QueryObjectReq.Validate if the designated constraints aren't met.
type QueryObjectReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryObjectReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryObjectReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryObjectReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryObjectReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryObjectReqValidationError) ErrorName() string { return "QueryObjectReqValidationError" }

// Error satisfies the builtin error interface
func (e QueryObjectReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryObjectReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryObjectReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryObjectReqValidationError{}

// Validate checks the field values on QueryObjectResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *QueryObjectResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryObjectResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryObjectRespMultiError, or nil if none found.
func (m *QueryObjectResp) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryObjectResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetObjects() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QueryObjectRespValidationError{
						field:  fmt.Sprintf("Objects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QueryObjectRespValidationError{
						field:  fmt.Sprintf("Objects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryObjectRespValidationError{
					field:  fmt.Sprintf("Objects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QueryObjectRespMultiError(errors)
	}

	return nil
}

// QueryObjectRespMultiError is an error wrapping multiple validation errors
// returned by QueryObjectResp.ValidateAll() if the designated constraints
// aren't met.
type QueryObjectRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryObjectRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryObjectRespMultiError) AllErrors() []error { return m }

// QueryObjectRespValidationError is the validation error returned by
// QueryObjectResp.Validate if the designated constraints aren't met.
type QueryObjectRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryObjectRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryObjectRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryObjectRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryObjectRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryObjectRespValidationError) ErrorName() string { return "QueryObjectRespValidationError" }

// Error satisfies the builtin error interface
func (e QueryObjectRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryObjectResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryObjectRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryObjectRespValidationError{}

// Validate checks the field values on GetLabelsReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetLabelsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLabelsReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetLabelsReqMultiError, or
// nil if none found.
func (m *GetLabelsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLabelsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Id

	if len(errors) > 0 {
		return GetLabelsReqMultiError(errors)
	}

	return nil
}

// GetLabelsReqMultiError is an error wrapping multiple validation errors
// returned by GetLabelsReq.ValidateAll() if the designated constraints aren't met.
type GetLabelsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLabelsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLabelsReqMultiError) AllErrors() []error { return m }

// GetLabelsReqValidationError is the validation error returned by
// GetLabelsReq.Validate if the designated constraints aren't met.
type GetLabelsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLabelsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLabelsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLabelsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLabelsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLabelsReqValidationError) ErrorName() string { return "GetLabelsReqValidationError" }

// Error satisfies the builtin error interface
func (e GetLabelsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLabelsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLabelsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLabelsReqValidationError{}

// Validate checks the field values on GetLabelsResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetLabelsResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLabelsResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetLabelsRespMultiError, or
// nil if none found.
func (m *GetLabelsResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLabelsResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetLabels() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetLabelsRespValidationError{
						field:  fmt.Sprintf("Labels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetLabelsRespValidationError{
						field:  fmt.Sprintf("Labels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetLabelsRespValidationError{
					field:  fmt.Sprintf("Labels[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetLabelsRespMultiError(errors)
	}

	return nil
}

// GetLabelsRespMultiError is an error wrapping multiple validation errors
// returned by GetLabelsResp.ValidateAll() if the designated constraints
// aren't met.
type GetLabelsRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLabelsRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLabelsRespMultiError) AllErrors() []error { return m }

// GetLabelsRespValidationError is the validation error returned by
// GetLabelsResp.Validate if the designated constraints aren't met.
type GetLabelsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLabelsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLabelsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLabelsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLabelsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLabelsRespValidationError) ErrorName() string { return "GetLabelsRespValidationError" }

// Error satisfies the builtin error interface
func (e GetLabelsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLabelsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLabelsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLabelsRespValidationError{}