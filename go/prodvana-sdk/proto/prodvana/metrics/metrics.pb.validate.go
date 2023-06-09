// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/metrics/metrics.proto

package metrics

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

// Validate checks the field values on DeploymentMetrics with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeploymentMetrics) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeploymentMetrics with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeploymentMetricsMultiError, or nil if none found.
func (m *DeploymentMetrics) ValidateAll() error {
	return m.validate(true)
}

func (m *DeploymentMetrics) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetValues()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeploymentMetricsValidationError{
					field:  "Values",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeploymentMetricsValidationError{
					field:  "Values",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValues()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeploymentMetricsValidationError{
				field:  "Values",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetTimeseries() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DeploymentMetricsValidationError{
						field:  fmt.Sprintf("Timeseries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DeploymentMetricsValidationError{
						field:  fmt.Sprintf("Timeseries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeploymentMetricsValidationError{
					field:  fmt.Sprintf("Timeseries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetDeltas()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeploymentMetricsValidationError{
					field:  "Deltas",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeploymentMetricsValidationError{
					field:  "Deltas",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeltas()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeploymentMetricsValidationError{
				field:  "Deltas",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeploymentMetricsMultiError(errors)
	}

	return nil
}

// DeploymentMetricsMultiError is an error wrapping multiple validation errors
// returned by DeploymentMetrics.ValidateAll() if the designated constraints
// aren't met.
type DeploymentMetricsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeploymentMetricsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeploymentMetricsMultiError) AllErrors() []error { return m }

// DeploymentMetricsValidationError is the validation error returned by
// DeploymentMetrics.Validate if the designated constraints aren't met.
type DeploymentMetricsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeploymentMetricsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeploymentMetricsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeploymentMetricsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeploymentMetricsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeploymentMetricsValidationError) ErrorName() string {
	return "DeploymentMetricsValidationError"
}

// Error satisfies the builtin error interface
func (e DeploymentMetricsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeploymentMetrics.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeploymentMetricsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeploymentMetricsValidationError{}

// Validate checks the field values on Cost with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Cost) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cost with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CostMultiError, or nil if none found.
func (m *Cost) ValidateAll() error {
	return m.validate(true)
}

func (m *Cost) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Configured

	// no validation rules for Cost

	if len(errors) > 0 {
		return CostMultiError(errors)
	}

	return nil
}

// CostMultiError is an error wrapping multiple validation errors returned by
// Cost.ValidateAll() if the designated constraints aren't met.
type CostMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CostMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CostMultiError) AllErrors() []error { return m }

// CostValidationError is the validation error returned by Cost.Validate if the
// designated constraints aren't met.
type CostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CostValidationError) ErrorName() string { return "CostValidationError" }

// Error satisfies the builtin error interface
func (e CostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CostValidationError{}

// Validate checks the field values on DeploymentMetrics_Deltas with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeploymentMetrics_Deltas) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeploymentMetrics_Deltas with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeploymentMetrics_DeltasMultiError, or nil if none found.
func (m *DeploymentMetrics_Deltas) ValidateAll() error {
	return m.validate(true)
}

func (m *DeploymentMetrics_Deltas) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalDeployments

	// no validation rules for SuccessfulDeployments

	// no validation rules for SuccessfulDeploymentsPerDay

	// no validation rules for FailedDeployments

	// no validation rules for MedianSuccessfulDeploymentDuration

	// no validation rules for DeploymentFailureRate

	// no validation rules for TotalRollbacks

	// no validation rules for SuccessfulRollbacks

	// no validation rules for MedianSuccessfulRollbackDuration

	if len(errors) > 0 {
		return DeploymentMetrics_DeltasMultiError(errors)
	}

	return nil
}

// DeploymentMetrics_DeltasMultiError is an error wrapping multiple validation
// errors returned by DeploymentMetrics_Deltas.ValidateAll() if the designated
// constraints aren't met.
type DeploymentMetrics_DeltasMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeploymentMetrics_DeltasMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeploymentMetrics_DeltasMultiError) AllErrors() []error { return m }

// DeploymentMetrics_DeltasValidationError is the validation error returned by
// DeploymentMetrics_Deltas.Validate if the designated constraints aren't met.
type DeploymentMetrics_DeltasValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeploymentMetrics_DeltasValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeploymentMetrics_DeltasValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeploymentMetrics_DeltasValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeploymentMetrics_DeltasValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeploymentMetrics_DeltasValidationError) ErrorName() string {
	return "DeploymentMetrics_DeltasValidationError"
}

// Error satisfies the builtin error interface
func (e DeploymentMetrics_DeltasValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeploymentMetrics_Deltas.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeploymentMetrics_DeltasValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeploymentMetrics_DeltasValidationError{}

// Validate checks the field values on DeploymentMetrics_Values with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeploymentMetrics_Values) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeploymentMetrics_Values with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeploymentMetrics_ValuesMultiError, or nil if none found.
func (m *DeploymentMetrics_Values) ValidateAll() error {
	return m.validate(true)
}

func (m *DeploymentMetrics_Values) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalDeployments

	// no validation rules for SuccessfulDeployments

	// no validation rules for SuccessfulDeploymentsPerDay

	// no validation rules for FailedDeployments

	if all {
		switch v := interface{}(m.GetMedianSuccessfulDeploymentDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeploymentMetrics_ValuesValidationError{
					field:  "MedianSuccessfulDeploymentDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeploymentMetrics_ValuesValidationError{
					field:  "MedianSuccessfulDeploymentDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMedianSuccessfulDeploymentDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeploymentMetrics_ValuesValidationError{
				field:  "MedianSuccessfulDeploymentDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DeploymentFailureRate

	// no validation rules for TotalRollbacks

	// no validation rules for SuccessfulRollbacks

	if all {
		switch v := interface{}(m.GetMedianSuccessfulRollbackDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeploymentMetrics_ValuesValidationError{
					field:  "MedianSuccessfulRollbackDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeploymentMetrics_ValuesValidationError{
					field:  "MedianSuccessfulRollbackDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMedianSuccessfulRollbackDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeploymentMetrics_ValuesValidationError{
				field:  "MedianSuccessfulRollbackDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLastSuccessfulDeploymentCompletionTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeploymentMetrics_ValuesValidationError{
					field:  "LastSuccessfulDeploymentCompletionTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeploymentMetrics_ValuesValidationError{
					field:  "LastSuccessfulDeploymentCompletionTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastSuccessfulDeploymentCompletionTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeploymentMetrics_ValuesValidationError{
				field:  "LastSuccessfulDeploymentCompletionTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeploymentMetrics_ValuesMultiError(errors)
	}

	return nil
}

// DeploymentMetrics_ValuesMultiError is an error wrapping multiple validation
// errors returned by DeploymentMetrics_Values.ValidateAll() if the designated
// constraints aren't met.
type DeploymentMetrics_ValuesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeploymentMetrics_ValuesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeploymentMetrics_ValuesMultiError) AllErrors() []error { return m }

// DeploymentMetrics_ValuesValidationError is the validation error returned by
// DeploymentMetrics_Values.Validate if the designated constraints aren't met.
type DeploymentMetrics_ValuesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeploymentMetrics_ValuesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeploymentMetrics_ValuesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeploymentMetrics_ValuesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeploymentMetrics_ValuesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeploymentMetrics_ValuesValidationError) ErrorName() string {
	return "DeploymentMetrics_ValuesValidationError"
}

// Error satisfies the builtin error interface
func (e DeploymentMetrics_ValuesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeploymentMetrics_Values.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeploymentMetrics_ValuesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeploymentMetrics_ValuesValidationError{}

// Validate checks the field values on DeploymentMetrics_TimeseriesMetrics with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *DeploymentMetrics_TimeseriesMetrics) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeploymentMetrics_TimeseriesMetrics
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// DeploymentMetrics_TimeseriesMetricsMultiError, or nil if none found.
func (m *DeploymentMetrics_TimeseriesMetrics) ValidateAll() error {
	return m.validate(true)
}

func (m *DeploymentMetrics_TimeseriesMetrics) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeploymentMetrics_TimeseriesMetricsValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeploymentMetrics_TimeseriesMetricsValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeploymentMetrics_TimeseriesMetricsValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SuccessfulDeployments

	// no validation rules for FailedDeployments

	if len(errors) > 0 {
		return DeploymentMetrics_TimeseriesMetricsMultiError(errors)
	}

	return nil
}

// DeploymentMetrics_TimeseriesMetricsMultiError is an error wrapping multiple
// validation errors returned by
// DeploymentMetrics_TimeseriesMetrics.ValidateAll() if the designated
// constraints aren't met.
type DeploymentMetrics_TimeseriesMetricsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeploymentMetrics_TimeseriesMetricsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeploymentMetrics_TimeseriesMetricsMultiError) AllErrors() []error { return m }

// DeploymentMetrics_TimeseriesMetricsValidationError is the validation error
// returned by DeploymentMetrics_TimeseriesMetrics.Validate if the designated
// constraints aren't met.
type DeploymentMetrics_TimeseriesMetricsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeploymentMetrics_TimeseriesMetricsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeploymentMetrics_TimeseriesMetricsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeploymentMetrics_TimeseriesMetricsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeploymentMetrics_TimeseriesMetricsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeploymentMetrics_TimeseriesMetricsValidationError) ErrorName() string {
	return "DeploymentMetrics_TimeseriesMetricsValidationError"
}

// Error satisfies the builtin error interface
func (e DeploymentMetrics_TimeseriesMetricsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeploymentMetrics_TimeseriesMetrics.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeploymentMetrics_TimeseriesMetricsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeploymentMetrics_TimeseriesMetricsValidationError{}
