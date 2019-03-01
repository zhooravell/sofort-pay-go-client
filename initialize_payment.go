package sofortpay

import (
	"net/url"

	"github.com/pkg/errors"
)

type InitializePayment struct {
	currencyID             string
	amount                 float64
	referenceTransactionID string
	language               string
	metadata               map[string]interface{}
	successURL             *url.URL
	abortURL               *url.URL
	webhookURL             *url.URL
}

func NewInitializePayment(currencyID string, amount float64, referenceTransactionID string) *InitializePayment {
	return &InitializePayment{
		currencyID:             currencyID,
		amount:                 amount,
		referenceTransactionID: referenceTransactionID,
		metadata:               map[string]interface{}{},
	}
}

// Set the language our payment form initially will be shown to the user
func (rcv *InitializePayment) SetLanguage(language string) *InitializePayment {
	rcv.language = language

	return rcv
}

// Set return URL if a transaction was successful
func (rcv *InitializePayment) SetSuccessURL(successURL *url.URL) *InitializePayment {
	rcv.successURL = successURL

	return rcv
}

// Set return URL if a transaction was aborted
func (rcv *InitializePayment) SetAbortURL(abortURL *url.URL) *InitializePayment {
	rcv.abortURL = abortURL

	return rcv
}

// Set URL to notify an endpoint about transaction events. Please use only HTTPS for production mode.
func (rcv *InitializePayment) SetWebhookURL(webhookURL *url.URL) *InitializePayment {
	rcv.webhookURL = webhookURL

	return rcv
}

// An object of data which will be passed back to your application
func (rcv *InitializePayment) AddMeta(key string, value interface{}) *InitializePayment {
	rcv.metadata[key] = value

	return rcv
}

func (rcv *InitializePayment) Valid() error {
	if len(rcv.currencyID) != 3 {
		return errors.New("The currency of your payment should bee from ISO 4217")
	}

	if rcv.language != "" && len(rcv.language) != 2 {
		return errors.New("The language of your payment should bee from ISO 639-1")
	}

	if rcv.referenceTransactionID == "" {
		return errors.New("The reference transaction ID should not be blank")
	}

	if len(rcv.metadata) > 3 {
		return errors.New("Metadata should contain 3 elements or less")
	}

	return nil
}
