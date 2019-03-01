package sofortpay

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestTransaction(t *testing.T) {
	var m = map[string]interface{}{}
	m["test"] = "test"

	id := uuid.New()

	successURL, _ := url.Parse("https://google.com/1")
	abortURL, _ := url.Parse("https://google.com/2")
	webhookURL, _ := url.Parse("https://google.com/3")

	obj := Transaction{
		amount:      0.21,
		currencyID:  "EUR",
		purpose:     "order: id",
		metadata:    m,
		language:    "en",
		sender:      &Sender{},
		successURL:  successURL,
		abortURL:    abortURL,
		webhookURL:  webhookURL,
		payFormCode: "payFormCode",
		recipient:   &Recipient{},
		uuid:        &id,
		status:      "PENDING",
		isTestMode:  true,
	}

	if obj.Amount() != obj.amount {
		t.Fail()
	}

	if obj.CurrencyID() != obj.currencyID {
		t.Fail()
	}

	if obj.Purpose() != obj.purpose {
		t.Fail()
	}

	if !reflect.DeepEqual(obj.Metadata(), obj.metadata) {
		t.Fail()
	}

	if obj.Language() != obj.language {
		t.Fail()
	}

	if obj.Sender() != obj.sender {
		t.Fail()
	}

	if obj.SuccessURL() != obj.successURL {
		t.Fail()
	}

	if obj.AbortURL() != obj.abortURL {
		t.Fail()
	}

	if obj.WebhookURL() != obj.webhookURL {
		t.Fail()
	}

	if obj.PayFormCode() != obj.payFormCode {
		t.Fail()
	}

	if obj.Recipient() != obj.recipient {
		t.Fail()
	}

	if obj.UUID() != obj.uuid {
		t.Fail()
	}

	if obj.Status() != obj.status {
		t.Fail()
	}

	if obj.IsTestMode() != obj.isTestMode {
		t.Fail()
	}
}
