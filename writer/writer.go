package writer

type Writer interface {
	Write(wc chan string) bool
}
