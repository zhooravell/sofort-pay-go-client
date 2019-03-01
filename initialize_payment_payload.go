package sofortpay

type initializePaymentPayload struct {
	Purpose    string                 `json:"purpose"`
	CurrencyID string                 `json:"currency_id"`
	Amount     float64                `json:"amount"`
	Metadata   map[string]interface{} `json:"metadata"`
	Language   string                 `json:"language,omitempty"`
	SuccessURL string                 `json:"success_url,omitempty"`
	AbortURL   string                 `json:"abort_url,omitempty"`
	WebhookURL string                 `json:"webhook_url,omitempty"`
}
