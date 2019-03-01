package sofortpay

// Recipient
type Recipient struct {
	holder    string
	iban      string
	bic       string
	bankName  string
	countryID string
	street    string
	city      string
	zip       string
}

// Name of the account holder
func (rcv *Recipient) Holder() string {
	return rcv.holder
}

// International Bank Account Number (IBAN) of the account (ISO 13616-1	)
func (rcv *Recipient) IBAN() string {
	return rcv.iban
}

// Business Identifier Code (BIC) of the account (ISO 9362)
func (rcv *Recipient) BIC() string {
	return rcv.bic
}

// Name of the accounts bank
func (rcv *Recipient) BankName() string {
	return rcv.bankName
}

// 2 letter country code of the account (ISO 3166-1 alpha-2	)
func (rcv *Recipient) CountryID() string {
	return rcv.countryID
}

// Street name
func (rcv *Recipient) Street() string {
	return rcv.street
}

// City name
func (rcv *Recipient) City() string {
	return rcv.city
}

// Zone Improvement Plan used in postal addresses
func (rcv *Recipient) Zip() string {
	return rcv.zip
}
