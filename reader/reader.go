package reader

type Reader interface {
	Read(rc chan string) bool
}

