package qrpayment

import (
	"errors"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"image/png"
	"os"
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

	// validQRTypes contains the list of valid QR types.
	validQRTypes = map[string]bool{
		QRTypeVietQR:      true,
		QRTypeVNPay:       true,
		QRTypeVNPayWallet: true,
		QRTypeMomo:        true,
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

	// Default to VN
	if config.CountryCode == EmptyString {
		config.CountryCode = CountryVN // Default to VN
	}

	// Default to VND
	if config.TransactionCurrency == EmptyString {
		config.TransactionCurrency = CurrencyVND // Default to VND
	}

	// Default to transfer to account
	if config.TransferType == EmptyString {
		config.TransferType = TransferToAccount // Default to transfer to account
	}

	// Default to VietQR
	if config.QRType == EmptyString {
		config.QRType = QRTypeVietQR
	}

	return config, nil
}

// SetQRType sets the QR type for the QRCode.
func (q *QRCode) SetQRType(qrType string) error {
	if _, ok := validQRTypes[qrType]; !ok {
		return errors.New("Invalid QR type. Use predefined QR type constants.")
	}
	q.QRType = qrType
	return nil
}

// SetBankBin sets the bank BIN for the QRCode.
func (q *QRCode) SetBankBin(bankBin string) error {
	if bankBin == EmptyString {
		return errors.New("Bank BIN is required")
	}
	q.BankBin = bankBin
	return nil
}

// SetAccountNo sets the account number for the QRCode.
func (q *QRCode) SetAccountNo(accountNo string) error {
	if accountNo == EmptyString {
		return errors.New("Account number is required")
	}
	q.AccountNo = accountNo
	return nil
}

// SetAmount sets the amount for the QRCode.
func (q *QRCode) SetAmount(amount int) error {
	if amount < 0 {
		return errors.New("Amount must be greater than 0")
	}
	q.Amount = &amount
	return nil
}

// SetNote sets the note for the QRCode.
func (q *QRCode) SetNote(note string) error {
	if note == EmptyString {
		return errors.New("Note is required")
	}
	q.AdditionalData[PurposeOfTransaction] = note
	return nil
}

// SetCardNo sets the card number for the QRCode.
func (q *QRCode) SetCardNo(cardNo string) error {
	if cardNo == EmptyString {
		return errors.New("Card number is required")
	}
	q.CardNo = cardNo
	return nil
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

// SetCityName sets the city name for the QRCode.
func (q *QRCode) SetCityName(city string) error {
	if city == EmptyString {
		return errors.New("City is required")
	}
	if len(city) > CityMaxLength {
		return errors.New(fmt.Sprintf("City must be less than %d characters", CityMaxLength))
	}
	q.CityName = city
	return nil
}

// SetPostalCode sets the postal code for the QRCode.
func (q *QRCode) SetPostalCode(postalCode string) error {
	if postalCode == EmptyString {
		return errors.New("Postal code is required")
	}
	if len(postalCode) > PostalCodeMaxLength {
		return errors.New(fmt.Sprintf("Postal code must be less than %d characters", PostalCodeMaxLength))
	}
	q.PostalCode = postalCode
	return nil
}

// SetBillNumber sets the bill number in the additional data of the QRCode.
func (q *QRCode) SetBillNumber(value string) error {
	if value == EmptyString {
		return errors.New("Bill number is required")
	}
	if len(value) > BillNumberMaxLength {
		return errors.New(fmt.Sprintf("Bill number must be less than %d characters", BillNumberMaxLength))
	}
	q.AdditionalData[BillNumber] = value
	return nil
}

// SetCustomerMobileNumber sets the customer's mobile number in the additional data of the QRCode.
func (q *QRCode) SetCustomerMobileNumber(value string) error {
	if value == EmptyString {
		return errors.New("Customer mobile number is required")
	}
	if len(value) > CustomerMobileNumberMaxLength {
		return errors.New(fmt.Sprintf("Customer mobile number must be less than %d characters", CustomerMobileNumberMaxLength))
	}
	q.AdditionalData[CustomerMobileNumber] = value
	return nil
}

// SetStoreLabel sets the store label in the additional data of the QRCode.
func (q *QRCode) SetStoreLabel(value string) error {
	if value == EmptyString {
		return errors.New("Store label is required")
	}
	if len(value) > StoreLabelMaxLength {
		return errors.New(fmt.Sprintf("Store label must be less than %d characters", StoreLabelMaxLength))
	}
	q.AdditionalData[StoreLabel] = value
	return nil
}

// SetReferenceLabel sets the reference label in the additional data of the QRCode.
func (q *QRCode) SetReferenceLabel(value string) error {
	if value == EmptyString {
		return errors.New("Reference label is required")
	}
	q.AdditionalData[ReferenceLabel] = value
	return nil
}

// SetCustomerLabel sets the customer label in the additional data of the QRCode.
func (q *QRCode) SetCustomerLabel(value string) error {
	if value == EmptyString {
		return errors.New("Customer label is required")
	}
	q.AdditionalData[CustomerLabel] = value
	return nil
}

// SetTerminalLabel sets the terminal label in the additional data of the QRCode.
func (q *QRCode) SetTerminalLabel(value string) error {
	if value == EmptyString {
		return errors.New("Terminal label is required")
	}
	if len(value) > TerminalLabelMaxLength {
		return errors.New(fmt.Sprintf("Terminal label must be less than %d characters", TerminalLabelMaxLength))
	}
	q.AdditionalData[TerminalLabel] = value
	return nil
}

// SetPurposeOfTransaction sets the purpose of transaction in the additional data of the QRCode.
func (q *QRCode) SetPurposeOfTransaction(value string) error {
	if value == EmptyString {
		return errors.New("Purpose of transaction is required")
	}
	q.AdditionalData[PurposeOfTransaction] = value
	return nil
}

// SetAdditionalConsumerDataRequest sets the additional consumer data request in the additional data of the QRCode.
func (q *QRCode) SetAdditionalConsumerDataRequest(value string) error {
	if value == EmptyString {
		return errors.New("Additional consumer data request is required")
	}
	q.AdditionalData[AdditionalConsumerDataRequest] = value
	return nil
}

// SetTransferType sets the transfer type for the QRCode.
func (q *QRCode) SetTransferType(transferType string) error {
	if !validTransferTypes[transferType] {
		return errors.New("Invalid transfer type. Use TransferToAccount or TransferToCard")
	}
	q.TransferType = transferType
	return nil
}

func (q *QRCode) SetMerchantId(merchantId string, isCompany bool) error {
	if merchantId == EmptyString {
		return errors.New("Merchant ID is required")
	}

	if q.QRType == EmptyString {
		return errors.New("QR type is required")
	}

	q.MerchantID = merchantId
	switch q.QRType {
	case QRTypeVNPay:
		if isCompany {
			q.MerchantIdentify = VNPAYMerchantCompanyService
		} else {
			q.MerchantIdentify = VNPAYMerchantPersonalService
		}
	}
	return nil
}

func (q *QRCode) SetMerchantName(merchantName string) error {
	if merchantName == EmptyString {
		return errors.New("Merchant name is required")
	}
	q.MerchantName = merchantName
	return nil
}

func (q *QRCode) SetMCC(mcc string) error {
	if mcc == EmptyString {
		return errors.New("MCC is required")
	}
	q.MCC = mcc
	return nil
}

func (q *QRCode) generateCommonQRData() map[string]interface{} {

	fields := map[string]interface{}{
		FieldQRVersion:  QRFormatVersion,
		FieldQRQRType:   StaticQR,
		FieldQRCurrency: q.TransactionCurrency,
	}

	if q.CountryCode != EmptyString {
		fields[FieldQRCountry] = q.CountryCode
	}

	if q.CityName != EmptyString {
		fields[FieldQRCityName] = q.CityName
	}

	if q.PostalCode != EmptyString {
		fields[FieldQRPostalCode] = q.PostalCode
	}

	if q.MCC != EmptyString {
		fields[FieldQRMCC] = q.MCC
	}

	if q.MerchantName != EmptyString {
		fields[FieldQRMerchantName] = q.MerchantName
	}

	if q.Amount != nil {
		fields[FieldQRAmount] = fmt.Sprintf("%d", *q.Amount)
		fields[FieldQRQRType] = DynamicQR
	}
	if len(q.AdditionalData) > 0 {
		subFields := map[string]string{}
		for key, value := range q.AdditionalData {
			subFields[key] = value
		}
		fields[FieldQRAdditionalData] = subFields
	}
	return fields
}

func (q *QRCode) IsValid() error {
	switch q.QRType {
	case QRTypeVietQR:
		return q.isValidVieQr()
	case QRTypeVNPay:
		return q.isValidVNPay()
	case QRTypeVNPayWallet:
		return q.isValidVNPayWallet()
	case QRTypeMomo:
		return q.isValidMomoWallet()
	}
	return errors.New("Invalid QR type. Use Predefined QR Code constants.")
}

// GenerateText generates the QR code text based on the QRCode data.
func (q *QRCode) GenerateText() (string, error) {
	switch q.QRType {
	case QRTypeVietQR:
		return q.generateVietQRText()
	case QRTypeVNPay:
		return q.generateVNPayQRText()
	case QRTypeVNPayWallet:
		return q.generateVNPayWalletQRText()
	case QRTypeMomo:
		return q.generateMoMoWalletQRText()
	}
	return "", errors.New("Invalid QR type. Use Predefined QR Code constants.")
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
