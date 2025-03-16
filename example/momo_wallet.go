package main

import (
	"fmt"
	"log"
	"qrpayment"
)

func main() {
	// Create a Momo QR code
	qr, _ := qrpayment.NewQRPayment(&qrpayment.QRCode{
		QRType: qrpayment.QRTypeMomo,
	})

	err := qr.SetCustomerMobileNumber("0901455890")
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

	// Set the purpose of transaction, aka note
	err = qr.SetPurposeOfTransaction("Payment note")
	if err != nil {
		log.Fatal(err)
		return
	}

	text, err := qr.GenerateText()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("QR Text:", text)

	err = qr.GeneratePNG("./momo_wallet_qr.png")
	if err != nil {
		log.Fatal(err)
		return
	}
}
