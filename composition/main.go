package main

func main() {
	my := myDoing{}
	DoNothing(my)
}

type Reader interface {
	Read() (int64, error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

type myDoing struct {
	Doer
}

func (do myDoing) Do() error {
	return nil
}

func DoNothing(do Doer) {
	do.Do()
}

type Doer interface {
	Do() error
	Error() string
}
