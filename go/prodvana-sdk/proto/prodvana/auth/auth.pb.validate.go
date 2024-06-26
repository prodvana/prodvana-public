// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prodvana/auth/auth.proto

package auth

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

// Validate checks the field values on AuthToken with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthToken) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthToken with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthTokenMultiError, or nil
// if none found.
func (m *AuthToken) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthToken) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return AuthTokenMultiError(errors)
	}

	return nil
}

// AuthTokenMultiError is an error wrapping multiple validation errors returned
// by AuthToken.ValidateAll() if the designated constraints aren't met.
type AuthTokenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthTokenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthTokenMultiError) AllErrors() []error { return m }

// AuthTokenValidationError is the validation error returned by
// AuthToken.Validate if the designated constraints aren't met.
type AuthTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthTokenValidationError) ErrorName() string { return "AuthTokenValidationError" }

// Error satisfies the builtin error interface
func (e AuthTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthTokenValidationError{}

// Validate checks the field values on ApiTokenInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ApiTokenInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApiTokenInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ApiTokenInfoMultiError, or
// nil if none found.
func (m *ApiTokenInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *ApiTokenInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Description

	if all {
		switch v := interface{}(m.GetExpiresTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApiTokenInfoValidationError{
					field:  "ExpiresTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApiTokenInfoValidationError{
					field:  "ExpiresTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiresTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApiTokenInfoValidationError{
				field:  "ExpiresTimestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreationTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ApiTokenInfoValidationError{
					field:  "CreationTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ApiTokenInfoValidationError{
					field:  "CreationTimestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreationTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApiTokenInfoValidationError{
				field:  "CreationTimestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ApiTokenInfoMultiError(errors)
	}

	return nil
}

// ApiTokenInfoMultiError is an error wrapping multiple validation errors
// returned by ApiTokenInfo.ValidateAll() if the designated constraints aren't met.
type ApiTokenInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApiTokenInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApiTokenInfoMultiError) AllErrors() []error { return m }

// ApiTokenInfoValidationError is the validation error returned by
// ApiTokenInfo.Validate if the designated constraints aren't met.
type ApiTokenInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApiTokenInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApiTokenInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApiTokenInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApiTokenInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApiTokenInfoValidationError) ErrorName() string { return "ApiTokenInfoValidationError" }

// Error satisfies the builtin error interface
func (e ApiTokenInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApiTokenInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApiTokenInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApiTokenInfoValidationError{}

// Validate checks the field values on AuthContext with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthContext) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthContext with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthContextMultiError, or
// nil if none found.
func (m *AuthContext) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthContext) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAuthToken()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AuthContextValidationError{
					field:  "AuthToken",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AuthContextValidationError{
					field:  "AuthToken",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuthContextValidationError{
				field:  "AuthToken",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ApiToken

	// no validation rules for Addr

	// no validation rules for SkipTls

	if len(errors) > 0 {
		return AuthContextMultiError(errors)
	}

	return nil
}

// AuthContextMultiError is an error wrapping multiple validation errors
// returned by AuthContext.ValidateAll() if the designated constraints aren't met.
type AuthContextMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthContextMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthContextMultiError) AllErrors() []error { return m }

// AuthContextValidationError is the validation error returned by
// AuthContext.Validate if the designated constraints aren't met.
type AuthContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthContextValidationError) ErrorName() string { return "AuthContextValidationError" }

// Error satisfies the builtin error interface
func (e AuthContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthContextValidationError{}

// Validate checks the field values on Session with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Session) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Session with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SessionMultiError, or nil if none found.
func (m *Session) ValidateAll() error {
	return m.validate(true)
}

func (m *Session) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetContexts()))
		i := 0
		for key := range m.GetContexts() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetContexts()[key]
			_ = val

			// no validation rules for Contexts[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, SessionValidationError{
							field:  fmt.Sprintf("Contexts[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, SessionValidationError{
							field:  fmt.Sprintf("Contexts[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return SessionValidationError{
						field:  fmt.Sprintf("Contexts[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	// no validation rules for CurrentContext

	{
		sorted_keys := make([]string, len(m.GetAdminContexts()))
		i := 0
		for key := range m.GetAdminContexts() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetAdminContexts()[key]
			_ = val

			// no validation rules for AdminContexts[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, SessionValidationError{
							field:  fmt.Sprintf("AdminContexts[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, SessionValidationError{
							field:  fmt.Sprintf("AdminContexts[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return SessionValidationError{
						field:  fmt.Sprintf("AdminContexts[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	// no validation rules for CurrentAdminContext

	if len(errors) > 0 {
		return SessionMultiError(errors)
	}

	return nil
}

// SessionMultiError is an error wrapping multiple validation errors returned
// by Session.ValidateAll() if the designated constraints aren't met.
type SessionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SessionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SessionMultiError) AllErrors() []error { return m }

// SessionValidationError is the validation error returned by Session.Validate
// if the designated constraints aren't met.
type SessionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SessionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SessionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SessionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SessionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SessionValidationError) ErrorName() string { return "SessionValidationError" }

// Error satisfies the builtin error interface
func (e SessionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSession.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SessionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SessionValidationError{}
