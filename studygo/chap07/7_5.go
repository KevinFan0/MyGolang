package main

import (
	"fmt"
	"bufio"
	"bytes"
	"io"
)

const debug = true

func main()  {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {}  {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
