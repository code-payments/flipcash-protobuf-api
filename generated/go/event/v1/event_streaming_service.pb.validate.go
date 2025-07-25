// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: event/v1/event_streaming_service.proto

package eventpb

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

// Validate checks the field values on StreamEventsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StreamEventsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamEventsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StreamEventsRequestMultiError, or nil if none found.
func (m *StreamEventsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamEventsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofTypePresent := false
	switch v := m.Type.(type) {
	case *StreamEventsRequest_Params_:
		if v == nil {
			err := StreamEventsRequestValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if all {
			switch v := interface{}(m.GetParams()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamEventsRequestValidationError{
						field:  "Params",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamEventsRequestValidationError{
						field:  "Params",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetParams()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamEventsRequestValidationError{
					field:  "Params",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *StreamEventsRequest_Pong:
		if v == nil {
			err := StreamEventsRequestValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if all {
			switch v := interface{}(m.GetPong()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamEventsRequestValidationError{
						field:  "Pong",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamEventsRequestValidationError{
						field:  "Pong",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPong()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamEventsRequestValidationError{
					field:  "Pong",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofTypePresent {
		err := StreamEventsRequestValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StreamEventsRequestMultiError(errors)
	}

	return nil
}

// StreamEventsRequestMultiError is an error wrapping multiple validation
// errors returned by StreamEventsRequest.ValidateAll() if the designated
// constraints aren't met.
type StreamEventsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamEventsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamEventsRequestMultiError) AllErrors() []error { return m }

// StreamEventsRequestValidationError is the validation error returned by
// StreamEventsRequest.Validate if the designated constraints aren't met.
type StreamEventsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamEventsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamEventsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamEventsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamEventsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamEventsRequestValidationError) ErrorName() string {
	return "StreamEventsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StreamEventsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamEventsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamEventsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamEventsRequestValidationError{}

// Validate checks the field values on StreamEventsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StreamEventsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamEventsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StreamEventsResponseMultiError, or nil if none found.
func (m *StreamEventsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamEventsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofTypePresent := false
	switch v := m.Type.(type) {
	case *StreamEventsResponse_Ping:
		if v == nil {
			err := StreamEventsResponseValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if all {
			switch v := interface{}(m.GetPing()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamEventsResponseValidationError{
						field:  "Ping",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamEventsResponseValidationError{
						field:  "Ping",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPing()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamEventsResponseValidationError{
					field:  "Ping",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *StreamEventsResponse_Error:
		if v == nil {
			err := StreamEventsResponseValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if all {
			switch v := interface{}(m.GetError()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamEventsResponseValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamEventsResponseValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamEventsResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *StreamEventsResponse_Events:
		if v == nil {
			err := StreamEventsResponseValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTypePresent = true

		if all {
			switch v := interface{}(m.GetEvents()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StreamEventsResponseValidationError{
						field:  "Events",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StreamEventsResponseValidationError{
						field:  "Events",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEvents()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamEventsResponseValidationError{
					field:  "Events",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofTypePresent {
		err := StreamEventsResponseValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StreamEventsResponseMultiError(errors)
	}

	return nil
}

// StreamEventsResponseMultiError is an error wrapping multiple validation
// errors returned by StreamEventsResponse.ValidateAll() if the designated
// constraints aren't met.
type StreamEventsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamEventsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamEventsResponseMultiError) AllErrors() []error { return m }

// StreamEventsResponseValidationError is the validation error returned by
// StreamEventsResponse.Validate if the designated constraints aren't met.
type StreamEventsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamEventsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamEventsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamEventsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamEventsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamEventsResponseValidationError) ErrorName() string {
	return "StreamEventsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StreamEventsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamEventsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamEventsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamEventsResponseValidationError{}

// Validate checks the field values on ForwardEventsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ForwardEventsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForwardEventsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ForwardEventsRequestMultiError, or nil if none found.
func (m *ForwardEventsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ForwardEventsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserEvents() == nil {
		err := ForwardEventsRequestValidationError{
			field:  "UserEvents",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetUserEvents()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ForwardEventsRequestValidationError{
					field:  "UserEvents",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForwardEventsRequestValidationError{
					field:  "UserEvents",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserEvents()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForwardEventsRequestValidationError{
				field:  "UserEvents",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ForwardEventsRequestMultiError(errors)
	}

	return nil
}

// ForwardEventsRequestMultiError is an error wrapping multiple validation
// errors returned by ForwardEventsRequest.ValidateAll() if the designated
// constraints aren't met.
type ForwardEventsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForwardEventsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForwardEventsRequestMultiError) AllErrors() []error { return m }

// ForwardEventsRequestValidationError is the validation error returned by
// ForwardEventsRequest.Validate if the designated constraints aren't met.
type ForwardEventsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForwardEventsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForwardEventsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForwardEventsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForwardEventsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForwardEventsRequestValidationError) ErrorName() string {
	return "ForwardEventsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ForwardEventsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForwardEventsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForwardEventsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForwardEventsRequestValidationError{}

// Validate checks the field values on ForwardEventsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ForwardEventsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForwardEventsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ForwardEventsResponseMultiError, or nil if none found.
func (m *ForwardEventsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ForwardEventsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return ForwardEventsResponseMultiError(errors)
	}

	return nil
}

// ForwardEventsResponseMultiError is an error wrapping multiple validation
// errors returned by ForwardEventsResponse.ValidateAll() if the designated
// constraints aren't met.
type ForwardEventsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForwardEventsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForwardEventsResponseMultiError) AllErrors() []error { return m }

// ForwardEventsResponseValidationError is the validation error returned by
// ForwardEventsResponse.Validate if the designated constraints aren't met.
type ForwardEventsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForwardEventsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForwardEventsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForwardEventsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForwardEventsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForwardEventsResponseValidationError) ErrorName() string {
	return "ForwardEventsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ForwardEventsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForwardEventsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForwardEventsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForwardEventsResponseValidationError{}

// Validate checks the field values on StreamEventsRequest_Params with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StreamEventsRequest_Params) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamEventsRequest_Params with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StreamEventsRequest_ParamsMultiError, or nil if none found.
func (m *StreamEventsRequest_Params) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamEventsRequest_Params) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAuth() == nil {
		err := StreamEventsRequest_ParamsValidationError{
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
				errors = append(errors, StreamEventsRequest_ParamsValidationError{
					field:  "Auth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StreamEventsRequest_ParamsValidationError{
					field:  "Auth",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuth()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamEventsRequest_ParamsValidationError{
				field:  "Auth",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetTs() == nil {
		err := StreamEventsRequest_ParamsValidationError{
			field:  "Ts",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StreamEventsRequest_ParamsMultiError(errors)
	}

	return nil
}

// StreamEventsRequest_ParamsMultiError is an error wrapping multiple
// validation errors returned by StreamEventsRequest_Params.ValidateAll() if
// the designated constraints aren't met.
type StreamEventsRequest_ParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamEventsRequest_ParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamEventsRequest_ParamsMultiError) AllErrors() []error { return m }

// StreamEventsRequest_ParamsValidationError is the validation error returned
// by StreamEventsRequest_Params.Validate if the designated constraints aren't met.
type StreamEventsRequest_ParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamEventsRequest_ParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamEventsRequest_ParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamEventsRequest_ParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamEventsRequest_ParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamEventsRequest_ParamsValidationError) ErrorName() string {
	return "StreamEventsRequest_ParamsValidationError"
}

// Error satisfies the builtin error interface
func (e StreamEventsRequest_ParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamEventsRequest_Params.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamEventsRequest_ParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamEventsRequest_ParamsValidationError{}

// Validate checks the field values on StreamEventsResponse_StreamError with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *StreamEventsResponse_StreamError) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamEventsResponse_StreamError with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// StreamEventsResponse_StreamErrorMultiError, or nil if none found.
func (m *StreamEventsResponse_StreamError) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamEventsResponse_StreamError) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	if len(errors) > 0 {
		return StreamEventsResponse_StreamErrorMultiError(errors)
	}

	return nil
}

// StreamEventsResponse_StreamErrorMultiError is an error wrapping multiple
// validation errors returned by
// StreamEventsResponse_StreamError.ValidateAll() if the designated
// constraints aren't met.
type StreamEventsResponse_StreamErrorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamEventsResponse_StreamErrorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamEventsResponse_StreamErrorMultiError) AllErrors() []error { return m }

// StreamEventsResponse_StreamErrorValidationError is the validation error
// returned by StreamEventsResponse_StreamError.Validate if the designated
// constraints aren't met.
type StreamEventsResponse_StreamErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamEventsResponse_StreamErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamEventsResponse_StreamErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamEventsResponse_StreamErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamEventsResponse_StreamErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamEventsResponse_StreamErrorValidationError) ErrorName() string {
	return "StreamEventsResponse_StreamErrorValidationError"
}

// Error satisfies the builtin error interface
func (e StreamEventsResponse_StreamErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamEventsResponse_StreamError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamEventsResponse_StreamErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamEventsResponse_StreamErrorValidationError{}
