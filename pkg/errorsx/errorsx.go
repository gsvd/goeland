package errorsx

import (
	"encoding/json"
	"errors"
)

type ErrorCode string

const (
	ErrCodePasswordRequired     ErrorCode = "ERR_PASSWORD_REQUIRED"
	ErrCodeAccountExists        ErrorCode = "ERR_ACCOUNT_EXISTS"
	ErrCodeAuthenticationFailed ErrorCode = "ERR_AUTHENTICATION_FAILED"

	ErrCodeInvalidInput         ErrorCode = "ERR_INVALID_INPUT"
	ErrCodeEmptyAddress         ErrorCode = "ERR_EMPTY_ADDRESS"
	ErrCodeInvalidAddressFormat ErrorCode = "ERR_INVALID_ADDRESS_FORMAT"
	ErrCodeUnknown              ErrorCode = "ERR_UNKNOWN"
)

type FieldError struct {
	Field string    `json:"field"`
	Code  ErrorCode `json:"code"`
}

type ValidationErrors struct {
	errors []FieldError
}

func (v *ValidationErrors) Add(field string, code ErrorCode) {
	v.errors = append(v.errors, FieldError{
		Field: field,
		Code:  code,
	})
}

func (v *ValidationErrors) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *ValidationErrors) Error() string {
	return string(ErrCodeInvalidInput)
}

func (v *ValidationErrors) AsAPIError() *APIError {
	return &APIError{
		Code:   string(ErrCodeInvalidInput),
		Errors: v.errors,
	}
}

type APIError struct {
	Code   string       `json:"code"`
	Errors []FieldError `json:"errors,omitempty"`
}

func (e *APIError) Error() string {
	return e.Code
}

func NewAPIError(code ErrorCode) *APIError {
	return &APIError{
		Code: string(code),
	}
}

func (e *APIError) Marshal() error {
	b, _ := json.Marshal(e)
	return errors.New(string(b))
}
