package sofortpay

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/uuid"
)

var server *httptest.Server
var baseUrl *url.URL
var transactionIDToDelete string
var transactionIDToGet string
var apiKey string

func init() {
	transactionIDToDelete = "50583e0a-6698-45a3-9b8f-26e869122b1d"
	transactionIDToGet = "d871afc2-c227-4c8a-a28a-eaa97dbbd254"

	transaction := `{
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

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"code":400,"message":"Bad Request"}`))
			return
		}

		if r.Header.Get("Authorization") != apiKey {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"code":401,"message":"Unauthorized"}`))
			return
		}

		if r.Method == "DELETE" {
			if fmt.Sprintf("/api/v1/payments/%s", transactionIDToDelete) != r.URL.String() {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"code":404,"message":"Not Found"}`))
				return
			}

			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method == "GET" {
			if fmt.Sprintf("/api/v1/payments/%s", transactionIDToGet) != r.URL.String() {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"code":404,"message":"Not Found"}`))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(transaction))
			return
		}

		if r.Method == "POST" {
			if "/api/v1/payments" != r.URL.String() {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"code":404,"message":"Not Found"}`))
				return
			}

			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(transaction))
			return
		}
	}))

	baseUrl, _ = url.Parse(server.URL)
}

func TestClient_DeletePayment(t *testing.T) {
	client := newClient(server.Client(), baseUrl, APIKey(apiKey))
	transactionUUID, _ := uuid.Parse(transactionIDToDelete)

	if err := client.DeletePayment(transactionUUID); err != nil {
		t.Error(err)
	}
}

func TestClient_DeletePaymentUnauthorized(t *testing.T) {
	client := newClient(server.Client(), baseUrl, APIKey("d5c0c073-992d-4128-9c5c-491fd56cf74f"))
	transactionUUID, _ := uuid.Parse(transactionIDToDelete)

	err := client.DeletePayment(transactionUUID)

	if err == nil || err.Error() != "sofort pay client: Unauthorized" {
		t.Fail()
	}
}

func TestClient_DeletePaymentNotFound(t *testing.T) {
	client := newClient(server.Client(), baseUrl, APIKey(apiKey))
	transactionUUID, _ := uuid.Parse("e63b981b-c904-48a3-9cc4-1061d285ab32")

	err := client.DeletePayment(transactionUUID)

	if err == nil || err.Error() != "sofort pay client: Not Found" {
		t.Fail()
	}
}

func TestClient_GetPaymentNotFound(t *testing.T) {
	client := newClient(server.Client(), baseUrl, APIKey(apiKey))
	transactionUUID, _ := uuid.Parse("e63b981b-c904-48a3-9cc4-1061d285ab32")

	transaction, err := client.GetPayment(transactionUUID)

	if transaction != nil {
		t.Fail()
	}

	if err == nil || err.Error() != "sofort pay client: Not Found" {
		t.Fail()
	}
}

func TestClient_GetPaymentUnauthorized(t *testing.T) {
	client := newClient(server.Client(), baseUrl, APIKey("d5c0c073-992d-4128-9c5c-491fd56cf74f"))
	transactionUUID, _ := uuid.Parse("e63b981b-c904-48a3-9cc4-1061d285ab32")

	transaction, err := client.GetPayment(transactionUUID)

	if transaction != nil {
		t.Fail()
	}

	if err == nil || err.Error() != "sofort pay client: Unauthorized" {
		t.Fail()
	}
}

func TestClient_GetPayment(t *testing.T) {
	client := newClient(server.Client(), baseUrl, APIKey(apiKey))
	transactionUUID, _ := uuid.Parse(transactionIDToGet)

	transaction, err := client.GetPayment(transactionUUID)

	if err != nil {
		t.Fail()
	}

	if transaction == nil {
		t.Fail()
	}

	if transaction.UUID().String() != "d871afc2-c227-4c8a-a28a-eaa97dbbd254" {
		t.Fail()
	}

	if transaction.Status() != "PENDING" {
		t.Fail()
	}

	if transaction.Amount() != 10.5 {
		t.Fail()
	}
}

func TestClient_InitializePayment(t *testing.T) {
	client := newClient(server.Client(), baseUrl, APIKey(apiKey))

	p := NewInitializePayment("EUR", 10.50, "123")

	transaction, err := client.InitializePayment(p)

	if err != nil {
		t.Fail()
	}

	if transaction == nil {
		t.Fail()
	}

	if transaction.UUID().String() != "d871afc2-c227-4c8a-a28a-eaa97dbbd254" {
		t.Fail()
	}

	if transaction.Status() != "PENDING" {
		t.Fail()
	}

	if transaction.Amount() != 10.5 {
		t.Fail()
	}
}

func TestClient_InitializePaymentUnauthorized(t *testing.T) {
	client := newClient(server.Client(), baseUrl, APIKey("d5c0c073-992d-4128-9c5c-491fd56cf74f"))

	p := NewInitializePayment("EUR", 10.50, "123")

	transaction, err := client.InitializePayment(p)

	if transaction != nil {
		t.Fail()
	}

	if err == nil || err.Error() != "sofort pay client: Unauthorized" {
		t.Fail()
	}
}
