// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/protection/builtins.proto

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

// Validate checks the field values on CommitDenylistProtectionConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommitDenylistProtectionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommitDenylistProtectionConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CommitDenylistProtectionConfigMultiError, or nil if none found.
func (m *CommitDenylistProtectionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *CommitDenylistProtectionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetRepository()) < 1 {
		err := CommitDenylistProtectionConfigValidationError{
			field:  "Repository",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetCommits() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommitDenylistProtectionConfigValidationError{
						field:  fmt.Sprintf("Commits[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommitDenylistProtectionConfigValidationError{
						field:  fmt.Sprintf("Commits[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommitDenylistProtectionConfigValidationError{
					field:  fmt.Sprintf("Commits[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CommitDenylistProtectionConfigMultiError(errors)
	}

	return nil
}

// CommitDenylistProtectionConfigMultiError is an error wrapping multiple
// validation errors returned by CommitDenylistProtectionConfig.ValidateAll()
// if the designated constraints aren't met.
type CommitDenylistProtectionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommitDenylistProtectionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommitDenylistProtectionConfigMultiError) AllErrors() []error { return m }

// CommitDenylistProtectionConfigValidationError is the validation error
// returned by CommitDenylistProtectionConfig.Validate if the designated
// constraints aren't met.
type CommitDenylistProtectionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommitDenylistProtectionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommitDenylistProtectionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommitDenylistProtectionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommitDenylistProtectionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommitDenylistProtectionConfigValidationError) ErrorName() string {
	return "CommitDenylistProtectionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e CommitDenylistProtectionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommitDenylistProtectionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommitDenylistProtectionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommitDenylistProtectionConfigValidationError{}

// Validate checks the field values on TimeWindow with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TimeWindow) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TimeWindow with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TimeWindowMultiError, or
// nil if none found.
func (m *TimeWindow) ValidateAll() error {
	return m.validate(true)
}

func (m *TimeWindow) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDays()) < 1 {
		err := TimeWindowValidationError{
			field:  "Days",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetHours()) < 1 {
		err := TimeWindowValidationError{
			field:  "Hours",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Tz

	if len(errors) > 0 {
		return TimeWindowMultiError(errors)
	}

	return nil
}

// TimeWindowMultiError is an error wrapping multiple validation errors
// returned by TimeWindow.ValidateAll() if the designated constraints aren't met.
type TimeWindowMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TimeWindowMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TimeWindowMultiError) AllErrors() []error { return m }

// TimeWindowValidationError is the validation error returned by
// TimeWindow.Validate if the designated constraints aren't met.
type TimeWindowValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimeWindowValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimeWindowValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimeWindowValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimeWindowValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimeWindowValidationError) ErrorName() string { return "TimeWindowValidationError" }

// Error satisfies the builtin error interface
func (e TimeWindowValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimeWindow.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimeWindowValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimeWindowValidationError{}

// Validate checks the field values on AllowedTimesProtectionConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AllowedTimesProtectionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AllowedTimesProtectionConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AllowedTimesProtectionConfigMultiError, or nil if none found.
func (m *AllowedTimesProtectionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *AllowedTimesProtectionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetWindows() {
		_, _ = idx, item

		if item == nil {
			err := AllowedTimesProtectionConfigValidationError{
				field:  fmt.Sprintf("Windows[%v]", idx),
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
					errors = append(errors, AllowedTimesProtectionConfigValidationError{
						field:  fmt.Sprintf("Windows[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AllowedTimesProtectionConfigValidationError{
						field:  fmt.Sprintf("Windows[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AllowedTimesProtectionConfigValidationError{
					field:  fmt.Sprintf("Windows[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AllowedTimesProtectionConfigMultiError(errors)
	}

	return nil
}

// AllowedTimesProtectionConfigMultiError is an error wrapping multiple
// validation errors returned by AllowedTimesProtectionConfig.ValidateAll() if
// the designated constraints aren't met.
type AllowedTimesProtectionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AllowedTimesProtectionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AllowedTimesProtectionConfigMultiError) AllErrors() []error { return m }

// AllowedTimesProtectionConfigValidationError is the validation error returned
// by AllowedTimesProtectionConfig.Validate if the designated constraints
// aren't met.
type AllowedTimesProtectionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AllowedTimesProtectionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AllowedTimesProtectionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AllowedTimesProtectionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AllowedTimesProtectionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AllowedTimesProtectionConfigValidationError) ErrorName() string {
	return "AllowedTimesProtectionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e AllowedTimesProtectionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAllowedTimesProtectionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AllowedTimesProtectionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AllowedTimesProtectionConfigValidationError{}

// Validate checks the field values on
// CommitDenylistProtectionConfig_CommitDefinition with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommitDenylistProtectionConfig_CommitDefinition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CommitDenylistProtectionConfig_CommitDefinition with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// CommitDenylistProtectionConfig_CommitDefinitionMultiError, or nil if none found.
func (m *CommitDenylistProtectionConfig_CommitDefinition) ValidateAll() error {
	return m.validate(true)
}

func (m *CommitDenylistProtectionConfig_CommitDefinition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofCommitOneofPresent := false
	switch v := m.CommitOneof.(type) {
	case *CommitDenylistProtectionConfig_CommitDefinition_Commit:
		if v == nil {
			err := CommitDenylistProtectionConfig_CommitDefinitionValidationError{
				field:  "CommitOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofCommitOneofPresent = true

		if utf8.RuneCountInString(m.GetCommit()) < 1 {
			err := CommitDenylistProtectionConfig_CommitDefinitionValidationError{
				field:  "Commit",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *CommitDenylistProtectionConfig_CommitDefinition_Range_:
		if v == nil {
			err := CommitDenylistProtectionConfig_CommitDefinitionValidationError{
				field:  "CommitOneof",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofCommitOneofPresent = true

		if m.GetRange() == nil {
			err := CommitDenylistProtectionConfig_CommitDefinitionValidationError{
				field:  "Range",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetRange()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommitDenylistProtectionConfig_CommitDefinitionValidationError{
						field:  "Range",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommitDenylistProtectionConfig_CommitDefinitionValidationError{
						field:  "Range",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommitDenylistProtectionConfig_CommitDefinitionValidationError{
					field:  "Range",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofCommitOneofPresent {
		err := CommitDenylistProtectionConfig_CommitDefinitionValidationError{
			field:  "CommitOneof",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CommitDenylistProtectionConfig_CommitDefinitionMultiError(errors)
	}

	return nil
}

// CommitDenylistProtectionConfig_CommitDefinitionMultiError is an error
// wrapping multiple validation errors returned by
// CommitDenylistProtectionConfig_CommitDefinition.ValidateAll() if the
// designated constraints aren't met.
type CommitDenylistProtectionConfig_CommitDefinitionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommitDenylistProtectionConfig_CommitDefinitionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommitDenylistProtectionConfig_CommitDefinitionMultiError) AllErrors() []error { return m }

// CommitDenylistProtectionConfig_CommitDefinitionValidationError is the
// validation error returned by
// CommitDenylistProtectionConfig_CommitDefinition.Validate if the designated
// constraints aren't met.
type CommitDenylistProtectionConfig_CommitDefinitionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommitDenylistProtectionConfig_CommitDefinitionValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e CommitDenylistProtectionConfig_CommitDefinitionValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e CommitDenylistProtectionConfig_CommitDefinitionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommitDenylistProtectionConfig_CommitDefinitionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommitDenylistProtectionConfig_CommitDefinitionValidationError) ErrorName() string {
	return "CommitDenylistProtectionConfig_CommitDefinitionValidationError"
}

// Error satisfies the builtin error interface
func (e CommitDenylistProtectionConfig_CommitDefinitionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommitDenylistProtectionConfig_CommitDefinition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommitDenylistProtectionConfig_CommitDefinitionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommitDenylistProtectionConfig_CommitDefinitionValidationError{}

// Validate checks the field values on
// CommitDenylistProtectionConfig_CommitDefinition_Range with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommitDenylistProtectionConfig_CommitDefinition_Range) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CommitDenylistProtectionConfig_CommitDefinition_Range with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommitDenylistProtectionConfig_CommitDefinition_RangeMultiError, or nil if
// none found.
func (m *CommitDenylistProtectionConfig_CommitDefinition_Range) ValidateAll() error {
	return m.validate(true)
}

func (m *CommitDenylistProtectionConfig_CommitDefinition_Range) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStartCommit()) < 1 {
		err := CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError{
			field:  "StartCommit",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetEndCommit()) < 1 {
		err := CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError{
			field:  "EndCommit",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CommitDenylistProtectionConfig_CommitDefinition_RangeMultiError(errors)
	}

	return nil
}

// CommitDenylistProtectionConfig_CommitDefinition_RangeMultiError is an error
// wrapping multiple validation errors returned by
// CommitDenylistProtectionConfig_CommitDefinition_Range.ValidateAll() if the
// designated constraints aren't met.
type CommitDenylistProtectionConfig_CommitDefinition_RangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommitDenylistProtectionConfig_CommitDefinition_RangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommitDenylistProtectionConfig_CommitDefinition_RangeMultiError) AllErrors() []error {
	return m
}

// CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError is the
// validation error returned by
// CommitDenylistProtectionConfig_CommitDefinition_Range.Validate if the
// designated constraints aren't met.
type CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError) Key() bool {
	return e.key
}

// ErrorName returns error name.
func (e CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError) ErrorName() string {
	return "CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError"
}

// Error satisfies the builtin error interface
func (e CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommitDenylistProtectionConfig_CommitDefinition_Range.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommitDenylistProtectionConfig_CommitDefinition_RangeValidationError{}
