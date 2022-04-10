package reader

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

type nginx_access_reader struct {
	// nginx access file path
	path string
	// 读取
	readRate int
}

func NewAccessReader() *nginx_access_reader {
	return &nginx_access_reader{}
}

func (nr *nginx_access_reader) SetPath(path string) {
	nr.path = path
}

func (nr *nginx_access_reader) GetPath() string {
	return nr.path
}

func (r *nginx_access_reader) Read(readChan chan string) {
	f, err := os.Open(r.path)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	defer f.Close()

	// 从文件末尾读取
	f.Seek(0, 2)

	rd := bufio.NewReader(f)
	for true {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if err != nil {
			log.Panicf("%s", err.Error())
		}

		readChan <- line[:len(line) - 1]
	}

}
