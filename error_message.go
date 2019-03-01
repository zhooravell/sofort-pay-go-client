package sofortpay

type errorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
