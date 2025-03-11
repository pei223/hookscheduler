package hookschedule_test

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/domain/hookschedule"
	"github.com/pei223/hook-scheduler/internal/test_common"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestValidation(t *testing.T) {
	baseParams := hookschedule.HookScheduleCreateParams{
		HookID:                uuid.New(),
		DisplayName:           "test",
		Description:           "test schedule",
		ScheduleFrequencyUnit: hookschedule.ScheduleFrequencyEveryMinute,
		ScheduleTimeMonth:     1,
		ScheduleTimeDay:       1,
		ScheduleTimeHour:      0,
		ScheduleTimeMinute:    0,
		ScheduleTimeSecond:    0,
	}

	t.Run("success", func(t *testing.T) {
		invalidParams := baseParams.Validate()
		require.Nil(t, invalidParams)
	})

	badRequestCases := []struct {
		caseName     string
		invalidCount int
		genBody      func(base hookschedule.HookScheduleCreateParams) hookschedule.HookScheduleCreateParams
	}{
		{
			caseName:     "too long displayName",
			invalidCount: 1,
			genBody: func(base hookschedule.HookScheduleCreateParams) hookschedule.HookScheduleCreateParams {
				base.DisplayName = "1234567890123456789012345678901234567890"
				return base
			},
		},
		{
			caseName:     "empty displayName",
			invalidCount: 1,
			genBody: func(base hookschedule.HookScheduleCreateParams) hookschedule.HookScheduleCreateParams {
				base.DisplayName = ""
				return base
			},
		},
		{
			caseName:     "invalid scheduleTimeMonth",
			invalidCount: 1,
			genBody: func(base hookschedule.HookScheduleCreateParams) hookschedule.HookScheduleCreateParams {
				base.ScheduleTimeMonth = 13
				return base
			},
		},
		{
			caseName:     "invalid scheduleTimeDay",
			invalidCount: 1,
			genBody: func(base hookschedule.HookScheduleCreateParams) hookschedule.HookScheduleCreateParams {
				base.ScheduleTimeDay = 32
				return base
			},
		},
		{
			caseName:     "invalid scheduleTimeHour",
			invalidCount: 1,
			genBody: func(base hookschedule.HookScheduleCreateParams) hookschedule.HookScheduleCreateParams {
				base.ScheduleTimeHour = 24
				return base
			},
		},
		{
			caseName:     "invalid scheduleTimeMinute",
			invalidCount: 1,
			genBody: func(base hookschedule.HookScheduleCreateParams) hookschedule.HookScheduleCreateParams {
				base.ScheduleTimeMinute = 60
				return base
			},
		},
		{
			caseName:     "invalid scheduleTimeSecond",
			invalidCount: 1,
			genBody: func(base hookschedule.HookScheduleCreateParams) hookschedule.HookScheduleCreateParams {
				base.ScheduleTimeSecond = 60
				return base
			},
		},
		{
			caseName:     "invalid scheduleFrequencyUnit",
			invalidCount: 1,
			genBody: func(base hookschedule.HookScheduleCreateParams) hookschedule.HookScheduleCreateParams {
				base.ScheduleFrequencyUnit = "invalid"
				return base
			},
		},
	}

	for _, c := range badRequestCases {
		t.Run(c.caseName, func(t *testing.T) {
			baseBody := lo.Must(test_common.DeepCopy(baseParams))
			body := c.genBody(*baseBody)

			invalidParams := body.Validate()
			require.Len(t, invalidParams, c.invalidCount)
			snaps.MatchJSON(t, invalidParams)
		})
	}
}
