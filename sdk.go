package qrpayment

import (
	"errors"
	"fmt"
	"image/png"
	"os"
	"strings"

	"github.com/sigurn/crc16"
	qrcode "github.com/skip2/go-qrcode"
)

var (
	// validCurrencies contains the list of valid currency codes.
	validCurrencies = map[string]bool{
		CurrencyVND: true,
		CurrencyJPY: true,
		CurrencyKRW: true,
		CurrencyMYR: true,
		CurrencyCNY: true,
		CurrencyIDR: true,
		CurrencyPHP: true,
		CurrencySGD: true,
		CurrencyTHB: true,
	}
	// validCountries contains the list of valid country codes.
	validCountries = map[string]bool{
		CountryVN: true,
		CountryJP: true,
		CountryKR: true,
		CountryMY: true,
		CountryRC: true,
		CountryRI: true,
		CountryRP: true,
		CountrySG: true,
		CountryTH: true,
	}
	// validTransferTypes contains the list of valid transfer types.
	validTransferTypes = map[string]bool{
		TransferToAccount: true,
		TransferToCard:    true,
	}
)

// NewQRPayment initializes a new QRCode with default values if not provided.
func NewQRPayment(configs ...*QRCode) (*QRCode, error) {
	config := &QRCode{}
	if len(configs) > 0 {
		config = configs[0]
	}
	if config == nil {
		config = &QRCode{}
	}

	if config.AdditionalData == nil {
		config.AdditionalData = make(map[string]string)
	}
	if config.CountryCode == "" {
		config.CountryCode = CountryVN // Default to VN
	}
	if config.TransactionCurrency == "" {
		config.TransactionCurrency = CurrencyVND // Default to VND
	}
	if config.TransferType == "" {
		config.TransferType = TransferToAccount // Default to transfer to account
	}
	return config, nil
}

// SetBankBin sets the bank BIN for the QRCode.
func (q *QRCode) SetBankBin(bankBin string) {
	q.BankBin = bankBin
}

// SetAccountNo sets the account number for the QRCode.
func (q *QRCode) SetAccountNo(accountNo string) {
	q.AccountNo = accountNo
}

// SetAmount sets the amount for the QRCode.
func (q *QRCode) SetAmount(amount int) {
	q.Amount = &amount
}

// SetNote sets the note for the QRCode.
func (q *QRCode) SetNote(note string) {
	q.AdditionalData[PurposeOfTransaction] = note
}

// SetCardNo sets the card number for the QRCode.
func (q *QRCode) SetCardNo(cardNo string) {
	q.CardNo = cardNo
}

// SetTransactionCurrency sets the transaction currency for the QRCode.
func (q *QRCode) SetTransactionCurrency(currency string) error {
	if !validCurrencies[currency] {
		return errors.New("Invalid currency. Use predefined currency constants.")
	}
	q.TransactionCurrency = currency
	return nil
}

// SetCountryCode sets the country code for the QRCode.
func (q *QRCode) SetCountryCode(country string) error {
	if !validCountries[country] {
		return errors.New("Invalid country code. Use predefined country constants.")
	}
	q.CountryCode = country
	return nil
}

// SetBillNumber sets the bill number in the additional data of the QRCode.
func (q *QRCode) SetBillNumber(value string) {
	q.AdditionalData[BillNumber] = value
}

// SetCustomerMobileNumber sets the customer's mobile number in the additional data of the QRCode.
func (q *QRCode) SetCustomerMobileNumber(value string) {
	q.AdditionalData[CustomerMobileNumber] = value
}

// SetStoreLabel sets the store label in the additional data of the QRCode.
func (q *QRCode) SetStoreLabel(value string) {
	q.AdditionalData[StoreLabel] = value
}

// SetReferenceLabel sets the reference label in the additional data of the QRCode.
func (q *QRCode) SetReferenceLabel(value string) {
	q.AdditionalData[ReferenceLabel] = value
}

// SetCustomerLabel sets the customer label in the additional data of the QRCode.
func (q *QRCode) SetCustomerLabel(value string) {
	q.AdditionalData[CustomerLabel] = value
}

// SetTerminalLabel sets the terminal label in the additional data of the QRCode.
func (q *QRCode) SetTerminalLabel(value string) {
	q.AdditionalData[TerminalLabel] = value
}

