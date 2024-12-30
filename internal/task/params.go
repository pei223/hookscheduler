package task

import (
	"github.com/pei223/hook-scheduler/pkg/errorcommon"
	"github.com/pei223/hook-scheduler/pkg/web"
)

type TaskCreateParams struct {
	Name string `json:"name" validate:"required,max=20,min=1"`
}

func (p *TaskCreateParams) Validate() *[]errorcommon.InvalidParam {
	invalidParams := web.SchemaValidate(p)
	if invalidParams != nil {
		return invalidParams
	}
	// ここでカスタムvalidateなど
	return nil
}
