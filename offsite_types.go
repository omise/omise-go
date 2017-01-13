package omise

// OffsiteTypes represents an enumeration of possible types of offsite charges, which can
// be one of the following list of constants:
type OffsiteTypes string

const (
	InternetBankingSCB OffsiteTypes = "internet_banking_scb"
	InternetBankingBBL OffsiteTypes = "internet_banking_bbl"
	InternetBankingKTB OffsiteTypes = "internet_banking_ktb"
	InternetBankingBAY OffsiteTypes = "internet_banking_bay"
)
