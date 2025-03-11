package webapi

import (
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/domain/hookschedule"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/web"
	"github.com/samber/lo"
)

type HookScheduleView struct {
	HookScheduleID        uuid.UUID                          `json:"hookScheduleId"`
	HookID                uuid.UUID                          `json:"hookId"`
	DisplayName           string                             `json:"displayName"`
	Description           string                             `json:"description"`
	ScheduleFrequencyUnit hookschedule.ScheduleFrequencyUnit `json:"scheduleFrequencyUnit"`
	ScheduleTimeMonth     int16                              `json:"scheduleTimeMonth"`
	ScheduleTimeDay       int16                              `json:"scheduleTimeDay"`
	ScheduleTimeHour      int16                              `json:"scheduleTimeHour"`
	ScheduleTimeMinute    int16                              `json:"scheduleTimeMinute"`
	ScheduleTimeSecond    int16                              `json:"scheduleTimeSecond"`
}

func fromHookScheduleModel(model *models.HookSchedule) HookScheduleView {
	return HookScheduleView{
		HookScheduleID:        model.HookScheduleID,
		HookID:                model.HookID,
		DisplayName:           model.DisplayName,
		Description:           model.Description,
		ScheduleFrequencyUnit: hookschedule.ScheduleFrequencyUnit(model.ScheduleFrequencyUnit),
		ScheduleTimeMonth:     model.ScheduleTimeMonth,
		ScheduleTimeDay:       model.ScheduleTimeDay,
		ScheduleTimeHour:      model.ScheduleTimeHour,
		ScheduleTimeMinute:    model.ScheduleTimeMinute,
		ScheduleTimeSecond:    model.ScheduleTimeSecond,
	}
}

func fromHookScheduleModels(hooks models.HookScheduleSlice, total int, limit int, offset int) web.ListRes[HookScheduleView] {
	items := lo.Map(hooks, func(m *models.HookSchedule, _ int) HookScheduleView {
		return fromHookScheduleModel(m)
	})
	return web.ListRes[HookScheduleView]{
		Items:  items,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}
}
