package sofortpay

import (
	"net/url"
	"testing"
)

func TestNewInitializePayment(t *testing.T) {
	p := NewInitializePayment("EUR", 10.50, "123")

	if err := p.Valid(); err != nil {
		t.Error(err)
	}

	if p.currencyID != "EUR" {
		t.Fail()
	}

	if p.amount != 10.50 {
		t.Fail()
	}

	if p.referenceTransactionID != "123" {
		t.Fail()
	}
}

func TestNewInitializePaymentInvalidCurrency(t *testing.T) {
	p := NewInitializePayment("EU", 10.50, "123")
	err := p.Valid()

	if err != nil && err.Error() != "The currency of your payment should bee from ISO 4217" {
		t.Error(err)
	}
}

func TestNewInitializePaymentInvalidReference(t *testing.T) {
	p := NewInitializePayment("EUR", 10.50, "")
	err := p.Valid()

	if err != nil && err.Error() != "The reference transaction ID should not be blank" {
		t.Error(err)
	}
}

func TestNewInitializePaymentInvalidMeta(t *testing.T) {
	p := NewInitializePayment("EUR", 10.50, "123")

	p.AddMeta("key_1", 1)
	p.AddMeta("key_2", "a")
	p.AddMeta("key_3", 2)
	p.AddMeta("key_4", "b")

	if p.metadata["key_1"] != 1 {
		t.Fail()
	}

	if p.metadata["key_2"] != "a" {
		t.Fail()
	}

	if p.metadata["key_3"] != 2 {
		t.Fail()
	}

	if p.metadata["key_4"] != "b" {
		t.Fail()
	}

	err := p.Valid()

	if err != nil && err.Error() != "Metadata should contain 3 elements or less" {
		t.Error(err)
	}
}

func TestNewInitializePaymentWithOptions(t *testing.T) {
	p := NewInitializePayment("EUR", 10.50, "123")
	p.SetLanguage("en")

	successURL, _ := url.Parse("https://example.com/success")
	abortURL, _ := url.Parse("https://example.com/abort")
	webhookURL, _ := url.Parse("https://example.com/webhook")

	p.SetSuccessURL(successURL)
	p.SetAbortURL(abortURL)
	p.SetWebhookURL(webhookURL)

	if err := p.Valid(); err != nil {
		t.Error(err)
	}

	if p.currencyID != "EUR" {
		t.Fail()
	}

	if p.amount != 10.50 {
		t.Fail()
	}

	if p.referenceTransactionID != "123" {
		t.Fail()
	}

	if p.language != "en" {
		t.Fail()
	}

	if p.successURL.String() != "https://example.com/success" {
		t.Fail()
	}

	if p.abortURL.String() != "https://example.com/abort" {
		t.Fail()
	}

	if p.webhookURL.String() != "https://example.com/webhook" {
		t.Fail()
	}
}

func TestNewInitializePaymentAddMeta(t *testing.T) {
	p := NewInitializePayment("EUR", 10.50, "123")
	p.AddMeta("key_1", 123)
	p.AddMeta("key_2", "test")

	if p.metadata["key_1"] != 123 {
		t.Fail()
	}

	if p.metadata["key_2"] != "test" {
		t.Fail()
	}
}
