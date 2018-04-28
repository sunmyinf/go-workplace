package decode

type Certificate struct {
	CertificatePem string `json:"certificate_pem"`
	CertHashSha256 string `json:"cert_hash_sha256"`
}
