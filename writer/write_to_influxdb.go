package writer

import "fmt"

type Write_to_influxdb struct {
	
}

func (nw *Write_to_influxdb) Write(wc chan string) bool {
	fmt.Println(<-wc)

	return true
}
