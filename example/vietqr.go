package main

import (
	"fmt"
	"log"
	"qrpayment"
)

func main() {
	qr, _ := qrpayment.NewQRPayment(&qrpayment.QRCode{
		QRType: qrpayment.QRTypeVietQR,
	})

	// Set the transfer type
	err := qr.SetTransferType(qrpayment.TransferToAccount)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the bank BIN
	err = qr.SetBankBin(qrpayment.BANKS.Vietcombank.BIN)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the account number
	err = qr.SetAccountNo("0441000670827")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the purpose of transaction, aka note
	err = qr.SetPurposeOfTransaction("Payment note")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the amount
	err = qr.SetAmount(100000)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Create the QR code - returns the QR code text
	text, err := qr.GenerateText()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("QR Text:", text)

	// Generate the QR code image
	err = qr.GeneratePNG("./vietqr.png")
	if err != nil {
		log.Fatal(err)
		return
	}
}
