package hook

import (
	"github.com/pei223/hook-scheduler/pkg/errorcommon"
	"github.com/pei223/hook-scheduler/pkg/types"
	"github.com/pei223/hook-scheduler/pkg/web"
	"github.com/samber/lo"
)

type HookCreateParams struct {
	DisplayName string      `json:"displayName" validate:"required,max=20,min=1"`
	URL         string      `json:"url" validate:"required,url"`
	Method      string      `json:"method" validate:"required,max=20,min=1"`
	Body        types.JSONB `json:"body" validate:"required"`
	Headers     types.JSONB `json:"headers" validate:"required"`
}

func (p *HookCreateParams) Validate() []errorcommon.InvalidParam {
	invalidParams := web.SchemaValidate(p)
	if invalidParams != nil {
		return invalidParams
	}
	if !lo.Contains(
		[]string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		p.Method,
	) {
		invalidParam := errorcommon.InvalidParam{
			Name:   "method",
			Reason: "method must be one of GET, POST, PUT, PATCH, DELETE",
		}
		if invalidParams == nil {
			return []errorcommon.InvalidParam{
				invalidParam,
			}
		}
		invalidParams = append(invalidParams, invalidParam)
	}
	return invalidParams
}
