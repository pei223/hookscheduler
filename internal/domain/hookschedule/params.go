package hookschedule

import (
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/pkg/errorcommon"
	"github.com/pei223/hook-scheduler/pkg/web"
)

type ScheduleFrequencyUnit string

const (
	ScheduleFrequencyEveryMinute ScheduleFrequencyUnit = "every_minute"
	ScheduleFrequencyEveryHour   ScheduleFrequencyUnit = "every_hour"
	ScheduleFrequencyEveryDay    ScheduleFrequencyUnit = "every_day"
	ScheduleFrequencyEveryMonth  ScheduleFrequencyUnit = "every_month"
	ScheduleFrequencyEveryYear   ScheduleFrequencyUnit = "every_year"
)

type HookScheduleCreateParams struct {
	HookID                uuid.UUID             `json:"hookId" validate:"required"`
	DisplayName           string                `json:"displayName" validate:"required,max=20,min=1"`
	Description           string                `json:"description" validate:"max=100"`
	ScheduleFrequencyUnit ScheduleFrequencyUnit `json:"scheduleFrequencyUnit" validate:"required,oneof=every_minute every_hour every_day every_month every_year"`
	ScheduleTimeMonth     int                   `json:"scheduleTimeMonth" validate:"max=12,min=1"`
	ScheduleTimeDay       int                   `json:"scheduleTimeDay" validate:"max=31,min=1"`
	ScheduleTimeHour      int                   `json:"scheduleTimeHour" validate:"max=23,min=0"`
	ScheduleTimeMinute    int                   `json:"scheduleTimeMinute" validate:"max=59,min=0"`
	ScheduleTimeSecond    int                   `json:"scheduleTimeSecond" validate:"max=59,min=0"`
}

func (p *HookScheduleCreateParams) Validate() []errorcommon.InvalidParam {
	invalidParams := web.SchemaValidate(p)
	if invalidParams != nil {
		return invalidParams
	}
	return invalidParams
}
