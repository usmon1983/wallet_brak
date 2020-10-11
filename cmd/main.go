package main

import (
	"github.com/usmon1983/wallet/pkg/wallet"

)
func main()  {
	vc := &wallet.Service{}
	vc .RegisterAccount("+992000000077")
	vc.Deposit(1, 10)
	vc .RegisterAccount("+992000000088")
	vc.Deposit(2, 20)
	vc .RegisterAccount("+992000000099")
	vc.Deposit(3, 30)
	vc.ExportToFile("data/exp.txt")
}