// SetPurposeOfTransaction sets the purpose of transaction in the additional data of the QRCode.
func (q *QRCode) SetPurposeOfTransaction(value string) {
	q.AdditionalData[PurposeOfTransaction] = value
}

// SetAdditionalConsumerDataRequest sets the additional consumer data request in the additional data of the QRCode.
func (q *QRCode) SetAdditionalConsumerDataRequest(value string) {
	q.AdditionalData[AdditionalConsumerDataRequest] = value
}

// SetTransferType sets the transfer type for the QRCode.
func (q *QRCode) SetTransferType(transferType string) error {
	if !validTransferTypes[transferType] {
		return errors.New("Invalid transfer type. Use TransferToAccount or TransferToCard")
	}
	q.TransferType = transferType
	return nil
}

// GenerateText generates the QR code text based on the QRCode data.
func (q *QRCode) GenerateText() (string, error) {
	if q.TransferType == TransferToAccount && q.AccountNo == EmptyString && q.BankBin == EmptyString {
		return EmptyString, errors.New("BankBin and AccountNo are required")
	}

	if q.TransferType == TransferToCard && q.CardNo == EmptyString && q.BankBin == EmptyString {
		return EmptyString, errors.New("BankBin and CardNo are required")
	}

	qrType := StaticQR
	if q.Amount != nil && *q.Amount > 0 {
		qrType = DynamicQR
	}

	fields := map[string]interface{}{
		"00": "01",
		"01": qrType,
		"38": map[string]interface{}{
			"00": VietQRService,
			"01": map[string]string{
				"00": q.BankBin,
				"01": func() string {
					if q.TransferType == TransferToAccount {
						return q.AccountNo
					}
					return q.CardNo
				}(),
			},
			"02": q.TransferType,
		},
		"53": q.TransactionCurrency,
		"58": q.CountryCode,
	}

	if q.Amount != nil && *q.Amount > 0 {
		fields["54"] = fmt.Sprintf("%d", *q.Amount)
	}
	if len(q.AdditionalData) > 0 {
		subFields := map[string]string{}
		for key, value := range q.AdditionalData {
			subFields[key] = value
		}
		fields["62"] = subFields
	}

	var tlv []string
	for key, value := range fields {
		switch v := value.(type) {
		case string:
			tlv = append(tlv, fmt.Sprintf("%s%02d%s", key, len(v), v))
		case map[string]string:
			var subTlv []string
			for subKey, subValue := range v {
				subTlv = append(subTlv, fmt.Sprintf("%s%02d%s", subKey, len(subValue), subValue))
			}
			joinedSub := strings.Join(subTlv, "")
			tlv = append(tlv, fmt.Sprintf("%s%02d%s", key, len(joinedSub), joinedSub))
		case map[string]interface{}:
			var subTlv []string
			for subKey, subValue := range v {
				switch subVal := subValue.(type) {
				case string:
					subTlv = append(subTlv, fmt.Sprintf("%s%02d%s", subKey, len(subVal), subVal))
				case map[string]string:
					var innerTlv []string
					for innerKey, innerValue := range subVal {
						innerTlv = append(innerTlv, fmt.Sprintf("%s%02d%s", innerKey, len(innerValue), innerValue))
					}
					joinedInner := strings.Join(innerTlv, "")
					subTlv = append(subTlv, fmt.Sprintf("%s%02d%s", subKey, len(joinedInner), joinedInner))
				}
			}
			joinedSub := strings.Join(subTlv, "")
			tlv = append(tlv, fmt.Sprintf("%s%02d%s", key, len(joinedSub), joinedSub))
		}
	}

	data := strings.Join(tlv, "")
	str := fmt.Sprintf("%s%s", data, CRC16Tag)
	checksum := crc16.Checksum([]byte(str), crc16.MakeTable(crc16.CRC16_CCITT_FALSE))
	return fmt.Sprintf("%s%04X", str, checksum), nil
}

// GeneratePNG generates a PNG image of the QR code and saves it to the specified filename.
func (q *QRCode) GeneratePNG(filename string) error {
	data, err := q.GenerateText()
	if err != nil {
		return err
	}

	qr, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, qr.Image(256))
}
