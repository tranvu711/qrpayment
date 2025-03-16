package main

import (
	"fmt"
	"log"
	"qrpayment"
)

func main() {
	// Create a VNPAY Merchant QR code
	qr, _ := qrpayment.NewQRPayment(&qrpayment.QRCode{
		QRType: qrpayment.QRTypeVNPayWallet,
	})

	err := qr.SetCustomerMobileNumber("0329092681")
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
