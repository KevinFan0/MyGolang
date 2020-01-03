// 下面这段逻辑和net/http包中web服务器负责写入HTTP头字段（例如："Content-type:text/html）的部分相似。io.Writer接口类型的变量w代表HTTP响应；写入它的字节最终被发送到某个人的web浏览器上。
func writeHeader(w io.Writer, contentType string) error {
	if _, err := w.Write([]byte("Content-Type: ")); err != nil {
		return err
	}
	if _, err := w.Write([]byte(contentType)); err != nil {
		return err
	}
// ...
}


func writeString(w io.Writer, s string) (n int, err error) {
	type stringWrite interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWrite); ok {
		return sw.WriteString(s)			// avoid a copy
	}
	return w.Write([]byte(s))				// allocate temporary copy
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}
}

// fmt.Fprintf函数怎么从其它所有值中区分满足error或者fmt.Stringer接口的值。在fmt.Fprintf内部，有一个将单个操作对象转换成一个字符串的步骤，像下面这样：
func formatOneValue(x interface{}) string {
	if err, ok := x.(error); ok {
		return err.Error()
	}
	if str, ok := x.(Stringer); ok {
		return str.String()
	}
}