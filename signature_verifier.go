package sofortpay

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// To ensure the authenticity and integrity of the payload of a webhook, we sign the payload with a shared secret.
// The signature is transmitted in a header with our request and is named X-Payload-Signature. The header has the following structure:
// X-Payload-Signature: v1=<v1_signature_hash>;
type SignatureVerifier interface {
	Verify(xPayloadSignature string, webHookPayload []byte) error
}

var versionToAlgorithmsMap = map[string]string{
	"v1": "sha256",
}

type signatureVerifier struct {
	sharedSecret string
}

// Verify signature using shared secret
func (rcv *signatureVerifier) Verify(xPayloadSignature string, webHookPayload []byte) error {
	for _, versionedSignature := range strings.Split(xPayloadSignature, ";") {
		var version, signature string

		idx := strings.IndexByte(versionedSignature, '=')

		if idx == -1 {
			return errors.New("sofort pay: invalid signature (separator not found)")
		}

		version = versionedSignature[:idx]
		adjustedPos := idx + 1

		if adjustedPos >= len(versionedSignature) {
			return errors.New("sofort pay: invalid signature")
		}

		signature = versionedSignature[adjustedPos:]
		sum, err := rcv.hashSumByVersion(version, webHookPayload)

		if err != nil {
			return err
		}

		if fmt.Sprintf("%x", sum) == signature {
			return nil
		}
	}

	return errors.New("sofort pay: invalid signature")
}

func (rcv *signatureVerifier) hashSumByVersion(version string, webHookPayload []byte) ([]byte, error) {
	switch version {
	case "v1":
		h := hmac.New(sha256.New, []byte(rcv.sharedSecret))
		h.Write(webHookPayload)
		return h.Sum(nil), nil
	}

	return nil, errors.New("sofort pay: hashing error")
}
