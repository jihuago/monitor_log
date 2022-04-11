package handler

import "strings"

// 专门负责处理nginx access log 的
type NginxAccessLogHandle struct {}

func (NginxAccessLog *NginxAccessLogHandle) Process(logCtx *LogContext) {
	for v := range logCtx.Rch {
		logCtx.Wch <- strings.ToUpper(v)
	}
}
