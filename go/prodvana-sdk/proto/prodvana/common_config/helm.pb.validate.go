// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/common_config/helm.proto

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

// Validate checks the field values on RemoteHelmChart with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RemoteHelmChart) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoteHelmChart with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoteHelmChartMultiError, or nil if none found.
func (m *RemoteHelmChart) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoteHelmChart) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetChart()) < 1 {
		err := RemoteHelmChartValidationError{
			field:  "Chart",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetChartVersion()) < 1 {
		err := RemoteHelmChartValidationError{
			field:  "ChartVersion",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	oneofRepoOneofPresent := false
	switch v := m.RepoOneof.(type) {
	case *RemoteHelmChart_Repo:
		if v == nil {
			err := RemoteHelmChartValidationError{
				field:  "RepoOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofRepoOneofPresent = true
		// no validation rules for Repo
	default:
		_ = v // ensures v is used
	}
	if !oneofRepoOneofPresent {
		err := RemoteHelmChartValidationError{
			field:  "RepoOneof",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RemoteHelmChartMultiError(errors)
	}

	return nil
}

// RemoteHelmChartMultiError is an error wrapping multiple validation errors
// returned by RemoteHelmChart.ValidateAll() if the designated constraints
// aren't met.
type RemoteHelmChartMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoteHelmChartMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoteHelmChartMultiError) AllErrors() []error { return m }

// RemoteHelmChartValidationError is the validation error returned by
// RemoteHelmChart.Validate if the designated constraints aren't met.
type RemoteHelmChartValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoteHelmChartValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoteHelmChartValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoteHelmChartValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoteHelmChartValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoteHelmChartValidationError) ErrorName() string { return "RemoteHelmChartValidationError" }

// Error satisfies the builtin error interface
func (e RemoteHelmChartValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoteHelmChart.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoteHelmChartValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoteHelmChartValidationError{}

// Validate checks the field values on HelmValuesOverrides with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HelmValuesOverrides) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HelmValuesOverrides with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HelmValuesOverridesMultiError, or nil if none found.
func (m *HelmValuesOverrides) ValidateAll() error {
	return m.validate(true)
}

func (m *HelmValuesOverrides) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofOverrideOneofPresent := false
	switch v := m.OverrideOneof.(type) {
	case *HelmValuesOverrides_Inlined:
		if v == nil {
			err := HelmValuesOverridesValidationError{
				field:  "OverrideOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofOverrideOneofPresent = true

		if utf8.RuneCountInString(m.GetInlined()) < 1 {
			err := HelmValuesOverridesValidationError{
				field:  "Inlined",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *HelmValuesOverrides_Local:
		if v == nil {
			err := HelmValuesOverridesValidationError{
				field:  "OverrideOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofOverrideOneofPresent = true

		if all {
			switch v := interface{}(m.GetLocal()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HelmValuesOverridesValidationError{
						field:  "Local",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HelmValuesOverridesValidationError{
						field:  "Local",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLocal()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HelmValuesOverridesValidationError{
					field:  "Local",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HelmValuesOverrides_Remote:
		if v == nil {
			err := HelmValuesOverridesValidationError{
				field:  "OverrideOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofOverrideOneofPresent = true

		if all {
			switch v := interface{}(m.GetRemote()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HelmValuesOverridesValidationError{
						field:  "Remote",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HelmValuesOverridesValidationError{
						field:  "Remote",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRemote()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HelmValuesOverridesValidationError{
					field:  "Remote",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofOverrideOneofPresent {
		err := HelmValuesOverridesValidationError{
			field:  "OverrideOneof",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HelmValuesOverridesMultiError(errors)
	}

	return nil
}

// HelmValuesOverridesMultiError is an error wrapping multiple validation
// errors returned by HelmValuesOverrides.ValidateAll() if the designated
// constraints aren't met.
type HelmValuesOverridesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelmValuesOverridesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelmValuesOverridesMultiError) AllErrors() []error { return m }

// HelmValuesOverridesValidationError is the validation error returned by
// HelmValuesOverrides.Validate if the designated constraints aren't met.
type HelmValuesOverridesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelmValuesOverridesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelmValuesOverridesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelmValuesOverridesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelmValuesOverridesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelmValuesOverridesValidationError) ErrorName() string {
	return "HelmValuesOverridesValidationError"
}

// Error satisfies the builtin error interface
func (e HelmValuesOverridesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelmValuesOverrides.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelmValuesOverridesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelmValuesOverridesValidationError{}

// Validate checks the field values on HelmConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HelmConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HelmConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HelmConfigMultiError, or
// nil if none found.
func (m *HelmConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *HelmConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetValuesOverrides() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HelmConfigValidationError{
						field:  fmt.Sprintf("ValuesOverrides[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HelmConfigValidationError{
						field:  fmt.Sprintf("ValuesOverrides[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HelmConfigValidationError{
					field:  fmt.Sprintf("ValuesOverrides[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for ReleaseName

	if utf8.RuneCountInString(m.GetNamespace()) > 0 {
		err := HelmConfigValidationError{
			field:  "Namespace",
			reason: "value length must be at most 0 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for FixupOwnership

	oneofChartOneofPresent := false
	switch v := m.ChartOneof.(type) {
	case *HelmConfig_Remote:
		if v == nil {
			err := HelmConfigValidationError{
				field:  "ChartOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofChartOneofPresent = true

		if all {
			switch v := interface{}(m.GetRemote()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HelmConfigValidationError{
						field:  "Remote",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HelmConfigValidationError{
						field:  "Remote",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRemote()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HelmConfigValidationError{
					field:  "Remote",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HelmConfig_Local:
		if v == nil {
			err := HelmConfigValidationError{
				field:  "ChartOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofChartOneofPresent = true

		if all {
			switch v := interface{}(m.GetLocal()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HelmConfigValidationError{
						field:  "Local",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HelmConfigValidationError{
						field:  "Local",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLocal()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HelmConfigValidationError{
					field:  "Local",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HelmConfig_HelmTarballBlobId:
		if v == nil {
			err := HelmConfigValidationError{
				field:  "ChartOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofChartOneofPresent = true
		// no validation rules for HelmTarballBlobId
	default:
		_ = v // ensures v is used
	}
	if !oneofChartOneofPresent {
		err := HelmConfigValidationError{
			field:  "ChartOneof",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HelmConfigMultiError(errors)
	}

	return nil
}

// HelmConfigMultiError is an error wrapping multiple validation errors
// returned by HelmConfig.ValidateAll() if the designated constraints aren't met.
type HelmConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelmConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelmConfigMultiError) AllErrors() []error { return m }

// HelmConfigValidationError is the validation error returned by
// HelmConfig.Validate if the designated constraints aren't met.
type HelmConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelmConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelmConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelmConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelmConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelmConfigValidationError) ErrorName() string { return "HelmConfigValidationError" }

// Error satisfies the builtin error interface
func (e HelmConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelmConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelmConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelmConfigValidationError{}
