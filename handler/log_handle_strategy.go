package handler

// 定义处理日志策略接口
type LogHandlerStrategy interface {
	Process(logCtx *LogContext)
}

