package sofortpay

import (
	"fmt"
	"strings"
)

// Convert initialize payment structure to payload
func prepareInitializePaymentPayload(p *InitializePayment) (*initializePaymentPayload, error) {
	if err := p.Valid(); err != nil {
		return nil, err
	}

	pp := new(initializePaymentPayload)
	pp.Purpose = fmt.Sprintf("Order ID: %s", p.referenceTransactionID)
	pp.CurrencyID = strings.ToUpper(p.currencyID) // a currency id of "eur" will fail and only upper case "EUR" is valid
	pp.Amount = p.amount
	pp.Language = p.language
	pp.Metadata = p.metadata

	if p.successURL != nil {
		pp.SuccessURL = p.successURL.String()
	}

	if p.abortURL != nil {
		pp.AbortURL = p.abortURL.String()
	}

	if p.webhookURL != nil {
		pp.WebhookURL = p.webhookURL.String()
	}

	return pp, nil
}
