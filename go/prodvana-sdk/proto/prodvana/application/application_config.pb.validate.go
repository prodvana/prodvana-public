// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/application/application_config.proto

package application

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

// Validate checks the field values on AnnotationsConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AnnotationsConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AnnotationsConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AnnotationsConfigMultiError, or nil if none found.
func (m *AnnotationsConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *AnnotationsConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLast9()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AnnotationsConfigValidationError{
					field:  "Last9",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AnnotationsConfigValidationError{
					field:  "Last9",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLast9()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AnnotationsConfigValidationError{
				field:  "Last9",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AnnotationsConfigMultiError(errors)
	}

	return nil
}

// AnnotationsConfigMultiError is an error wrapping multiple validation errors
// returned by AnnotationsConfig.ValidateAll() if the designated constraints
// aren't met.
type AnnotationsConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AnnotationsConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AnnotationsConfigMultiError) AllErrors() []error { return m }

// AnnotationsConfigValidationError is the validation error returned by
// AnnotationsConfig.Validate if the designated constraints aren't met.
type AnnotationsConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AnnotationsConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AnnotationsConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AnnotationsConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AnnotationsConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AnnotationsConfigValidationError) ErrorName() string {
	return "AnnotationsConfigValidationError"
}

// Error satisfies the builtin error interface
func (e AnnotationsConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAnnotationsConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AnnotationsConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AnnotationsConfigValidationError{}

// Validate checks the field values on ApplicationConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ApplicationConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApplicationConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ApplicationConfigMultiError, or nil if none found.
func (m *ApplicationConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ApplicationConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 63 {
		err := ApplicationConfigValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 63 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_ApplicationConfig_Name_Pattern.MatchString(m.GetName()) {
		err := ApplicationConfigValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetReleaseChannels() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ApplicationConfigValidationError{
						field:  fmt.Sprintf("ReleaseChannels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ApplicationConfigValidationError{
						field:  fmt.Sprintf("ReleaseChannels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationConfigValidationError{
					field:  fmt.Sprintf("ReleaseChannels[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetReleaseChannelGroups() {
		_, _ = idx, item

		if item == nil {
			err := ApplicationConfigValidationError{
				field:  fmt.Sprintf("ReleaseChannelGroups[%v]", idx),
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
					errors = append(errors, ApplicationConfigValidationError{
						field:  fmt.Sprintf("ReleaseChannelGroups[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ApplicationConfigValidationError{
						field:  fmt.Sprintf("ReleaseChannelGroups[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationConfigValidationError{
					field:  fmt.Sprintf("ReleaseChannelGroups[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetNotifications()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationConfigValidationError{
					field:  "Notifications",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationConfigValidationError{
					field:  "Notifications",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNotifications()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationConfigValidationError{
				field:  "Notifications",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAnnotations()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationConfigValidationError{
					field:  "Annotations",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationConfigValidationError{
					field:  "Annotations",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAnnotations()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationConfigValidationError{
				field:  "Annotations",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetReleaseSettings()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationConfigValidationError{
					field:  "ReleaseSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationConfigValidationError{
					field:  "ReleaseSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReleaseSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationConfigValidationError{
				field:  "ReleaseSettings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAlerts()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApplicationConfigValidationError{
					field:  "Alerts",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApplicationConfigValidationError{
					field:  "Alerts",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlerts()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApplicationConfigValidationError{
				field:  "Alerts",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetCapabilities() {
		_, _ = idx, item

		if item == nil {
			err := ApplicationConfigValidationError{
				field:  fmt.Sprintf("Capabilities[%v]", idx),
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
					errors = append(errors, ApplicationConfigValidationError{
						field:  fmt.Sprintf("Capabilities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ApplicationConfigValidationError{
						field:  fmt.Sprintf("Capabilities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationConfigValidationError{
					field:  fmt.Sprintf("Capabilities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetCapabilityInstances() {
		_, _ = idx, item

		if item == nil {
			err := ApplicationConfigValidationError{
				field:  fmt.Sprintf("CapabilityInstances[%v]", idx),
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
					errors = append(errors, ApplicationConfigValidationError{
						field:  fmt.Sprintf("CapabilityInstances[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ApplicationConfigValidationError{
						field:  fmt.Sprintf("CapabilityInstances[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ApplicationConfigValidationError{
					field:  fmt.Sprintf("CapabilityInstances[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ApplicationConfigMultiError(errors)
	}

	return nil
}

// ApplicationConfigMultiError is an error wrapping multiple validation errors
// returned by ApplicationConfig.ValidateAll() if the designated constraints
// aren't met.
type ApplicationConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApplicationConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApplicationConfigMultiError) AllErrors() []error { return m }

// ApplicationConfigValidationError is the validation error returned by
// ApplicationConfig.Validate if the designated constraints aren't met.
type ApplicationConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplicationConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplicationConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplicationConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplicationConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplicationConfigValidationError) ErrorName() string {
	return "ApplicationConfigValidationError"
}

// Error satisfies the builtin error interface
func (e ApplicationConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplicationConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplicationConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplicationConfigValidationError{}

var _ApplicationConfig_Name_Pattern = regexp.MustCompile("^[a-z]([a-z0-9-]*[a-z0-9]){0,1}$")

// Validate checks the field values on AnnotationsConfig_Last9 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AnnotationsConfig_Last9) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AnnotationsConfig_Last9 with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AnnotationsConfig_Last9MultiError, or nil if none found.
func (m *AnnotationsConfig_Last9) ValidateAll() error {
	return m.validate(true)
}

func (m *AnnotationsConfig_Last9) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDataSource()) < 1 {
		err := AnnotationsConfig_Last9ValidationError{
			field:  "DataSource",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AnnotationsConfig_Last9MultiError(errors)
	}

	return nil
}

// AnnotationsConfig_Last9MultiError is an error wrapping multiple validation
// errors returned by AnnotationsConfig_Last9.ValidateAll() if the designated
// constraints aren't met.
type AnnotationsConfig_Last9MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AnnotationsConfig_Last9MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AnnotationsConfig_Last9MultiError) AllErrors() []error { return m }

// AnnotationsConfig_Last9ValidationError is the validation error returned by
// AnnotationsConfig_Last9.Validate if the designated constraints aren't met.
type AnnotationsConfig_Last9ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AnnotationsConfig_Last9ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AnnotationsConfig_Last9ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AnnotationsConfig_Last9ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AnnotationsConfig_Last9ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AnnotationsConfig_Last9ValidationError) ErrorName() string {
	return "AnnotationsConfig_Last9ValidationError"
}

// Error satisfies the builtin error interface
func (e AnnotationsConfig_Last9ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAnnotationsConfig_Last9.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AnnotationsConfig_Last9ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AnnotationsConfig_Last9ValidationError{}
