// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: thirdparty/v1/third_party_service.proto

package thirdpartypb

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

// Validate checks the field values on GetJwtRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetJwtRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetJwtRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetJwtRequestMultiError, or
// nil if none found.
func (m *GetJwtRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetJwtRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetApiKey() == nil {
		err := GetJwtRequestValidationError{
			field:  "ApiKey",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetApiKey()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetJwtRequestValidationError{
					field:  "ApiKey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetJwtRequestValidationError{
					field:  "ApiKey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApiKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetJwtRequestValidationError{
				field:  "ApiKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetMethod()); l < 3 || l > 4 {
		err := GetJwtRequestValidationError{
			field:  "Method",
			reason: "value length must be between 3 and 4 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetHost()); l < 1 || l > 1024 {
		err := GetJwtRequestValidationError{
			field:  "Host",
			reason: "value length must be between 1 and 1024 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPath()); l < 1 || l > 1024 {
		err := GetJwtRequestValidationError{
			field:  "Path",
			reason: "value length must be between 1 and 1024 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAuth() == nil {
		err := GetJwtRequestValidationError{
			field:  "Auth",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAuth()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetJwtRequestValidationError{
					field:  "Auth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetJwtRequestValidationError{
					field:  "Auth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuth()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetJwtRequestValidationError{
				field:  "Auth",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetJwtRequestMultiError(errors)
	}

	return nil
}

// GetJwtRequestMultiError is an error wrapping multiple validation errors
// returned by GetJwtRequest.ValidateAll() if the designated constraints
// aren't met.
type GetJwtRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetJwtRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetJwtRequestMultiError) AllErrors() []error { return m }

// GetJwtRequestValidationError is the validation error returned by
// GetJwtRequest.Validate if the designated constraints aren't met.
type GetJwtRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetJwtRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetJwtRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetJwtRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetJwtRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetJwtRequestValidationError) ErrorName() string { return "GetJwtRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetJwtRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetJwtRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetJwtRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetJwtRequestValidationError{}

// Validate checks the field values on GetJwtResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetJwtResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetJwtResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetJwtResponseMultiError,
// or nil if none found.
func (m *GetJwtResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetJwtResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if all {
		switch v := interface{}(m.GetJwt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetJwtResponseValidationError{
					field:  "Jwt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetJwtResponseValidationError{
					field:  "Jwt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetJwt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetJwtResponseValidationError{
				field:  "Jwt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetJwtResponseMultiError(errors)
	}

	return nil
}

// GetJwtResponseMultiError is an error wrapping multiple validation errors
// returned by GetJwtResponse.ValidateAll() if the designated constraints
// aren't met.
type GetJwtResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetJwtResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetJwtResponseMultiError) AllErrors() []error { return m }

// GetJwtResponseValidationError is the validation error returned by
// GetJwtResponse.Validate if the designated constraints aren't met.
type GetJwtResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetJwtResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetJwtResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetJwtResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetJwtResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetJwtResponseValidationError) ErrorName() string { return "GetJwtResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetJwtResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetJwtResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetJwtResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetJwtResponseValidationError{}
