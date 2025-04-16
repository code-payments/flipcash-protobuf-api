// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: activity/v1/model.proto

package activitypb

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

// Validate checks the field values on NotificationId with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NotificationId) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotificationId with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NotificationIdMultiError,
// or nil if none found.
func (m *NotificationId) ValidateAll() error {
	return m.validate(true)
}

func (m *NotificationId) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetValue()) != 32 {
		err := NotificationIdValidationError{
			field:  "Value",
			reason: "value length must be 32 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return NotificationIdMultiError(errors)
	}

	return nil
}

// NotificationIdMultiError is an error wrapping multiple validation errors
// returned by NotificationId.ValidateAll() if the designated constraints
// aren't met.
type NotificationIdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotificationIdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotificationIdMultiError) AllErrors() []error { return m }

// NotificationIdValidationError is the validation error returned by
// NotificationId.Validate if the designated constraints aren't met.
type NotificationIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationIdValidationError) ErrorName() string { return "NotificationIdValidationError" }

// Error satisfies the builtin error interface
func (e NotificationIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotificationId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationIdValidationError{}

// Validate checks the field values on Notification with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Notification) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Notification with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NotificationMultiError, or
// nil if none found.
func (m *Notification) ValidateAll() error {
	return m.validate(true)
}

func (m *Notification) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() == nil {
		err := NotificationValidationError{
			field:  "Id",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NotificationValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NotificationValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NotificationValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetLocalizedText()); l < 1 || l > 256 {
		err := NotificationValidationError{
			field:  "LocalizedText",
			reason: "value length must be between 1 and 256 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAmount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NotificationValidationError{
					field:  "Amount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NotificationValidationError{
					field:  "Amount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAmount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NotificationValidationError{
				field:  "Amount",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetTs() == nil {
		err := NotificationValidationError{
			field:  "Ts",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	switch v := m.AdditionalMetadata.(type) {
	case *Notification_WelcomeBonus:
		if v == nil {
			err := NotificationValidationError{
				field:  "AdditionalMetadata",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetWelcomeBonus()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NotificationValidationError{
						field:  "WelcomeBonus",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NotificationValidationError{
						field:  "WelcomeBonus",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetWelcomeBonus()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NotificationValidationError{
					field:  "WelcomeBonus",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Notification_GaveUsdc:
		if v == nil {
			err := NotificationValidationError{
				field:  "AdditionalMetadata",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetGaveUsdc()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NotificationValidationError{
						field:  "GaveUsdc",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NotificationValidationError{
						field:  "GaveUsdc",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGaveUsdc()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NotificationValidationError{
					field:  "GaveUsdc",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Notification_ReceivedUsdc:
		if v == nil {
			err := NotificationValidationError{
				field:  "AdditionalMetadata",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetReceivedUsdc()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NotificationValidationError{
						field:  "ReceivedUsdc",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NotificationValidationError{
						field:  "ReceivedUsdc",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetReceivedUsdc()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NotificationValidationError{
					field:  "ReceivedUsdc",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return NotificationMultiError(errors)
	}

	return nil
}

// NotificationMultiError is an error wrapping multiple validation errors
// returned by Notification.ValidateAll() if the designated constraints aren't met.
type NotificationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotificationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotificationMultiError) AllErrors() []error { return m }

// NotificationValidationError is the validation error returned by
// Notification.Validate if the designated constraints aren't met.
type NotificationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationValidationError) ErrorName() string { return "NotificationValidationError" }

// Error satisfies the builtin error interface
func (e NotificationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotification.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationValidationError{}

// Validate checks the field values on WelcomeBonusNotificationMetadata with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *WelcomeBonusNotificationMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WelcomeBonusNotificationMetadata with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// WelcomeBonusNotificationMetadataMultiError, or nil if none found.
func (m *WelcomeBonusNotificationMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *WelcomeBonusNotificationMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return WelcomeBonusNotificationMetadataMultiError(errors)
	}

	return nil
}

// WelcomeBonusNotificationMetadataMultiError is an error wrapping multiple
// validation errors returned by
// WelcomeBonusNotificationMetadata.ValidateAll() if the designated
// constraints aren't met.
type WelcomeBonusNotificationMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WelcomeBonusNotificationMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WelcomeBonusNotificationMetadataMultiError) AllErrors() []error { return m }

// WelcomeBonusNotificationMetadataValidationError is the validation error
// returned by WelcomeBonusNotificationMetadata.Validate if the designated
// constraints aren't met.
type WelcomeBonusNotificationMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WelcomeBonusNotificationMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WelcomeBonusNotificationMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WelcomeBonusNotificationMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WelcomeBonusNotificationMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WelcomeBonusNotificationMetadataValidationError) ErrorName() string {
	return "WelcomeBonusNotificationMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e WelcomeBonusNotificationMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWelcomeBonusNotificationMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WelcomeBonusNotificationMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WelcomeBonusNotificationMetadataValidationError{}

// Validate checks the field values on GaveUsdcNotificationMetadata with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GaveUsdcNotificationMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GaveUsdcNotificationMetadata with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GaveUsdcNotificationMetadataMultiError, or nil if none found.
func (m *GaveUsdcNotificationMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *GaveUsdcNotificationMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GaveUsdcNotificationMetadataMultiError(errors)
	}

	return nil
}

// GaveUsdcNotificationMetadataMultiError is an error wrapping multiple
// validation errors returned by GaveUsdcNotificationMetadata.ValidateAll() if
// the designated constraints aren't met.
type GaveUsdcNotificationMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GaveUsdcNotificationMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GaveUsdcNotificationMetadataMultiError) AllErrors() []error { return m }

// GaveUsdcNotificationMetadataValidationError is the validation error returned
// by GaveUsdcNotificationMetadata.Validate if the designated constraints
// aren't met.
type GaveUsdcNotificationMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GaveUsdcNotificationMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GaveUsdcNotificationMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GaveUsdcNotificationMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GaveUsdcNotificationMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GaveUsdcNotificationMetadataValidationError) ErrorName() string {
	return "GaveUsdcNotificationMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e GaveUsdcNotificationMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGaveUsdcNotificationMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GaveUsdcNotificationMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GaveUsdcNotificationMetadataValidationError{}

// Validate checks the field values on ReceivedUsdcNotificationMetadata with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *ReceivedUsdcNotificationMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceivedUsdcNotificationMetadata with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ReceivedUsdcNotificationMetadataMultiError, or nil if none found.
func (m *ReceivedUsdcNotificationMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceivedUsdcNotificationMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ReceivedUsdcNotificationMetadataMultiError(errors)
	}

	return nil
}

// ReceivedUsdcNotificationMetadataMultiError is an error wrapping multiple
// validation errors returned by
// ReceivedUsdcNotificationMetadata.ValidateAll() if the designated
// constraints aren't met.
type ReceivedUsdcNotificationMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceivedUsdcNotificationMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceivedUsdcNotificationMetadataMultiError) AllErrors() []error { return m }

// ReceivedUsdcNotificationMetadataValidationError is the validation error
// returned by ReceivedUsdcNotificationMetadata.Validate if the designated
// constraints aren't met.
type ReceivedUsdcNotificationMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceivedUsdcNotificationMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceivedUsdcNotificationMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceivedUsdcNotificationMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceivedUsdcNotificationMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceivedUsdcNotificationMetadataValidationError) ErrorName() string {
	return "ReceivedUsdcNotificationMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e ReceivedUsdcNotificationMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceivedUsdcNotificationMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceivedUsdcNotificationMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceivedUsdcNotificationMetadataValidationError{}
