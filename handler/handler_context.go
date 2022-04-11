package handler

// 日志处理上下文
type LogContext struct {
	Rch chan string
	Wch chan string
}
