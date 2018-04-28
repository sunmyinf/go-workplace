package webhook

import "testing"

func TestVerifySignature(t *testing.T) {
	tests := []struct {
		sig     string
		secret  string
		payload []byte
	}{
		{
			sig:     "sha1=cfd0346a9aa97984082638fbd2c4d31ffcb0762a",
			secret:  "secret",
			payload: []byte("hogehogehoge"),
		},
	}

	for _, tt := range tests {
		if err := verifySignature(tt.sig, tt.secret, tt.payload); err != nil {
			t.Error(err)
		}
	}
}
