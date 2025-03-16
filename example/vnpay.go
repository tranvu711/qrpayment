package main

import (
	"fmt"
	"log"
	"qrpayment"
	"strconv"
	"time"
)

func main() {
	// Create a VNPAY Merchant QR code
	qr, _ := qrpayment.NewQRPayment(&qrpayment.QRCode{
		QRType: qrpayment.QRTypeVNPay,
	})

	// Set the merchant ID
	err := qr.SetMerchantId("0316213597", true)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the merchant name
	err = qr.SetMerchantName("WESCAN VIETNAM")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the merchant city
	err = qr.SetCityName("HCM")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the merchant postal code
	err = qr.SetPostalCode("700000")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the terminal ID
	err = qr.SetTerminalLabel("VUTNT")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the store label
	err = qr.SetStoreLabel("TRAN NGO TUAN VU")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the mcc code
	err = qr.SetMCC("4816")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Set bill number
	err = qr.SetBillNumber(strconv.FormatInt(time.Now().UTC().Unix(), 10))
	if err != nil {
		log.Fatal(err)
		return
	}
	// Set purpose of transaction, aka note
	err = qr.SetPurposeOfTransaction("Paymentnote")
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

	err = qr.GeneratePNG("./vnpayqr.png")
	if err != nil {
		log.Fatal(err)
		return
	}
}
