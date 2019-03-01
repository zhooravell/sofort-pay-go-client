package sofortpay

// Object holding information about a sender account
type Sender struct {
	holder    string
	iban      string
	bic       string
	bankName  string
	countryID string
}

// Name of the account holder
func (rcv *Sender) Holder() string {
	return rcv.holder
}

// International Bank Account Number (IBAN) of the account
func (rcv *Sender) IBAN() string {
	return rcv.iban
}

// Business Identifier Code (BIC) of the account
func (rcv *Sender) BIC() string {
	return rcv.bic
}

// Name of the accounts bank
func (rcv *Sender) BankName() string {
	return rcv.bankName
}

// 2 letter country code of the account
func (rcv *Sender) CountryID() string {
	return rcv.countryID
}
