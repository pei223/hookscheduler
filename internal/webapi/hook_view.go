package webapi

import (
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/types"
	"github.com/pei223/hook-scheduler/pkg/web"
	"github.com/samber/lo"
)

type HookView struct {
	HookID      uuid.UUID   `json:"hookId"`
	DisplayName string      `json:"displayName"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	Method      string      `json:"method"`
	Body        types.JSONB `json:"body"`
	Headers     types.JSONB `json:"header"`
}

func fromModel(model *models.Hook) HookView {
	return HookView{
		HookID:      model.HookID,
		DisplayName: model.DisplayName,
		Description: model.Description,
		URL:         model.URL,
		Method:      model.Method,
		Body:        model.Body,
		Headers:     model.Headers,
	}
}

func fromModels(hooks models.HookSlice, total int, limit int, offset int) web.ListRes[HookView] {
	items := lo.Map(hooks, func(m *models.Hook, _ int) HookView {
		return fromModel(m)
	})
	return web.ListRes[HookView]{
		Items:  items,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}
}
