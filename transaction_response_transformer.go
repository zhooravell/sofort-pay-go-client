package sofortpay

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

// Extract Transaction from response body
func prepareTransaction(r *http.Response) (*Transaction, error) {
	defer r.Body.Close()

	var data struct {
		Amount     float64                `json:"amount"`
		CurrencyID string                 `json:"currency_id"`
		Purpose    string                 `json:"purpose"`
		Metadata   map[string]interface{} `json:"metadata"`
		Language   string                 `json:"language"`
		Sender     struct {
			Holder    string `json:"holder"`
			Iban      string `json:"iban"`
			Bic       string `json:"bic"`
			BankName  string `json:"bank_name"`
			CountryID string `json:"country_id"`
		} `json:"sender"`
		SuccessURL  string `json:"success_url"`
		AbortURL    string `json:"abort_url"`
		WebhookURL  string `json:"webhook_url"`
		PayFormCode string `json:"payform_code"`
		Recipient   struct {
			Holder    string `json:"holder"`
			Iban      string `json:"iban"`
			Bic       string `json:"bic"`
			BankName  string `json:"bank_name"`
			CountryID string `json:"country_id"`
			Street    string `json:"street"`
			City      string `json:"city"`
			Zip       string `json:"zip"`
		} `json:"recipient"`
		UUID     string `json:"uuid"`
		Status   string `json:"status"`
		TestMode bool   `json:"testmode"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}

	uuiID, err := uuid.Parse(data.UUID)

	if err != nil {
		return nil, err
	}

	data.SuccessURL = strings.TrimSpace(data.SuccessURL)
	data.AbortURL = strings.TrimSpace(data.AbortURL)
	data.WebhookURL = strings.TrimSpace(data.WebhookURL)

	var successURL *url.URL
	var abortURL *url.URL
	var webhookURL *url.URL

	if data.SuccessURL != "" {
		successURL, err = url.Parse(data.SuccessURL)
		if err != nil {
			return nil, err
		}
	}

	if data.AbortURL != "" {
		abortURL, err = url.Parse(data.AbortURL)

		if err != nil {
			return nil, err
		}
	}

	if data.WebhookURL != "" {
		webhookURL, err = url.Parse(data.WebhookURL)

		if err != nil {
			return nil, err
		}
	}

	transaction := new(Transaction)
	transaction.uuid = &uuiID
	transaction.amount = data.Amount
	transaction.currencyID = data.CurrencyID
	transaction.purpose = data.Purpose
	transaction.metadata = data.Metadata
	transaction.language = data.Language
	transaction.payFormCode = data.PayFormCode
	transaction.status = data.Status
	transaction.isTestMode = data.TestMode
	transaction.successURL = successURL
	transaction.abortURL = abortURL
	transaction.webhookURL = webhookURL

	transaction.sender = new(Sender)
	transaction.sender.holder = data.Sender.Holder
	transaction.sender.iban = data.Sender.Iban
	transaction.sender.bic = data.Sender.Bic
	transaction.sender.bankName = data.Sender.BankName
	transaction.sender.countryID = data.Sender.CountryID

	transaction.recipient = new(Recipient)
	transaction.recipient.holder = data.Recipient.Holder
	transaction.recipient.iban = data.Recipient.Iban
	transaction.recipient.bic = data.Recipient.Bic
	transaction.recipient.bankName = data.Recipient.BankName
	transaction.recipient.countryID = data.Recipient.CountryID
	transaction.recipient.street = data.Recipient.Street
	transaction.recipient.city = data.Recipient.City
	transaction.recipient.zip = data.Recipient.Zip

	return transaction, nil
}
