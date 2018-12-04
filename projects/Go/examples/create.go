package main

import "fmt"
import "../proto/fbe"
import "../proto/proto"

func main() {
	// Create a new account using FBE model
	var account = proto.NewAccountModel(fbe.NewEmptyBuffer())
	modelBegin := account.CreateBegin()
	accountBegin, _ := account.Model().SetBegin()
	_ = account.Model().Uid.Set(1)
	_ = account.Model().Name.Set("Test")
	_ = account.Model().State.SetValue(proto.State_good)
	walletBegin, _ := account.Model().Wallet.SetBegin()
	_ = account.Model().Wallet.Currency.Set("USD")
	_ = account.Model().Wallet.Amount.Set(1000.0)
	account.Model().Wallet.SetEnd(walletBegin)
	account.Model().SetEnd(accountBegin)
	account.CreateEnd(modelBegin)
	if ok := account.Verify(); !ok {
		panic("verify error")
	}

	// Show the serialized FBE size
	fmt.Printf("FBE size: %d\n", account.Buffer().Size())

	// Access the account using the FBE model
	access := proto.NewAccountModel(fbe.NewAttachedBuffer(account.Buffer()))
	if ok := access.Verify(); !ok {
		panic("verify error")
	}

	accountBegin, _ = access.Model().GetBegin()
	uid, _ := access.Model().Uid.Get()
	name, _ := access.Model().Name.Get()
	state, _ := access.Model().State.Get()
	walletBegin, _ = access.Model().Wallet.GetBegin()
	walletCurrency, _ := access.Model().Wallet.Currency.Get()
	walletAmount, _ := access.Model().Wallet.Amount.Get()
	access.Model().Wallet.GetEnd(walletBegin)
	access.Model().GetEnd(accountBegin)

	// Show account content
	fmt.Println()
	fmt.Printf("account.uid = %v\n", uid)
	fmt.Printf("account.name = %v\n", name)
	fmt.Printf("account.state = %v\n", state)
	fmt.Printf("account.wallet.currency = %v\n", walletCurrency)
	fmt.Printf("account.wallet.amount = %v\n", walletAmount)
}