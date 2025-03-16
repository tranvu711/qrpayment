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
	CityName            string
	PostalCode          string
	TransactionCurrency string
	BankBin             string
	AccountNo           string
	CardNo              string
	Amount              *int
	Content             string
	TransferType        string
	QRType              string

	// VNPay data
	MerchantID       string
	MerchantIdentify string
	MerchantName     string
	TerminalID       string
	TerminalName     string
	MCC              string
}

const (
	VietQRService                = "A000000727"
	VNPAYMerchantCompanyService  = "A000000775"
	VNPAYMerchantPersonalService = "970436"
	VNPAYWALLETService           = "908401"
)

const (
	QRFormatVersion = "01"
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
	CRC16Tag                      = "6304"
	StaticQR                      = "11"
	DynamicQR                     = "12"
	VNPayWalletQR                 = "10"
	TransferToAccount             = "TransferToAccount"
	TransferToCard                = "TransferToCard"
)

const (
	QRTypeVietQR      = "VIETQR"
	QRTypeVNPay       = "VNPAY"
	QRTypeVNPayWallet = "VNPayWallet"
	QRTypeMomo        = "MOMO"
)

const (
	FieldQRVersion        = "00"
	FieldQRQRType         = "01"
	FieldQRMCC            = "52"
	FieldQRCurrency       = "53"
	FieldQRCountry        = "58"
	FieldQRMerchantName   = "59"
	FieldQRAmount         = "54"
	FieldQRAdditionalData = "62"
	FieldQRCityName       = "60"
	FieldQRPostalCode     = "61"

	// VietQR
	FieldVietQRIdentity                        = "38"
	FiledVietQRIdentitySubService              = "00"
	FiledVietQRIdentitySubCustomer             = "01"
	FiledVietQRIdentitySubCustomerBankBin      = "00"
	FiledVietQRIdentitySubCustomerBankInfo     = "01"
	FiledVietQRIdentitySubCustomerBankInfoType = "02"

	// VNPAY
	FieldVNPayIdentity              = "26"
	FieldVNPayIdentitySubService    = "00"
	FieldVNPayIdentitySubMerchantId = "01"
)

const (
	CityMaxLength                 = 15
	PostalCodeMaxLength           = 10
	StoreLabelMaxLength           = 25
	TerminalLabelMaxLength        = 25
	BillNumberMaxLength           = 25
	CustomerMobileNumberMaxLength = 25
)
