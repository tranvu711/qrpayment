package qrpayment

import (
	"errors"
	"strconv"
	"strings"
)

func (q *QRCode) isValidMomoWallet() error {
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
func (q *QRCode) generateMoMoWalletQRText() (string, error) {

	err := q.isValidVNPayWallet()
	if err != nil {
		return EmptyString, err
	}
	if q.Amount == nil {
		defaultAmount := 0
		q.Amount = &defaultAmount
	}

	data := []string{
		"2",
		"99",
		q.AdditionalData[CustomerMobileNumber],
		q.AdditionalData[CustomerMobileNumber],
		"",
		"0",
		"0",
		strconv.Itoa(*q.Amount),
		q.AdditionalData[PurposeOfTransaction],
		"transfer_myqr",
	}

	qrText := strings.Join(data, "|")
	return qrText, nil
}
