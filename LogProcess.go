package monitor_log

import "github.com/jihuago/monitor_log/handler"

type log_handle struct {
	context *handler.LogContext
	strategy handler.LogHandlerStrategy
}

func NewLogHandle(logCtx *handler.LogContext, strategy handler.LogHandlerStrategy) *log_handle {
	return &log_handle{
		context: logCtx,
		strategy: strategy,
	}
}

func (logHandle *log_handle) Process()  {
	logHandle.strategy.Process(logHandle.context)
}
