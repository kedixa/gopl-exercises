package bank

import (
	"fmt"
	"testing"
	"time"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("Deposit 200, balance is ", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		fmt.Println("Deposit 100, balance is ", Balance())
		done <- struct{}{}
	}()

	go func() {
		// wait a moument
		time.Sleep(1 * time.Microsecond)
		flag := Withdraw(100)
		fmt.Printf("Withdraw 100, %t, balance is %d\n", flag, Balance())
		done <- struct{}{}
	}()
	go func() {
		flag := Withdraw(400)
		fmt.Printf("Withdraw 400, %t, balance is %d\n", flag, Balance())
		done <- struct{}{}
	}()

	// Wait for transactions.
	<-done
	<-done
	<-done
	<-done

}
