# QRPayment SDK

QRPayment SDK is a Go library that supports generating multiple types of payment QR codes, including VietQR, VNPay, VNPay Wallet, and MoMo Wallet.

## Installation

```sh
go get github.com/tranvu711/qrpayment
```

## Usage

### Initialize QRCode

```go
import "github.com/tranvu711/qrpayment"

qr, err := qrpayment.NewQRPayment()
if err != nil {
    panic(err)
}
```

## Generate QR Codes by Type

### VietQR Example

```go
qr, _ := qrpayment.NewQRPayment()
qr.SetQRType(qrpayment.QRTypeVietQR)
qr.SetBankBin(qrpayment.BANKS.Vietcombank.BIN)
qr.SetAccountNo("0441000670827")
qr.SetAmount(100000)
qr.SetNote("Payment for order 123")
qr.SetTransactionCurrency(qrpayment.CurrencyVND)
qr.SetCountryCode(qrpayment.CountryVN)

qrText, _ := qr.GenerateText()
fmt.Println("VietQR Code Text:", qrText)
```

### VNPay Example

```go
qr, _ := qrpayment.NewQRPayment()
qr.SetQRType(qrpayment.QRTypeVNPay)
qr.SetMCC("4829")
qr.SetMerchantName("My Shop")
qr.SetMerchantId("123456789", true)
qr.SetTerminalID("T12345")
qr.SetTerminalName("Main Terminal")
qr.SetCityName("Hanoi")
qr.SetPostalCode("100000")
qr.SetAmount(500000)

qrText, _ := qr.GenerateText()
fmt.Println("VNPay QR Code Text:", qrText)
```

### VNPay Wallet Example

```go
qr, _ := qrpayment.NewQRPayment()
qr.SetQRType(qrpayment.QRTypeVNPayWallet)
qr.SetCustomerMobileNumber("0901455890")
qr.SetAmount(200000)
qr.SetCountryCode(qrpayment.CountryVN)

qrText, _ := qr.GenerateText()
fmt.Println("VNPay Wallet QR Code Text:", qrText)
```

### MoMo Wallet Example - By Mobile Number

```go
qr, _ := qrpayment.NewQRPayment()
qr.SetQRType(qrpayment.QRTypeMomo)
qr.SetCustomerMobileNumber("0901455890")
qr.SetAmount(150000)
qr.SetNote("Momo payment")

qrText, _ := qr.GenerateText()
fmt.Println("MoMo Wallet QR Code Text:", qrText)
```

### MoMo Wallet Example - By Bank Account
- Currently, MoMo's multi-purpose QR codes are using BVBank (Bản Việt Bank) to receive payments.
- Each MoMo account is assigned a corresponding account number at BVBank. Any funds transferred to this account number will be forwarded to the MoMo wallet.
- You can find this account number on the details page of the payment QR code within the MoMo app.

```go
qr, _ := qrpayment.NewQRPayment()
qr.SetQRType(qrpayment.QRTypeVietQR)
qr.SetBankBin(qrpayment.BANKS.BVBank.BIN)
qr.SetAccountNo("99MM23307M35010091")
qr.SetAmount(100000)
qr.SetNote("Payment for order 123")
qr.SetTransactionCurrency(qrpayment.CurrencyVND)
qr.SetCountryCode(qrpayment.CountryVN)

qrText, _ := qr.GenerateText()
fmt.Println("VietQR Code Text:", qrText)
```

## Available Methods

- `NewQRPayment(configs ...*QRCode) (*QRCode, error)`: Initializes a new QRPayment instance.
- `SetQRType(qrType string) error`: Sets the QR type.
- `SetBankBin(bankBin string) error`: Sets the bank BIN.
- `SetAccountNo(accountNo string) error`: Sets the account number.
- `SetAmount(amount int) error`: Sets the payment amount.
- `SetNote(note string) error`: Sets the transaction note.
- `SetCardNo(cardNo string) error`: Sets the card number.
- `SetTransactionCurrency(currency string) error`: Sets the transaction currency.
- `SetCountryCode(country string) error`: Sets the country code.
- `SetMerchantId(merchantId string, isCompany bool) error`: Sets the merchant ID.
- `SetMerchantName(merchantName string) error`: Sets the merchant name.
- `SetMCC(mcc string) error`: Sets the Merchant Category Code.
- `GenerateText() (string, error)`: Generates the QR code text.
- `GeneratePNG(filename string) error`: Generates a QR code image and saves it to a file.

## Required Fields & Validation per QR Type

### VietQR
- **Required Fields:** `BankBin`, `AccountNo` or `CardNo`, `Amount`
- **Validation:**
  - `BankBin` and `AccountNo` are required for account transfers.
  - `BankBin` and `CardNo` are required for card transfers.
  - `Amount` must be greater than 0.

### VNPay
- **Required Fields:** `MCC`, `MerchantName`, `MerchantID`, `TerminalID`, `TerminalName`, `CityName`, `PostalCode`
- **Validation:**
  - All required fields must be provided and non-empty.
  - `Amount` must be greater than 0.

### VNPay Wallet
- **Required Fields:** `CustomerMobileNumber`
- **Validation:**
  - `CustomerMobileNumber` must be provided.
  - `Amount`, if provided, must be greater than 0.

### MoMo Wallet - By Mobile Number
- **Required Fields:** `CustomerMobileNumber`
- **Validation:**
  - `CustomerMobileNumber` must be provided.
  - `Amount`, if provided, must be greater than 0.

## Support

If you encounter any issues or have questions, please create an issue on the project's GitHub repository.

## Donate

If you find this project useful and would like to support its development, consider making a donation:

- **Website:** [tranvu.info](https://tranvu.info)
- **PayPal:** [paypal.me/movewho](https://paypal.me/movewho)
- **Buy Me a Coffee:** [buymeacoffee.com/tranvu711](https://buymeacoffee.com/tranvu711)

Your support is greatly appreciated! 😊

