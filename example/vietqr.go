package main

import (
	"fmt"
	"log"
	"qrpayment"
)

func main() {
	qr, _ := qrpayment.NewQRPayment()
	qr.SetBankBin(qrpayment.BANKS.Vietcombank.BIN)
	qr.SetAccountNo("0441000670827")
	qr.SetAmount(100000)
	qr.SetPurposeOfTransaction("Payment note")
	err := qr.SetTransferType(qrpayment.TransferToAccount)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = qr.SetTransactionCurrency(qrpayment.CurrencyVND)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = qr.SetCountryCode(qrpayment.CountryVN)
	if err != nil {
		log.Fatal(err)
		return
	}

	text, err := qr.GenerateText()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("QR Text:", text)

	err = qr.GeneratePNG("./vietqr.png")
	if err != nil {
		log.Fatal(err)
		return
	}
}
