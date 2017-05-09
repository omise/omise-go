package omise

// OffsiteTypes represents an enumeration of possible types of offsite charges.
type OffsiteTypes string

// OffsiteTypes can be one of the following list of constants:
const (
	InternetBankingSCB OffsiteTypes = "internet_banking_scb"
	InternetBankingBBL OffsiteTypes = "internet_banking_bbl"
	InternetBankingKTB OffsiteTypes = "internet_banking_ktb"
	InternetBankingBAY OffsiteTypes = "internet_banking_bay"
	Alipay             OffsiteTypes = "alipay"
)
