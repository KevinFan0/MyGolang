/*

package xml

type Name struct {
	Local	string
}

type Attr struct {
	Name	Name
	Value	string
}

type Token interface {}
type StartElement struct {
	Name	Name
	Attr	[]Attr
}
type EndElement struct { Name Name}

type CharData []byte
type Comment []byte
type Decoder struct { //,,, }

func NewDecoder(io.Reader) *Decoder {
	
}

func (*Decoder) Token() (Token, error) {
	
}
*/

// xmlselect程序获取和打印在一个XML文档树中确定的元素下找到的文本。使用上面的API，它可以在输入上一次完成它的工作而从来不要具体化这个文档树。
package main

import (
	"fmt"
	"encoding/xml"
	"io"
	"os"
	"strings"
)

func main()  {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)		// push
		case xml.EndElement:
			stack = stack[:len(stack)-1]				// pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}