package webapi

//go:generate mockgen -destination=mock_webapi/mock_hook_usecase.go . HookUsecaseIF
//go:generate mockgen -destination=mock_webapi/mock_hook_schedule_usecase.go . HookScheduleUsecaseIF
