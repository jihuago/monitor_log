package test

import (
	"fmt"
	"github.com/jihuago/monitor_log/reader"
	"github.com/jihuago/monitor_log/tools"
	"os"
	"sync"
	"testing"
)

// 并发往 access.log 写入数据 go test -v test/write_log_test.go
func TestConcurrentWriteToLog(t *testing.T) {

	rd := reader.NewAccessReader()
	rd.SetPath("../example/access.txt2")

	file, err := os.OpenFile(rd.GetPath(), os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	defer file.Close()

	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			_, err := file.WriteString(fmt.Sprintf("%s\n", tools.RandomString(5)))
			if err != nil {
				t.Fatalf("%s", err.Error())
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}