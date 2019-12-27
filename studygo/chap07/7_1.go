package fmt

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error) {}
func Printf(format string, args ...interface{}) (int, error)  {
	return Fprintf(os.Stdout, format, args...)
}

func Sprintf(format string, args ...interface{}) string {
	var buf bytes.Buffer
	Fprintf(&buf, format, args...)
	return buf.String()
}

// 下面*ByteCounter类型里的Write方法，仅仅在丢失写向它的字节前统计它们的长度。(在这个+=赋值语句中，让len(p)的类型和*c的类型匹配的转换是必须的。)
type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

var c ByteCounter
c.Write([]byte("hello"))
fmt.Println(c)
c = 0
var name = "Dolly"
fmt.Fprintf(&c, "hello, %s\n", name)
fmt.Println(c)