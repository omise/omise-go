package omise

// SourceType represents an enumeration of possible values for Source.
type SourceType string

// SourceType can be one of the following list of constants:
const (
	SourceAlipay SourceType = "alipay"
	SourceBarcodeAlipay SourceType = "barcode_alipay"
	SourceBarcodeWechat SourceType = "barcode_wechat"
	SourceBillPaymentTescoLotus SourceType = "bill_payment_tesco_lotus"
	SourceEcontext SourceType = "econtext"
	SourceInstallmentBay SourceType = "installment_bay"
	SourceInstallmentBbl SourceType = "installment_bbl"
	SourceInstallmentFirstChoice SourceType = "installment_first_choice"
	SourceInstallmentKbank SourceType = "installment_kbank"
	SourceInstallmentKtc SourceType = "installment_ktc"
	SourceInternetBankingBay SourceType = "internet_banking_bay"
	SourceInternetBankingBbl SourceType = "internet_banking_bbl"
	SourceInternetBankingKtb SourceType = "internet_banking_ktb"
	SourceInternetBankingScb SourceType = "internet_banking_scb"
	SourcePaynow SourceType = "paynow"
	SourcePointsCiti SourceType = "points_citi"
	SourcePromptpay SourceType = "promptpay"
	SourceTruemoney SourceType = "truemoney"
)

