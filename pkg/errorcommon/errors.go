package errorcommon

import "errors"

type ErrorType string

var (
	ErrParamsInvalid    ErrorType = "err_params"
	ErrNoData           ErrorType = "err_no_data"
	ErrConflict         ErrorType = "err_conflict"
	ErrNoAuthorization  ErrorType = "err_no_authorization"
	ErrNotAuthenticated ErrorType = "err_not_authenticated"
	ErrOthers           ErrorType = "err_others"
)

type InvalidParam struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type CommonError struct {
	Err           error
	Title         string
	Type          ErrorType
	InvalidParams []InvalidParam `json:"omitempty"`
}

func NewCommonError(err error, title string, errType ErrorType, invalidParams []InvalidParam) *CommonError {
	return &CommonError{
		Err:           err,
		Title:         title,
		Type:          errType,
		InvalidParams: invalidParams,
	}
}

func (e *CommonError) Error() string {
	return e.Err.Error()
}

func NewParseError(err error) *CommonError {
	return &CommonError{
		Err:   err,
		Title: "failed to parse",
		Type:  ErrParamsInvalid,
	}
}

func NewInvalidParamsError(invalidParams []InvalidParam) *CommonError {
	return &CommonError{
		Err:           errors.New("invalid params"),
		Title:         "invalid params",
		Type:          ErrParamsInvalid,
		InvalidParams: invalidParams,
	}
}
