package bank

var deposits = make(chan int)   // send amount to deposit
var balances = make(chan int)   // receive balance
var withdraws = make(chan item) // send amount to withdraw

type item struct {
	amount int
	flag   chan bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	var flag = make(chan bool)
	withdraws <- item{amount, flag}
	return <-flag
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case it := <-withdraws:
			if it.amount <= balance {
				balance -= it.amount
				it.flag <- true
			} else {
				it.flag <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
