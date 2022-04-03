package wallet

import (
	"errors"
	"fmt"
	"os"
)

type Wallet struct {
	Name    string
	Balance int
}

func CreateWallet(name string) (*Wallet, error) {
	wallet := &Wallet{Name: name,
		Balance: 0}
	if wallet.Balance != 0 {
		return nil, fmt.Errorf("wallet balance is %v, should be 0", wallet.Balance)
	}
	return wallet, nil
}

func DepositWallet(wallet *Wallet, money int) (*Wallet, error) {
	oldBalance := wallet.Balance
	wallet.Balance += money
	if wallet.Balance <= oldBalance {
		return nil, fmt.Errorf("new balance %v is not greater than the old balance %v", wallet.Balance, oldBalance)
	}
	return wallet, nil
}

func WithdrawWallet(wallet *Wallet, money int) (*Wallet, error) {

	oldBalance := wallet.Balance
	wallet.Balance -= money
	if wallet.Balance >= oldBalance {
		return nil, fmt.Errorf("new balance %v is not less than the old balance %v", wallet.Balance, oldBalance)
	}
	return wallet, nil
}

func (wallet *Wallet) ReturnBalance() (string, error) {
	balanceStatement := fmt.Sprintf("hello %s, your balance is %v \n", wallet.Name, wallet.Balance)
	if balanceStatement == "" {
		return "", errors.New("unable to return the balance")
	}
	return balanceStatement, nil
}

func SendMoney(sourceWallet *Wallet, destWallet *Wallet, money int) (*Wallet, *Wallet, error) {
	sourceWallet, err := WithdrawWallet(sourceWallet, money)
	if err != nil {
		return sourceWallet, nil, err
	}

	destWallet, err = DepositWallet(destWallet, money)
	if err != nil {
		return nil, destWallet, err
	}
	return sourceWallet, destWallet, nil
}

func RunWallet() {
	jim, err := CreateWallet("jim")
	if err != nil {
		fmt.Printf("sorry there has been an error: %v \n", err)
		os.Exit(1)
	}
	fmt.Println(jim)

	addCashJim, err := DepositWallet(jim, 5)
	if err != nil {
		fmt.Printf("sorry there has been an error: %v \n", err)
		os.Exit(1)
	}
	fmt.Println(addCashJim)

	withdrawCashJim, err := WithdrawWallet(jim, 1)
	if err != nil {
		fmt.Printf("sorry there has been an error: %v \n", err)
		os.Exit(1)
	}
	fmt.Println(withdrawCashJim)

	balance, err := jim.ReturnBalance()
	if err != nil {
		fmt.Printf("sorry there has been an error: %v \n", err)
		os.Exit(1)
	}
	fmt.Println(balance)

	rosy, err := CreateWallet("rosy")
	if err != nil {
		fmt.Println("sorry there has been an error")
	}

	fmt.Println("send 3 from jim to rosy....")

	SendMoney(jim, rosy, 3)
	fmt.Println("money sent to rosy")
	fmt.Printf("\n")

	fmt.Println("new balances:")
	jimNewBalance, err := jim.ReturnBalance()
	if err != nil {
		fmt.Printf("sorry there has been an error: %v \n", err)
		os.Exit(1)
	}
	fmt.Println(jimNewBalance)

	rosyNewBalance, err := rosy.ReturnBalance()
	if err != nil {
		fmt.Printf("sorry there has been an error: %v \n", err)
		os.Exit(1)
	}
	fmt.Println(rosyNewBalance)
}
