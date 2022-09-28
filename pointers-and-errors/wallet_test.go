package pointers_and_errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}
		err := wallet.Withdraw(Bitcoin(30))

		assert.NoError(t, err)
		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(35)}
		err := wallet.Withdraw(Bitcoin(100))

		if assert.Error(t, err) {
			assert.Equal(t, ErrInsufficientFunds, err)

		}
		assertBalance(t, wallet, Bitcoin(35))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	assert.Equal(t, want, got)
}
