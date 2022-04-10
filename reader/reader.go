package reader

type Reader interface {
	Read(chan string)
}

