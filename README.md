# wallet

Pre-reqs

`go get github.com/google/go-cmp/cmp`

In order to run the program

`go run cmd/main.go`

In order to test the program

`go test`


Notes
- model a wallet as a struct with only 2 attributes, name and an account balance.  This is at minimum what I needed to implemement the program.
- deliberately chose not to implement pesistence, wanted to focus on code architecture, structure and tests.
- program should be read from cmd/main, wallet is a reusable library consumed by itself.
- tests should dogfood as much functionality of the library as possible, this what I aimed to do.

Areas to be improved
- wallet has no way of easily differenting users from one another, it is conceivable there will be 2 wallet owners with the same name.  Wallet struct should also include a UID, an immutable hash that differentiates their account from another or maybe an account number.
- currently only "ReturnBalance()" is implemented as a method. DepositWallet(), WithdrawWallet() could all potentially be implemented as a method straight off of the struct.  Arguable the others too, each operation is tightly coupled to the struct right now.
- DepositWallet() and WithdrawWallet() could potentionally only return an error instead of an errror and new pointer to the struct.
- SendMoney() error handling needs to be improved.  In the program for demonstration purposes I just call it directly but I do not do anything with any errors being returned.  That needs thought and needs to to be improved.
- There should be some pre-flight checks before money is attempted to be deposited, withdrawn or sent.  At the moment the program attempts to do them, it does exit if output was unexpected but there could be ways of checking before the operation is done.  For example DepositWallet() could check that the money is positive.  Also there could a concept of "pre-auth" where money is taken out of your account then attempted to be sent, at the moment if money is succesfully withdrawn but not sent into it has no way of going back into your wallet.
