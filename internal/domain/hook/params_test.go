package hook_test

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/pei223/hook-scheduler/internal/domain/hook"
	"github.com/pei223/hook-scheduler/internal/test_common"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestValidation(t *testing.T) {
	dummyHeaders := map[string]any{
		"createtestkey": "createvalue",
		"testlist": []string{
			"test1", "test2",
		},
	}
	dummyBody := map[string]any{
		"createtestkey": "createvalue",
		"testlist": []string{
			"test1", "test2",
		},
	}
	baseParams := hook.HookCreateParams{
		DisplayName: "test",
		URL:         "http://test.test",
		Method:      "GET",
		Body:        dummyBody,
		Headers:     dummyHeaders,
	}

	t.Run("success", func(t *testing.T) {
		invalidParams := baseParams.Validate()
		require.Nil(t, invalidParams)
	})

	badRequestCases := []struct {
		caseName     string
		invalidCount int
		genBody      func(base hook.HookCreateParams) hook.HookCreateParams
	}{
		{
			caseName:     "too long displayName",
			invalidCount: 1,
			genBody: func(base hook.HookCreateParams) hook.HookCreateParams {
				base.DisplayName = "1234567890123456789012345678901234567890"
				return base
			},
		},
		{
			caseName:     "empty displayName",
			invalidCount: 1,
			genBody: func(base hook.HookCreateParams) hook.HookCreateParams {
				base.DisplayName = ""
				return base
			},
		},
		{
			caseName:     "invalid url",
			invalidCount: 1,
			genBody: func(base hook.HookCreateParams) hook.HookCreateParams {
				base.URL = "invalidurl"
				return base
			},
		},
		{
			caseName:     "empty url",
			invalidCount: 1,
			genBody: func(base hook.HookCreateParams) hook.HookCreateParams {
				base.URL = ""
				return base
			},
		},
		{
			caseName:     "invalid method",
			invalidCount: 1,
			genBody: func(base hook.HookCreateParams) hook.HookCreateParams {
				base.Method = "INVALID"
				return base
			},
		},
		{
			caseName:     "empty method",
			invalidCount: 1,
			genBody: func(base hook.HookCreateParams) hook.HookCreateParams {
				base.Method = ""
				return base
			},
		},
		{
			caseName:     "empty method and displayName",
			invalidCount: 2,
			genBody: func(base hook.HookCreateParams) hook.HookCreateParams {
				base.Method = ""
				base.DisplayName = ""
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
