# QRPayment SDK

QRPayment SDK is a Go library that supports generating payment QR codes following the standard of VietQR, VNPayQR,...

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

### Configure Payment Information

```go
qr.SetBankBin(qrpayment.BANKS.Vietcombank.BIN)
qr.SetAccountNo("0441000670827")
qr.SetAmount(100000)
qr.SetNote("Payment for order 123")
qr.SetTransactionCurrency(qrpayment.CurrencyVND)
qr.SetCountryCode(qrpayment.CountryVN)
```

### Generate QR Code Text

```go
qrText, err := qr.GenerateText()
if err != nil {
    panic(err)
}
fmt.Println("QR Code Text:", qrText)
```

### Generate QR Code Image

```go
err = qr.GeneratePNG("qrcode.png")
if err != nil {
    panic(err)
}
```

## Available Methods

- `NewQRPayment(configs ...*QRCode) (*QRCode, error)`: Initializes a new QRPayment instance.
- `SetBankBin(bankBin string)`: Sets the bank BIN.
- `SetAccountNo(accountNo string)`: Sets the account number.
- `SetAmount(amount int)`: Sets the payment amount.
- `SetNote(note string)`: Sets the transaction note.
- `SetCardNo(cardNo string)`: Sets the card number.
- `SetTransactionCurrency(currency string) error`: Sets the transaction currency.
- `SetCountryCode(country string) error`: Sets the country code.
- `SetBillNumber(value string)`: Sets the bill number.
- `SetCustomerMobileNumber(value string)`: Sets the customer's mobile number.
- `SetStoreLabel(value string)`: Sets the store label.
- `SetReferenceLabel(value string)`: Sets the reference label.
- `SetCustomerLabel(value string)`: Sets the customer label.
- `SetTerminalLabel(value string)`: Sets the terminal label.
- `SetPurposeOfTransaction(value string)`: Sets the purpose of transaction.
- `SetAdditionalConsumerDataRequest(value string)`: Sets additional consumer data.
- `SetTransferType(transferType string) error`: Sets the transfer type.
- `GenerateText() (string, error)`: Generates the QR code text.
- `GeneratePNG(filename string) error`: Generates a QR code image and saves it to a file.

## Support

If you encounter any issues or have questions, please create an issue on the project's GitHub repository.

