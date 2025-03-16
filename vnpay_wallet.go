package qrpayment

import (
	"errors"
	"fmt"
)

const (
	FieldVNPayWalletIdentity = "07"
	FieldVNPayWalletService  = "03"
	FieldVNPayWallerCountry  = "10"
	FieldVNPAYWalletAmount   = "09"
)

func (q *QRCode) isValidVNPayWallet() error {
	// Validate amount
	if q.Amount != nil && *q.Amount < 0 {
		return errors.New("amount must be greater than 0")
	}
	if q.AdditionalData[CustomerMobileNumber] == EmptyString {
		return errors.New("CustomerMobileNumber is required")
	}
	return nil
}

// GenerateText generates the QR code text based on the QRCode data.
func (q *QRCode) generateVNPayWalletQRText() (string, error) {

	err := q.isValidVNPayWallet()
	if err != nil {
		return EmptyString, err
	}

	fields := map[string]interface{}{
		FieldQRVersion:           QRFormatVersion,
		FieldQRQRType:            VNPayWalletQR,
		FieldVNPayWalletIdentity: q.AdditionalData[CustomerMobileNumber],
		FieldVNPayWalletService:  VNPAYWALLETService,
		FieldVNPayWallerCountry:  q.CountryCode,
		"04":                     "F5",
		"02":                     "10",
		FieldVNPAYWalletAmount:   " ",
	}
	if q.Amount != nil {
		fields[FieldVNPAYWalletAmount] = fmt.Sprintf("%d", *q.Amount)
	}

	data := generateTLVFormat(fields)
	qrText := generateChecksumWithCRC16(data)
	return qrText, nil
}
