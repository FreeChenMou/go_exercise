package main

//!+
import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
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

//exercise9.1
func Withdraw(amout int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amout)
	if balance < 0 {
		Deposit(amout)
		return false
	}
	return true
}
