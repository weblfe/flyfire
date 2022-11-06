// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/account.proto

package v1

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

// Validate checks the field values on GetUserInfoParams with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUserInfoParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserInfoParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserInfoParamsMultiError, or nil if none found.
func (m *GetUserInfoParams) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserInfoParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetId()); l < 1 || l > 128 {
		err := GetUserInfoParamsValidationError{
			field:  "Id",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetUserInfoParamsMultiError(errors)
	}

	return nil
}

// GetUserInfoParamsMultiError is an error wrapping multiple validation errors
// returned by GetUserInfoParams.ValidateAll() if the designated constraints
// aren't met.
type GetUserInfoParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserInfoParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserInfoParamsMultiError) AllErrors() []error { return m }

// GetUserInfoParamsValidationError is the validation error returned by
// GetUserInfoParams.Validate if the designated constraints aren't met.
type GetUserInfoParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserInfoParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserInfoParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserInfoParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserInfoParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserInfoParamsValidationError) ErrorName() string {
	return "GetUserInfoParamsValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserInfoParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserInfoParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserInfoParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserInfoParamsValidationError{}

// Validate checks the field values on GetUserInfoReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUserInfoReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserInfoReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserInfoReplyMultiError, or nil if none found.
func (m *GetUserInfoReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserInfoReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Username

	// no validation rules for RoleType

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserInfoReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserInfoReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserInfoReplyValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserInfoReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserInfoReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserInfoReplyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetUserInfoReplyMultiError(errors)
	}

	return nil
}

// GetUserInfoReplyMultiError is an error wrapping multiple validation errors
// returned by GetUserInfoReply.ValidateAll() if the designated constraints
// aren't met.
type GetUserInfoReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserInfoReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserInfoReplyMultiError) AllErrors() []error { return m }

// GetUserInfoReplyValidationError is the validation error returned by
// GetUserInfoReply.Validate if the designated constraints aren't met.
type GetUserInfoReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserInfoReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserInfoReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserInfoReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserInfoReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserInfoReplyValidationError) ErrorName() string { return "GetUserInfoReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUserInfoReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserInfoReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserInfoReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserInfoReplyValidationError{}