package sofortpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

const sofortPayBaseUrl = "https://api.sofort-pay.com"

// Please note that sofortpay is a Payment Initiation Service,
// which means that your client can initiate a payment with their bank account via this payment method.
//
// https://manage.sofort-pay.com/merchant/documentation/payments/api/payment
type Client interface {
	InitializePayment(p *InitializePayment) (*Transaction, error)
	GetPayment(uuid uuid.UUID) (*Transaction, error)
	DeletePayment(uuid uuid.UUID) error
}

// SofortPay HTTP Client constructor (public)
func NewSofortPayClient(httpClient *http.Client, apiKey APIKey) (Client, error) {
	baseUrl, err := url.Parse(sofortPayBaseUrl)

	if err != nil {
		return nil, err
	}

	if err := apiKey.Valid(); err != nil {
		return nil, err
	}

	return newClient(httpClient, baseUrl, apiKey), nil
}

// SofortPay HTTP Client
type client struct {
	baseURL    *url.URL
	apiKey     APIKey
	httpClient *http.Client
}

// SofortPay HTTP Client constructor
func newClient(httpClient *http.Client, baseUrl *url.URL, apiKey APIKey) *client {
	return &client{
		baseURL:    baseUrl,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// Initialize a payment
func (rcv *client) InitializePayment(p *InitializePayment) (*Transaction, error) {
	payload, err := prepareInitializePaymentPayload(p)

	if err != nil {
		return nil, err
	}

	r, err := rcv.newRequest("POST", "/api/v1/payments", payload)

	if err != nil {
		return nil, err
	}

	res, err := rcv.httpClient.Do(r)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusAccepted {
		return nil, prepareError(res)
	}

	return prepareTransaction(res)
}

// Return Payment Transaction Object
func (rcv *client) GetPayment(uuid uuid.UUID) (*Transaction, error) {
	r, err := rcv.newRequest("GET", fmt.Sprintf("/api/v1/payments/%s", uuid), nil)

	if err != nil {
		return nil, err
	}

	res, err := rcv.httpClient.Do(r)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, prepareError(res)
	}

	return prepareTransaction(res)
}

// Delete a Payment Transaction Object
func (rcv *client) DeletePayment(uuid uuid.UUID) error {
	r, err := rcv.newRequest("DELETE", fmt.Sprintf("/api/v1/payments/%s", uuid), nil)

	if err != nil {
		return err
	}

	res, err := rcv.httpClient.Do(r)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return prepareError(res)
	}

	return nil
}

// Prepare request
func (rcv *client) newRequest(method string, path string, body interface{}) (*http.Request, error) {
	u := rcv.baseURL.ResolveReference(&url.URL{Path: path})
	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	r, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	if body != nil {
		r.Header.Set("Content-Type", "application/json")
	}

	r.Header.Set("Accept", "application/json")
	r.Header.Set("Authorization", string(rcv.apiKey))

	return r, nil
}
