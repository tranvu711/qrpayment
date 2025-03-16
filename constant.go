package qrpayment

type Bank struct {
	BIN       string
	Name      string
	ShortName string
	Code      string
	SwiftCode string
}

type QRCode struct {
	AdditionalData      map[string]string
	CountryCode         string
	TransactionCurrency string
	BankBin             string
	AccountNo           string
	CardNo              string
	Amount              *int
	Content             string
	TransferType        string
}

const (
	VietQRService = "A000000727"
)
const EmptyString = ""
const (
	BillNumber                    = "01"
	CustomerMobileNumber          = "02"
	StoreLabel                    = "03"
	ReferenceLabel                = "05"
	CustomerLabel                 = "06"
	TerminalLabel                 = "07"
	PurposeOfTransaction          = "08"
	AdditionalConsumerDataRequest = "09"
	CountryVN                     = "VN"
	CountryJP                     = "JP"
	CountryKR                     = "KR"
	CountryMY                     = "MY"
	CountryRC                     = "RC"
	CountryRI                     = "RI"
	CountryRP                     = "RP"
	CountrySG                     = "SG"
	CountryTH                     = "TH"
	CurrencyVND                   = "704"
	CurrencyJPY                   = "392"
	CurrencyKRW                   = "410"
	CurrencyMYR                   = "458"
	CurrencyCNY                   = "156"
	CurrencyIDR                   = "360"
	CurrencyPHP                   = "608"
	CurrencySGD                   = "702"
	CurrencyTHB                   = "764"
	PayloadFormatIndicator        = "000201"
	TransactionCurrency           = "53"
	CountryCode                   = "5802VN"
	CRC16Tag                      = "6304"
	StaticQR                      = "11"
	DynamicQR                     = "12"
	TransferToAccount             = "TransferToAccount" // Chuyển nhanh Napas247 đến tài khoản
	TransferToCard                = "TransferToCard"    // Chuyển nhanh Napas247 đến thẻ
)
