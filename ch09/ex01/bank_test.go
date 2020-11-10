package bank_test

import (
	"fmt"
	"testing"

	bank "github.com/yyamada12/go_exercises/ch09/ex01"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	if got, want := bank.Withdraw(400), false; got != want {

	}
}

func TestWithdraw(t *testing.T) {

	bank.Deposit(200)

	if bank.Withdraw(300) {
		t.Errorf("Withdraw(300) should not return true if amount is 200")
	}
	if got, want := bank.Balance(), 200; got != want {
		t.Errorf("After Withdraw(300) the amount got %d, want %d", got, want)
	}

	if !bank.Withdraw(50) {
		t.Errorf("Withdraw(50) should return true if amount is 200")
	}
	if got, want := bank.Balance(), 150; got != want {
		t.Errorf("After Withdraw(50) amount got %d, want %d", got, want)
	}

}
