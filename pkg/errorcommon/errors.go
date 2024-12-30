package errorcommon

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
	InvalidParams *[]InvalidParam
}

func (e *CommonError) Error() string {
	return e.Err.Error()
}
