// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/desired_state/model/entity.proto

package model

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

	common_config "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/common_config"
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

	_ = common_config.TaskLifecycle(0)
)

// Validate checks the field values on Entity with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Entity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Entity with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EntityMultiError, or nil if none found.
func (m *Entity) ValidateAll() error {
	return m.validate(true)
}

func (m *Entity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DesiredStateId

	// no validation rules for RootDesiredStateId

	if all {
		switch v := interface{}(m.GetAnnotations()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "Annotations",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "Annotations",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAnnotations()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityValidationError{
				field:  "Annotations",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	// no validation rules for SimpleStatus

	if all {
		switch v := interface{}(m.GetStartingState()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "StartingState",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "StartingState",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStartingState()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityValidationError{
				field:  "StartingState",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLastSeenState()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "LastSeenState",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "LastSeenState",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastSeenState()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityValidationError{
				field:  "LastSeenState",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDesiredState()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "DesiredState",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "DesiredState",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDesiredState()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityValidationError{
				field:  "DesiredState",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTargetState()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "TargetState",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "TargetState",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTargetState()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityValidationError{
				field:  "TargetState",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetPreconditionStatuses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EntityValidationError{
						field:  fmt.Sprintf("PreconditionStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EntityValidationError{
						field:  fmt.Sprintf("PreconditionStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityValidationError{
					field:  fmt.Sprintf("PreconditionStatuses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetStatusExplanations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EntityValidationError{
						field:  fmt.Sprintf("StatusExplanations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EntityValidationError{
						field:  fmt.Sprintf("StatusExplanations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityValidationError{
					field:  fmt.Sprintf("StatusExplanations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetLogs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EntityValidationError{
						field:  fmt.Sprintf("Logs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EntityValidationError{
						field:  fmt.Sprintf("Logs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityValidationError{
					field:  fmt.Sprintf("Logs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetActionExplanation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "ActionExplanation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "ActionExplanation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetActionExplanation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityValidationError{
				field:  "ActionExplanation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLastUpdateTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "LastUpdateTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "LastUpdateTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastUpdateTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityValidationError{
				field:  "LastUpdateTimestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLastFetchedTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "LastFetchedTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "LastFetchedTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastFetchedTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityValidationError{
				field:  "LastFetchedTimestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLastAppliedTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "LastAppliedTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityValidationError{
					field:  "LastAppliedTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastAppliedTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityValidationError{
				field:  "LastAppliedTimestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetDependencies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EntityValidationError{
						field:  fmt.Sprintf("Dependencies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EntityValidationError{
						field:  fmt.Sprintf("Dependencies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityValidationError{
					field:  fmt.Sprintf("Dependencies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Lifecycle

	if len(errors) > 0 {
		return EntityMultiError(errors)
	}

	return nil
}

// EntityMultiError is an error wrapping multiple validation errors returned by
// Entity.ValidateAll() if the designated constraints aren't met.
type EntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntityMultiError) AllErrors() []error { return m }

// EntityValidationError is the validation error returned by Entity.Validate if
// the designated constraints aren't met.
type EntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityValidationError) ErrorName() string { return "EntityValidationError" }

// Error satisfies the builtin error interface
func (e EntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityValidationError{}

// Validate checks the field values on EntityGraph with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EntityGraph) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EntityGraph with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EntityGraphMultiError, or
// nil if none found.
func (m *EntityGraph) ValidateAll() error {
	return m.validate(true)
}

func (m *EntityGraph) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRoot()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityGraphValidationError{
					field:  "Root",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityGraphValidationError{
					field:  "Root",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRoot()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityGraphValidationError{
				field:  "Root",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetEntities() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EntityGraphValidationError{
						field:  fmt.Sprintf("Entities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EntityGraphValidationError{
						field:  fmt.Sprintf("Entities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityGraphValidationError{
					field:  fmt.Sprintf("Entities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EntityGraphMultiError(errors)
	}

	return nil
}

// EntityGraphMultiError is an error wrapping multiple validation errors
// returned by EntityGraph.ValidateAll() if the designated constraints aren't met.
type EntityGraphMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntityGraphMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntityGraphMultiError) AllErrors() []error { return m }

// EntityGraphValidationError is the validation error returned by
// EntityGraph.Validate if the designated constraints aren't met.
type EntityGraphValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityGraphValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityGraphValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityGraphValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityGraphValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityGraphValidationError) ErrorName() string { return "EntityGraphValidationError" }

// Error satisfies the builtin error interface
func (e EntityGraphValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntityGraph.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityGraphValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityGraphValidationError{}
