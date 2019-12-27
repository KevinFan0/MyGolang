package main

import (
	"fmt"
	"io"
	"os"
	"bufio"
	"time"
)

var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = time.Second

var rwc io.ReadWriteCloser
rwc = os.Stdout
rwc = new(bytes.Buffer)

var any interface{}
any = true
any = 12.34