package omise

// QrCode represents resource object.
type QrCode struct {
	Base
	Email string `json:"email"`
	MfaProvisioningURI string `json:"mfa_provisioning_uri"`
	Secret string `json:"secret"`
}

// QrCodeService represents resource service.
type QrCodeService struct {
	*Client
}

// QrCode defines resource service.
func (c *Client) QrCode() *QrCodeService {
	return &QrCodeService{c}
}

