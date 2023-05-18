// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/config_file/config.proto

package config_file

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

// Validate checks the field values on ProdvanaConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProdvanaConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProdvanaConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProdvanaConfigMultiError,
// or nil if none found.
func (m *ProdvanaConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ProdvanaConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofConfigOneofPresent := false
	switch v := m.ConfigOneof.(type) {
	case *ProdvanaConfig_Application:
		if v == nil {
			err := ProdvanaConfigValidationError{
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
			switch v := interface{}(m.GetApplication()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "Application",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "Application",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetApplication()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProdvanaConfigValidationError{
					field:  "Application",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProdvanaConfig_Service:
		if v == nil {
			err := ProdvanaConfigValidationError{
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
			switch v := interface{}(m.GetService()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "Service",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "Service",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetService()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProdvanaConfigValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProdvanaConfig_Protection:
		if v == nil {
			err := ProdvanaConfigValidationError{
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
			switch v := interface{}(m.GetProtection()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "Protection",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "Protection",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetProtection()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProdvanaConfigValidationError{
					field:  "Protection",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProdvanaConfig_Runtime:
		if v == nil {
			err := ProdvanaConfigValidationError{
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
			switch v := interface{}(m.GetRuntime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "Runtime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "Runtime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRuntime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProdvanaConfigValidationError{
					field:  "Runtime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProdvanaConfig_DeliveryModule:
		if v == nil {
			err := ProdvanaConfigValidationError{
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
			switch v := interface{}(m.GetDeliveryModule()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "DeliveryModule",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "DeliveryModule",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeliveryModule()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProdvanaConfigValidationError{
					field:  "DeliveryModule",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofConfigOneofPresent {
		err := ProdvanaConfigValidationError{
			field:  "ConfigOneof",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}
	switch v := m.MetadataOneof.(type) {
	case *ProdvanaConfig_ApplicationMetadata:
		if v == nil {
			err := ProdvanaConfigValidationError{
				field:  "MetadataOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetApplicationMetadata()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "ApplicationMetadata",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "ApplicationMetadata",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetApplicationMetadata()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProdvanaConfigValidationError{
					field:  "ApplicationMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ProdvanaConfig_ServiceMetadata:
		if v == nil {
			err := ProdvanaConfigValidationError{
				field:  "MetadataOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetServiceMetadata()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "ServiceMetadata",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProdvanaConfigValidationError{
						field:  "ServiceMetadata",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetServiceMetadata()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProdvanaConfigValidationError{
					field:  "ServiceMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ProdvanaConfigMultiError(errors)
	}

	return nil
}

// ProdvanaConfigMultiError is an error wrapping multiple validation errors
// returned by ProdvanaConfig.ValidateAll() if the designated constraints
// aren't met.
type ProdvanaConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProdvanaConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProdvanaConfigMultiError) AllErrors() []error { return m }

// ProdvanaConfigValidationError is the validation error returned by
// ProdvanaConfig.Validate if the designated constraints aren't met.
type ProdvanaConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProdvanaConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProdvanaConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProdvanaConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProdvanaConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProdvanaConfigValidationError) ErrorName() string { return "ProdvanaConfigValidationError" }

// Error satisfies the builtin error interface
func (e ProdvanaConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProdvanaConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProdvanaConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProdvanaConfigValidationError{}
