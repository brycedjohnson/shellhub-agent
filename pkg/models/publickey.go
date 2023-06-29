package models

type PublicKeyAuthRequest struct {
	Fingerprint string `json:"fingerprint"`
	Data        string `json:"data"`
}

type PublicKeyAuthResponse struct {
	Signature string `json:"signature"`
}
