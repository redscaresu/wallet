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

func TestSendMoney(t *testing.T) {
	t.Parallel()

	jim, err := wallet.CreateWallet("jim")
	if err != nil {
		t.Fatal(err)
	}

	rosy, err := wallet.CreateWallet("rosy")
	if err != nil {
		t.Fatal(err)
	}

	jimBalance, err := wallet.DepositWallet(jim, 5)
	if err != nil {
		t.Fatal(err)
	}

	wallet.SendMoney(jimBalance, rosy, 3)

	want := &wallet.Wallet{
		Name:    "rosy",
		Balance: 3,
	}

	_, got, err := wallet.SendMoney(jimBalance, rosy, 3)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
