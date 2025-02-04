package webapi

import (
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/types"
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
