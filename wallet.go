package wallet

import (
	"fmt"
)

type Wallet struct {
	Name    string
	Balance int
}

func CreateWallet(name string) (*Wallet, error) {
	wallet := &Wallet{Name: name,
		Balance: 0}
	return wallet, nil
}

func DepositWallet(wallet *Wallet, money int) (*Wallet, error) {
	wallet.Balance += money
	return wallet, nil
}

func WithdrawWallet(wallet *Wallet, money int) (*Wallet, error) {
	wallet.Balance -= money
	return wallet, nil
}

func returnBalance(wallet *Wallet) {
	fmt.Printf("hello %s, your balance is %v \n", wallet.Name, wallet.Balance)
}

func sendMoney(sourceWallet *Wallet, destWallet *Wallet, money int) (*Wallet, *Wallet, error) {
	sourceWallet, err := WithdrawWallet(sourceWallet, money)
	if err != nil {
		fmt.Println("sorry there has been an error")
	}

	destWallet, err = DepositWallet(destWallet, money)
	if err != nil {
		fmt.Println("sorry there has been an error")
	}
	return sourceWallet, destWallet, err
}

func RunWallet() {
	jim, err := CreateWallet("jim")
	if err != nil {
		fmt.Println("sorry there has been an error")
	}
	fmt.Println(jim)

	addCashJim, err := DepositWallet(jim, 5)
	if err != nil {
		fmt.Println("sorry there has been an error")
	}
	fmt.Println(addCashJim)

	withdrawCashJim, err := WithdrawWallet(jim, 1)
	if err != nil {
		fmt.Println("sorry there has been an error")
	}
	fmt.Println(withdrawCashJim)

	returnBalance(jim)

	rosy, err := CreateWallet("rosy")
	if err != nil {
		fmt.Println("sorry there has been an error")
	}

	fmt.Println("send 3 to rosy from jim....")
	sendMoney(jim, rosy, 3)
	returnBalance(jim)
	returnBalance(rosy)
}
