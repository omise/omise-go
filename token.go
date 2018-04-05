package omise

// Token represents Omise's token object.
// See https://www.omise.co/tokens-api for more information.
//
// !!! IMPORTANT !!! - Full Credit Card data should never go through your server.
//
// Sending card data from server requires a valid PCI-DSS certification. You can learn
// more about this on the Security Best Practices page at
// https://www.omise.co/security-best-practices
type Token struct {
	Base
	Used bool  `json:"used"`
	Card *Card `json:"card"`
}
