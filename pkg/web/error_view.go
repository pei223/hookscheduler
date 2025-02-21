package web

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pei223/hook-scheduler/pkg/errorcommon"
)

var errStatusMap = map[errorcommon.ErrorType]int{
	errorcommon.ErrNoData:           http.StatusNotFound,
	errorcommon.ErrParamsInvalid:    http.StatusBadRequest,
	errorcommon.ErrConflict:         http.StatusConflict,
	errorcommon.ErrOthers:           http.StatusInternalServerError,
	errorcommon.ErrNoAuthorization:  http.StatusForbidden,
	errorcommon.ErrNotAuthenticated: http.StatusUnauthorized,
}

// RFC 7807
type ErrorResponse struct {
	Type          errorcommon.ErrorType      `json:"type,omitempty"`
	Title         string                     `json:"title"`
	Detail        string                     `json:"detail,omitempty"`
	Instance      string                     `json:"instance,omitempty"`
	InvalidParams []errorcommon.InvalidParam `json:"invalidParams,omitempty"` // フィールドごとのエラー
}

func ErrorResFrom(c *gin.Context, err error) (int, ErrorResponse) {
	if errors.Is(err, sql.ErrNoRows) {
		return http.StatusNotFound, ErrorResponse{
			Title:    "no data",
			Type:     errorcommon.ErrNoData,
			Instance: c.Request.URL.Path,
		}
	}
	var commonErr *errorcommon.CommonError
	if ok := errors.As(err, &commonErr); ok {
		status := errStatusMap[commonErr.Type]
		if status == 0 {
			status = http.StatusInternalServerError
		}
		return status, ErrorResponse{
			Instance:      c.Request.URL.Path,
			Title:         commonErr.Title,
			Type:          commonErr.Type,
			InvalidParams: commonErr.InvalidParams,
		}
	}
	return http.StatusInternalServerError, ErrorResponse{
		Instance: c.Request.URL.Path,
		Title:    err.Error(),
		Type:     errorcommon.ErrOthers,
	}
}
