var mu sync.Mutex
var balance int

func Balance() int {
	mu.RLockk()				// readers lock
	defer mu.RUnlock()
	return balance
}