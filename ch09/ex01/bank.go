// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int)           // send amount to deposit
var balances = make(chan int)           // receive balance
var withdrawals = make(chan withdrawal) // withdraw amount

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	res := make(chan bool)
	withdrawals <- withdrawal{amount, res}
	return <-res
}

type withdrawal struct {
	amount int
	result chan bool
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdrawals:
			if w.amount > balance {
				w.result <- false
			} else {
				balance -= w.amount
				w.result <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
