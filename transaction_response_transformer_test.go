package sofortpay

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPrepareTransaction(t *testing.T) {
	j := `{
  "amount": 10.5,
  "currency_id": "EUR",
  "purpose": "Order ID: 1234",
  "metadata": {
    "key1": "value1",
    "key2": "value2",
    "key3": "value3"
  },
  "language": "de",
  "sender": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE"
  },
  "success_url": "https://example.com/success",
  "abort_url": "https://example.com/abort",
  "webhook_url": "https://example.com/webhook",
  "payform_code": "default",
  "recipient": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE",
    "street": "string",
    "city": "string",
    "zip": "string"
  },
  "uuid": "d871afc2-c227-4c8a-a28a-eaa97dbbd254",
  "status": "PENDING",
  "testmode": false
}`

	b := ioutil.NopCloser(bytes.NewReader([]byte(j)))
	defer b.Close()

	r := &http.Response{
		StatusCode: 200,
		Body:       b,
	}

	transaction, err := prepareTransaction(r)

	if err != nil {
		t.Fail()
	}

	if transaction.Status() != "PENDING" {
		t.Fail()
	}

	if transaction.UUID().String() != "d871afc2-c227-4c8a-a28a-eaa97dbbd254" {
		t.Fail()
	}

	if transaction.Amount() != 10.5 {
		t.Fail()
	}

	if transaction.SuccessURL().String() != "https://example.com/success" {
		t.Fail()
	}
}

func TestPrepareTransactionInvalidBody(t *testing.T) {
	j := `true`

	b := ioutil.NopCloser(bytes.NewReader([]byte(j)))
	defer b.Close()

	r := &http.Response{
		StatusCode: 200,
		Body:       b,
	}

	transaction, err := prepareTransaction(r)

	if err == nil {
		t.Fail()
	}

	if transaction != nil {
		t.Fail()
	}
}

func TestPrepareTransactionInvalidUUID(t *testing.T) {
	j := `{
  "amount": 10.5,
  "currency_id": "EUR",
  "purpose": "Order ID: 1234",
  "metadata": {
    "key1": "value1",
    "key2": "value2",
    "key3": "value3"
  },
  "language": "de",
  "sender": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE"
  },
  "success_url": "https://example.com/success",
  "abort_url": "https://example.com/abort",
  "webhook_url": "https://example.com/webhook",
  "payform_code": "default",
  "recipient": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE",
    "street": "string",
    "city": "string",
    "zip": "string"
  },
  "uuid": "123",
  "status": "PENDING",
  "testmode": false
}`

	b := ioutil.NopCloser(bytes.NewReader([]byte(j)))
	defer b.Close()

	r := &http.Response{
		StatusCode: 200,
		Body:       b,
	}

	transaction, err := prepareTransaction(r)

	if err == nil {
		t.Fail()
	}

	if transaction != nil {
		t.Fail()
	}
}

func TestPrepareTransactionInvalidSuccessUrl(t *testing.T) {
	j := `{
  "amount": 10.5,
  "currency_id": "EUR",
  "purpose": "Order ID: 1234",
  "metadata": {
    "key1": "value1",
    "key2": "value2",
    "key3": "value3"
  },
  "language": "de",
  "sender": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE"
  },
  "success_url": "[192:168:26:2::3]:1900",
  "abort_url": "https://example.com/abort",
  "webhook_url": "https://example.com/webhook",
  "payform_code": "default",
  "recipient": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE",
    "street": "string",
    "city": "string",
    "zip": "string"
  },
  "uuid": "d871afc2-c227-4c8a-a28a-eaa97dbbd254",
  "status": "PENDING",
  "testmode": false
}`

	b := ioutil.NopCloser(bytes.NewReader([]byte(j)))
	defer b.Close()

	r := &http.Response{
		StatusCode: 200,
		Body:       b,
	}

	transaction, err := prepareTransaction(r)

	if err == nil {
		t.Fail()
	}

	if transaction != nil {
		t.Fail()
	}
}

func TestPrepareTransactionInvalidAbortUrl(t *testing.T) {
	j := `{
  "amount": 10.5,
  "currency_id": "EUR",
  "purpose": "Order ID: 1234",
  "metadata": {
    "key1": "value1",
    "key2": "value2",
    "key3": "value3"
  },
  "language": "de",
  "sender": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE"
  },
  "success_url": "https://example.com",
  "abort_url": "[192:168:26:2::3]:1900",
  "webhook_url": "https://example.com/webhook",
  "payform_code": "default",
  "recipient": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE",
    "street": "string",
    "city": "string",
    "zip": "string"
  },
  "uuid": "d871afc2-c227-4c8a-a28a-eaa97dbbd254",
  "status": "PENDING",
  "testmode": false
}`

	b := ioutil.NopCloser(bytes.NewReader([]byte(j)))
	defer b.Close()

	r := &http.Response{
		StatusCode: 200,
		Body:       b,
	}

	transaction, err := prepareTransaction(r)

	if err == nil {
		t.Fail()
	}

	if transaction != nil {
		t.Fail()
	}
}

func TestPrepareTransactionInvalidWebhookUrl(t *testing.T) {
	j := `{
  "amount": 10.5,
  "currency_id": "EUR",
  "purpose": "Order ID: 1234",
  "metadata": {
    "key1": "value1",
    "key2": "value2",
    "key3": "value3"
  },
  "language": "de",
  "sender": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE"
  },
  "success_url": "https://example.com",
  "abort_url": "https://example.com",
  "webhook_url": "[192:168:26:2::3]:1900",
  "payform_code": "default",
  "recipient": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE",
    "street": "string",
    "city": "string",
    "zip": "string"
  },
  "uuid": "d871afc2-c227-4c8a-a28a-eaa97dbbd254",
  "status": "PENDING",
  "testmode": false
}`

	b := ioutil.NopCloser(bytes.NewReader([]byte(j)))
	defer b.Close()

	r := &http.Response{
		StatusCode: 200,
		Body:       b,
	}

	transaction, err := prepareTransaction(r)

	if err == nil {
		t.Fail()
	}

	if transaction != nil {
		t.Fail()
	}
}

func BenchmarkPrepareTransaction(b *testing.B) {
	j := `{
  "amount": 10.5,
  "currency_id": "EUR",
  "purpose": "Order ID: 1234",
  "metadata": {
    "key1": "value1",
    "key2": "value2",
    "key3": "value3"
  },
  "language": "de",
  "sender": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE"
  },
  "success_url": "https://example.com/success",
  "abort_url": "https://example.com/abort",
  "webhook_url": "https://example.com/webhook",
  "payform_code": "default",
  "recipient": {
    "holder": "John Doe",
    "iban": "DE04888888880087654321",
    "bic": "TESTDE88XXX",
    "bank_name": "Test Bank",
    "country_id": "DE",
    "street": "string",
    "city": "string",
    "zip": "string"
  },
  "uuid": "123",
  "status": "PENDING",
  "testmode": false
}`

	for n := 0; n < b.N; n++ {
		prepareTransaction(&http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(j))),
		})
	}
}
