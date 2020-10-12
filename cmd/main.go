package main

import (
	"github.com/usmon1983/wallet/pkg/wallet"
)

func main()  {
	vc := &wallet.Service{}
	vc.ImportFromFile("data/exp.txt")
}