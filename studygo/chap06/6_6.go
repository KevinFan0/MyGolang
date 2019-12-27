type IntSet struct {
	words []uint64
}

type Buffer struct {
	buf		 []byte
	initial	 [64]byte
}

func (b *Buffer) Grow(n int) {
	if b.buf == nil {
		b.buf = b.initial[:0]
	}
	if len(b.buf) + n > cap(b.buf) {
		buf := make([]byte, b.Len(), 2*cap(b.buf) + n)
		copy(buf, b.buf)
		b.buf = buf
	}
}

//下面的Counter类型允许调用方来增加counter变量的值，并且允许将这个值reset为0，但是不允许随便设置这个值(译注：因为压根就访问不到)：
type Counter struct { n int }
func (c *Counter) N() int { return c.n }
func (c *Counter) Increment() { c.n++ }
func (c *Counter) Reset()	{ c.n = 0}

package log

type Logger struct {
	flags	int
	prefix	string
}

func (l *Logger) Flags() int
func (l *Logger) SetFlags(flags int)
func (l *Logger) Prefix() string
func (l *Logger) SetPrefix(prefix string)