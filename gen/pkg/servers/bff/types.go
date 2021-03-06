// Code generated by sysl DO NOT EDIT.
package bff

import (
	"time"

	"github.com/anz-bank/sysl-go/validator"

	"github.com/rickb777/date"
)

// Reference imports to suppress unused errors
var _ = time.Parse

// Reference imports to suppress unused errors
var _ = date.Parse

// GenericError ...
type GenericError struct {
	Status GenericErrorStatus `json:"status"`
}

// GenericErrorStatus ...
type GenericErrorStatus struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// GetAccountResponse ...
type GetAccountResponse struct {
	Amount string `json:"amount"`
}

// GetV1AccountsListRequest ...
type GetV1AccountsListRequest struct {
}

// *GenericError validator
func (s *GenericError) Validate() error {
	return validator.Validate(s)
}

// *GenericErrorStatus validator
func (s *GenericErrorStatus) Validate() error {
	return validator.Validate(s)
}

// *GetAccountResponse validator
func (s *GetAccountResponse) Validate() error {
	return validator.Validate(s)
}
