package wallet_test

import (
	"testing"
	"wallet"

	"github.com/google/go-cmp/cmp"
)

func TestWalletCreation(t *testing.T) {
	t.Parallel()
	got, err := wallet.CreateWallet("jim")
	if err != nil {
		t.Fatal(err)
	}

	want := &wallet.Wallet{
		Name:    "jim",
		Balance: 0,
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestWalletDeposit(t *testing.T) {
	t.Parallel()

	jim, err := wallet.CreateWallet("jim")
	if err != nil {
		t.Fatal(err)
	}

	got, err := wallet.DepositWallet(jim, 5)
	if err != nil {
		t.Fatal(err)
	}

	want := &wallet.Wallet{
		Name:    "jim",
		Balance: 5,
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
