package main

import (
	"fmt"
	"github.com/tranvu711/qrpayment"
	"log"
)

func main() {
	// Create a VNPAY Wallet QR code
	qr, _ := qrpayment.NewQRPayment(&qrpayment.QRCode{
		QRType: qrpayment.QRTypeVNPayWallet,
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

	text, err := qr.GenerateText()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("QR Text:", text)

	err = qr.GeneratePNG("./vnpay_wallet_qr.png")
	if err != nil {
		log.Fatal(err)
		return
	}
}
