package main

import (
	"log"
	"github.com/usmon1983/wallet/pkg/wallet"
)

func main()  {
	vc := &wallet.Service{}
	//vc.ImportFromFile("data/exp.txt")
	//vc.Export("/data")
	//vc.Import("/data")
	paymentsExport, err := vc.ExportAccountHistory(1)
	if err != nil {
		log.Print(err)
		return
	}
	vc.HistoryToFiles(paymentsExport, "/data", 2)
}