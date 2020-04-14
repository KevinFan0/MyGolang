import (
	"sync"
)

var (
	mu		sync.Mutex
	balance	int
)

func Deposit(amount int)  {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func Balance2() int {
    mu.Lock()
    defer mu.Unlock()
    return balance
}

// NOTE: not atomic!
// 成功的时候，它会正确地减掉余额并返回true。但如果银行记录资金对交易来说不足，那么取款就会恢复余额，并返回false
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if Balance() < 0 {
		deposit(amount)
		return false				// insufficient funds
	}
	return true
}


func Deposit(amount int)  {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func deposit(amount int) { balance += amount}
