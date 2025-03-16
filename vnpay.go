package qrpayment

import (
	"errors"
)

func (q *QRCode) isValidVNPay() error {
	// Validate amount
	if q.Amount != nil && *q.Amount < 0 {
		return errors.New("amount must be greater than 0")
	}

	validations := []struct {
		value    string
		errorMsg string
	}{
		{q.MCC, "MCC is required"},
		{q.MerchantName, "Merchant name is required"},
		{q.MerchantID, "Merchant ID is required"},
		{q.TerminalID, "Terminal ID is required"},
		{q.TerminalName, "Terminal name is required"},
		{q.CityName, "City name is required"},
		{q.PostalCode, "Postal code is required"},
	}

	for _, v := range validations {
		if v.value == EmptyString {
			return errors.New(v.errorMsg)
		}
	}

	return nil
}

// GenerateText generates the QR code text based on the QRCode data.
func (q *QRCode) generateVNPayQRText() (string, error) {

	err := q.isValidVNPay()
	if err != nil {
		return EmptyString, err
	}

	fields := q.generateCommonQRData()
	fields[FieldVNPayIdentity] = map[string]interface{}{
		FieldVNPayIdentitySubService:    q.MerchantIdentify,
		FieldVNPayIdentitySubMerchantId: q.MerchantID,
	}

	data := generateTLVFormat(fields)
	qrText := generateChecksumWithCRC16(data)
	return qrText, nil
}
