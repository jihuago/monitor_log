package monitor_log

import (
	"strings"
)

type logProcess struct {
	Rch chan string
	Wch chan string
}

func (lp *logProcess) Process(ch chan string)  {
	for v := range ch {
		lp.Wch <- strings.ToUpper(v)
	}
}

func New() *logProcess {
	return &logProcess{
		Rch: make(chan string),
		Wch: make(chan string),
	}
}