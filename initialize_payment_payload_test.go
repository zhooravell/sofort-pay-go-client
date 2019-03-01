package sofortpay

import (
	"encoding/json"
	"testing"
)

func TestInitializePaymentPayloadJsonEncode(t *testing.T) {
	payload := initializePaymentPayload{
		Purpose:    "test_1",
		CurrencyID: "EUR",
		Amount:     33.55,
		Metadata: map[string]interface{}{
			"transaction_id": "test_2",
		},
	}

	j, err := json.Marshal(payload)

	if err != nil {
		t.Error(err)
	}

	expected := `{"purpose":"test_1","currency_id":"EUR","amount":33.55,"metadata":{"transaction_id":"test_2"}}`
	actual := string(j)

	if expected != actual {
		t.Error("bad format")
	}
}

func TestInitializePaymentPayloadJsonEncodeWithOmitempty(t *testing.T) {
	payload := initializePaymentPayload{
		Purpose:    "test_1",
		CurrencyID: "EUR",
		Amount:     33.55,
		Metadata: map[string]interface{}{
			"transaction_id": "test_2",
		},
		Language:   "EN",
		WebhookURL: "https://google.com/1",
		SuccessURL: "https://google.com/2",
		AbortURL:   "https://google.com/3",
	}

	j, err := json.Marshal(payload)

	if err != nil {
		t.Error(err)
	}

	expected := `{"purpose":"test_1","currency_id":"EUR","amount":33.55,"metadata":{"transaction_id":"test_2"},"language":"EN","success_url":"https://google.com/2","abort_url":"https://google.com/3","webhook_url":"https://google.com/1"}`
	actual := string(j)

	if expected != actual {
		t.Error("bad format")
	}
}
