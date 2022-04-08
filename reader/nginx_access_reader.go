package reader

type Nginx_access_reader struct {
}

func (r *Nginx_access_reader) Read(rc chan string) bool {
	// 读取文件内容
	line := "message"
	rc <- line

	return true
}
