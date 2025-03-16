package qrpayment

import (
	"errors"
)

func (q *QRCode) isValidVieQr() error {
	// Validate transfer to account
	if q.TransferType == TransferToAccount && q.AccountNo == EmptyString && q.BankBin == EmptyString {
		return errors.New("BankBin and AccountNo are required for transfer to account")
	}

	// Validate transfer to card
	if q.TransferType == TransferToCard && q.CardNo == EmptyString && q.BankBin == EmptyString {
		return errors.New("BankBin and CardNo are required for transfer to card")
	}

	// Validate amount
	if q.Amount != nil && *q.Amount < 0 {
		return errors.New("amount must be greater than 0")
	}

	return nil
}

// GenerateText generates the QR code text based on the QRCode data.
func (q *QRCode) generateVietQRText() (string, error) {
	err := q.isValidVieQr()
	if err != nil {
		return EmptyString, err
	}

	fields := q.generateCommonQRData()
	fields[FieldVietQRIdentity] = map[string]interface{}{
		FiledVietQRIdentitySubService: VietQRService,
		FiledVietQRIdentitySubCustomer: map[string]string{
			FiledVietQRIdentitySubCustomerBankBin: q.BankBin,
			FiledVietQRIdentitySubCustomerBankInfo: func() string {
				if q.TransferType == TransferToAccount {
					return q.AccountNo
				}
				return q.CardNo
			}(),
		},
		FiledVietQRIdentitySubCustomerBankInfoType: q.TransferType,
	}

	data := generateTLVFormat(fields)
	qrText := generateChecksumWithCRC16(data)
	return qrText, nil
}
