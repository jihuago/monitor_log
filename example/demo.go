package main

import (
	"github.com/jihuago/monitor_log"
	"github.com/jihuago/monitor_log/handler"
	"github.com/jihuago/monitor_log/reader"
	"github.com/jihuago/monitor_log/writer"
)

func main()  {
	exampleStrategy()

	for true {

	}
}

func exampleStrategy()  {

	handle_ctx := &handler.LogContext{
		Rch: make(chan string),
		Wch: make(chan string),
	}

	// 读
	rd := reader.NewAccessReader()
	rd.SetPath("./example/access.txt")
	go rd.Read(handle_ctx.Rch)

	ng_access_log_handle := &handler.NginxAccessLogHandle{}
	// 策略模式的
	logHandler := monitor_log.NewLogHandle(handle_ctx, ng_access_log_handle)
	go logHandler.Process()

	wd := &writer.Write_to_influxdb{}
	go wd.Write(handle_ctx.Wch)
}
