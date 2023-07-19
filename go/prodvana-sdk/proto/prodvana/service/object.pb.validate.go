// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/service/object.proto

package service

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

// Validate checks the field values on ServiceInstanceProtectionAttachment with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *ServiceInstanceProtectionAttachment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceInstanceProtectionAttachment
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// ServiceInstanceProtectionAttachmentMultiError, or nil if none found.
func (m *ServiceInstanceProtectionAttachment) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceInstanceProtectionAttachment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Protection

	// no validation rules for Attachment

	// no validation rules for DesiredStateId

	// no validation rules for AttachmentId

	if len(errors) > 0 {
		return ServiceInstanceProtectionAttachmentMultiError(errors)
	}

	return nil
}

// ServiceInstanceProtectionAttachmentMultiError is an error wrapping multiple
// validation errors returned by
// ServiceInstanceProtectionAttachment.ValidateAll() if the designated
// constraints aren't met.
type ServiceInstanceProtectionAttachmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceInstanceProtectionAttachmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceInstanceProtectionAttachmentMultiError) AllErrors() []error { return m }

// ServiceInstanceProtectionAttachmentValidationError is the validation error
// returned by ServiceInstanceProtectionAttachment.Validate if the designated
// constraints aren't met.
type ServiceInstanceProtectionAttachmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceInstanceProtectionAttachmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceInstanceProtectionAttachmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceInstanceProtectionAttachmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceInstanceProtectionAttachmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceInstanceProtectionAttachmentValidationError) ErrorName() string {
	return "ServiceInstanceProtectionAttachmentValidationError"
}

// Error satisfies the builtin error interface
func (e ServiceInstanceProtectionAttachmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceInstanceProtectionAttachment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceInstanceProtectionAttachmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceInstanceProtectionAttachmentValidationError{}

// Validate checks the field values on ExternalAddr with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExternalAddr) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExternalAddr with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExternalAddrMultiError, or
// nil if none found.
func (m *ExternalAddr) ValidateAll() error {
	return m.validate(true)
}

func (m *ExternalAddr) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Addr

	// no validation rules for Type

	if len(errors) > 0 {
		return ExternalAddrMultiError(errors)
	}

	return nil
}

// ExternalAddrMultiError is an error wrapping multiple validation errors
// returned by ExternalAddr.ValidateAll() if the designated constraints aren't met.
type ExternalAddrMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExternalAddrMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExternalAddrMultiError) AllErrors() []error { return m }

// ExternalAddrValidationError is the validation error returned by
// ExternalAddr.Validate if the designated constraints aren't met.
type ExternalAddrValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExternalAddrValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExternalAddrValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExternalAddrValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExternalAddrValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExternalAddrValidationError) ErrorName() string { return "ExternalAddrValidationError" }

// Error satisfies the builtin error interface
func (e ExternalAddrValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExternalAddr.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExternalAddrValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExternalAddrValidationError{}

// Validate checks the field values on ServiceState with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServiceState) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceState with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServiceStateMultiError, or
// nil if none found.
func (m *ServiceState) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceState) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ServiceStateMultiError(errors)
	}

	return nil
}

// ServiceStateMultiError is an error wrapping multiple validation errors
// returned by ServiceState.ValidateAll() if the designated constraints aren't met.
type ServiceStateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceStateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceStateMultiError) AllErrors() []error { return m }

// ServiceStateValidationError is the validation error returned by
// ServiceState.Validate if the designated constraints aren't met.
type ServiceStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceStateValidationError) ErrorName() string { return "ServiceStateValidationError" }

// Error satisfies the builtin error interface
func (e ServiceStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceStateValidationError{}

// Validate checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Service) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ServiceMultiError, or nil if none found.
func (m *Service) ValidateAll() error {
	return m.validate(true)
}

func (m *Service) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "Meta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "Meta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceValidationError{
				field:  "Meta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetState()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "State",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "State",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetState()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceValidationError{
				field:  "State",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUserMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "UserMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceValidationError{
					field:  "UserMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceValidationError{
				field:  "UserMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServiceMultiError(errors)
	}

	return nil
}

// ServiceMultiError is an error wrapping multiple validation errors returned
// by Service.ValidateAll() if the designated constraints aren't met.
type ServiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceMultiError) AllErrors() []error { return m }

// ServiceValidationError is the validation error returned by Service.Validate
// if the designated constraints aren't met.
type ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceValidationError) ErrorName() string { return "ServiceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceValidationError{}

// Validate checks the field values on ServiceInstanceState with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ServiceInstanceState) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceInstanceState with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServiceInstanceStateMultiError, or nil if none found.
func (m *ServiceInstanceState) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceInstanceState) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetExternalAddrs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServiceInstanceStateValidationError{
						field:  fmt.Sprintf("ExternalAddrs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServiceInstanceStateValidationError{
						field:  fmt.Sprintf("ExternalAddrs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServiceInstanceStateValidationError{
					field:  fmt.Sprintf("ExternalAddrs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetProtectionAttachments() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServiceInstanceStateValidationError{
						field:  fmt.Sprintf("ProtectionAttachments[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServiceInstanceStateValidationError{
						field:  fmt.Sprintf("ProtectionAttachments[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServiceInstanceStateValidationError{
					field:  fmt.Sprintf("ProtectionAttachments[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ServiceInstanceStateMultiError(errors)
	}

	return nil
}

// ServiceInstanceStateMultiError is an error wrapping multiple validation
// errors returned by ServiceInstanceState.ValidateAll() if the designated
// constraints aren't met.
type ServiceInstanceStateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceInstanceStateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceInstanceStateMultiError) AllErrors() []error { return m }

// ServiceInstanceStateValidationError is the validation error returned by
// ServiceInstanceState.Validate if the designated constraints aren't met.
type ServiceInstanceStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceInstanceStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceInstanceStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceInstanceStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceInstanceStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceInstanceStateValidationError) ErrorName() string {
	return "ServiceInstanceStateValidationError"
}

// Error satisfies the builtin error interface
func (e ServiceInstanceStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceInstanceState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceInstanceStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceInstanceStateValidationError{}

// Validate checks the field values on ServiceInstance with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ServiceInstance) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceInstance with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServiceInstanceMultiError, or nil if none found.
func (m *ServiceInstance) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceInstance) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceInstanceValidationError{
					field:  "Meta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceInstanceValidationError{
					field:  "Meta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceInstanceValidationError{
				field:  "Meta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceInstanceValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceInstanceValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceInstanceValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetState()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceInstanceValidationError{
					field:  "State",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceInstanceValidationError{
					field:  "State",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetState()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceInstanceValidationError{
				field:  "State",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServiceInstanceMultiError(errors)
	}

	return nil
}

// ServiceInstanceMultiError is an error wrapping multiple validation errors
// returned by ServiceInstance.ValidateAll() if the designated constraints
// aren't met.
type ServiceInstanceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceInstanceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceInstanceMultiError) AllErrors() []error { return m }

// ServiceInstanceValidationError is the validation error returned by
// ServiceInstance.Validate if the designated constraints aren't met.
type ServiceInstanceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceInstanceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceInstanceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceInstanceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceInstanceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceInstanceValidationError) ErrorName() string { return "ServiceInstanceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceInstanceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceInstance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceInstanceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceInstanceValidationError{}
