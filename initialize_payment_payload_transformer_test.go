package sofortpay

import (
	"net/url"
	"testing"
)

func TestPrepareInitializePaymentPayload(t *testing.T) {
	p := NewInitializePayment("eur", 10.50, "123")

	webhookURL, _ := url.Parse("https://example.com/webhook")

	p.SetWebhookURL(webhookURL)
	p.SetLanguage("de")

	pp, err := prepareInitializePaymentPayload(p)

	if err != nil {
		t.Error(err)
	}

	if pp == nil {
		t.Fail()
	}

	if pp.CurrencyID != "EUR" {
		t.Fail()
	}

	if pp.Amount != p.amount {
		t.Fail()
	}

	if len(pp.Metadata) != len(p.metadata) {
		t.Fail()
	}

	if pp.SuccessURL != "" {
		t.Fail()
	}

	if pp.AbortURL != "" {
		t.Fail()
	}

	if pp.WebhookURL != "https://example.com/webhook" {
		t.Fail()
	}

	if pp.Language != p.language {
		t.Fail()
	}
}

func TestPrepareInitializePaymentPayloadFail(t *testing.T) {
	p := NewInitializePayment("eur", 10.50, "")

	pp, err := prepareInitializePaymentPayload(p)

	if err == nil {
		t.Fail()
	}

	if pp != nil {
		t.Fail()
	}
}
