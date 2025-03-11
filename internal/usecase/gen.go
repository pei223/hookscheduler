package usecase

//go:generate mockgen -destination=mock_usecase/mock_hook_service.go . HookServiceIF
//go:generate mockgen -destination=mock_usecase/mock_hook_exec_service.go . HookExecServiceIF
//go:generate mockgen -destination=mock_usecase/mock_hook_schedule_service.go . HookScheduleServiceIF
