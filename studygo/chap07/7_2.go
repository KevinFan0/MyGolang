package io

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReaderWriter interface {
	Reader
	Writer
}

type ReadWriterCloser interface {
	Reader
	Writer
	Closer
}

// 接口内嵌
type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 混合风格
type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Writer
}