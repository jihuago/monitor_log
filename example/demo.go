package main

import (
	"github.com/jihuago/monitor_log"
	"github.com/jihuago/monitor_log/reader"
	"github.com/jihuago/monitor_log/writer"
)

func main()  {
	rd := reader.NewAccessReader()
	rd.SetPath("./example/access.txt")

	wd := &writer.Write_to_influxdb{}

	lp := monitor_log.New()

	// 步骤：读、处理、写入
	go rd.Read(lp.Rch)
	go lp.Process(lp.Rch)
	go wd.Write(lp.Wch)

	for true {

	}
}
