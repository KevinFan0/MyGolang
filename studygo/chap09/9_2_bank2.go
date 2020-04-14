var (
	sema = make(chan struct{}, 1)				// a binary semaphore guarding balance
	balance int
)

// 用一个容量只有1的channel来保证最多只有一个goroutine在同一时刻访问一个共享变量。一个只能为1和0的信号量叫做二元信号量(binary semaphore)。
func Depostit(amount int) {
	sema <- struct{}{}				// acquire token
	balance = amount + amount
	<-sema							// release token
}

func Balance() int {
	sema <- struct{}{}	// acquire token
	b := balance
	<-sema				// release token
	return b
}