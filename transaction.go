package sofortpay

import (
	"net/url"

	"github.com/google/uuid"
)

type Transaction struct {
	amount      float64
	currencyID  string
	purpose     string
	metadata    map[string]interface{}
	language    string
	sender      *Sender
	successURL  *url.URL
	abortURL    *url.URL
	webhookURL  *url.URL
	payFormCode string
	recipient   *Recipient
	uuid        *uuid.UUID
	status      string
	isTestMode  bool
}

// Is test mode
func (rcv *Transaction) IsTestMode() bool {
	return rcv.isTestMode
}

// Transaction status
func (rcv *Transaction) Status() string {
	return rcv.status
}

// Transaction ID
func (rcv *Transaction) UUID() *uuid.UUID {
	return rcv.uuid
}

// Recipient
func (rcv *Transaction) Recipient() *Recipient {
	return rcv.recipient
}

// Code to specify which customized Payform template should be used.
func (rcv *Transaction) PayFormCode() string {
	return rcv.payFormCode
}

// URL to notify an endpoint about transaction events. Please use only HTTPS for production mode.
func (rcv *Transaction) WebhookURL() *url.URL {
	return rcv.webhookURL
}

// Return URL if a transaction was aborted
func (rcv *Transaction) AbortURL() *url.URL {
	return rcv.abortURL
}

// Return URL if a transaction was successful
func (rcv *Transaction) SuccessURL() *url.URL {
	return rcv.successURL
}

// Object holding information about a sender account
func (rcv *Transaction) Sender() *Sender {
	return rcv.sender
}

// The language our payment form initially will be shown to the user (ISO 639-1)
func (rcv *Transaction) Language() string {
	return rcv.language
}

// An object of data which will be passed back to your application
func (rcv *Transaction) Metadata() map[string]interface{} {
	return rcv.metadata
}

// The purpose (aka subject) of your payment
func (rcv *Transaction) Purpose() string {
	return rcv.purpose
}

// The currency of your payment (ISO 4217)
func (rcv *Transaction) CurrencyID() string {
	return rcv.currencyID
}

// The amount of your payment
func (rcv *Transaction) Amount() float64 {
	return rcv.amount
}
