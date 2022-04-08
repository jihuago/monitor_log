package main

import (
	"github.com/jihuago/monitor_log"
	"time"
)

func main()  {
	lp := monitor_log.Process()

	go lp.Reader.Read(lp.Rc)
	go lp.Process()
	go lp.Writer.Write(lp.Wc)

	time.Sleep(1 * time.Second)
}
