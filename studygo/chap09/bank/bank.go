package bank

var balances = make(chan int)
var deposits = make(chan int)


func Deposit(amount int) { deposits <- amount }
func Balance() int { return <- balances }

func teller()  {
	var balance int				// balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init()  {
	go teller()					// start the monitor goroutine
}