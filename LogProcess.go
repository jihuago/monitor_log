package monitor_log

import (
	"github.com/jihuago/monitor_log/reader"
	"github.com/jihuago/monitor_log/writer"
	"strings"
)

type LogProcess struct {
	Rc chan string
	Wc chan string
	Reader reader.Reader
	Writer writer.Writer
}

func (lp *LogProcess) Process()  {
	data := <- lp.Rc
	lp.Wc <- strings.ToUpper(data)
}

func Process() *LogProcess {
	lp := &LogProcess{
		Rc: make(chan string),
		Wc: make(chan string),
		Reader: &reader.Nginx_access_reader{},
		Writer: &writer.Write_to_influxdb{},
	}

	return lp
}