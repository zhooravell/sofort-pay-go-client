package sofortpay

import "testing"

func TestSignatureVerifier_Verify(t *testing.T) {
	sv := signatureVerifier{
		sharedSecret: "V;*s5ii@316w=HmuW1fPC:35?Js$5UH$",
	}

	xPayloadSignature := "v1=b95ac0a6fbb0f868eb2e66fafbc79bc1e31995b0773b78de515c9c1dc5d7038a"

	payload := `{
    "uuid": "65452679-c7fe-4c63-84d9-1ea42cad8a1b",
    "purpose": "kKR8L Order ID 1234",
    "amount": 10.5,
    "currency_id": "EUR",
    "metadata": {
        "id": "bla-bla-bla"
    },
    "testmode": true,
    "status": "PENDING",
    "abort_url": "https:\/\/webhook.site\/67cd7da5-dfe5-47f2-81e7-497fe52998a7",
    "success_url": "https:\/\/google.com.ua",
    "webhook_url": "https:\/\/webhook.site\/67cd7da5-dfe5-47f2-81e7-497fe52998a7",
    "payform_code": null,
    "language": "ru",
    "recipient": {
        "holder": "Centrobill (Cyprus) Limited",
        "iban": "DE31512308000000060396",
        "bic": "WIREDEMMXXX",
        "bank_name": "Wirecard Bank",
        "country_id": "DE"
    },
    "sender": {
        "holder": "MUSTERMANN, HARTMUT",
        "iban": "DE62888888880012345678",
        "bic": "TESTDE88XXX",
        "bank_name": "Testbank",
        "country_id": "DE"
    }
}`

	if err := sv.Verify(xPayloadSignature, []byte(payload)); err != nil {
		t.Error(err)
	}
}

func TestSignatureVerifier_VerifyFail(t *testing.T) {
	sv := signatureVerifier{
		sharedSecret: "test",
	}

	xPayloadSignature := "v1=b95ac0a6fbb0f868eb2e66fafbc79bc1e31995b0773b78de515c9c1dc5d7038a"

	payload := `{
    "uuid": "65452679-c7fe-4c63-84d9-1ea42cad8a1b",
    "purpose": "kKR8L Order ID 1234",
    "amount": 10.5,
    "currency_id": "EUR",
    "metadata": {
        "id": "bla-bla-bla"
    },
    "testmode": true,
    "status": "PENDING",
    "abort_url": "https:\/\/webhook.site\/67cd7da5-dfe5-47f2-81e7-497fe52998a7",
    "success_url": "https:\/\/google.com.ua",
    "webhook_url": "https:\/\/webhook.site\/67cd7da5-dfe5-47f2-81e7-497fe52998a7",
    "payform_code": null,
    "language": "ru",
    "recipient": {
        "holder": "Centrobill (Cyprus) Limited",
        "iban": "DE31512308000000060396",
        "bic": "WIREDEMMXXX",
        "bank_name": "Wirecard Bank",
        "country_id": "DE"
    },
    "sender": {
        "holder": "MUSTERMANN, HARTMUT",
        "iban": "DE62888888880012345678",
        "bic": "TESTDE88XXX",
        "bank_name": "Testbank",
        "country_id": "DE"
    }
}`

	if err := sv.Verify(xPayloadSignature, []byte(payload)); err == nil {
		t.Fail()
	}
}
