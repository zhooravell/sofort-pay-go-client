package sofortpay

import "testing"

func TestSender(t *testing.T) {
	s := Sender{
		holder:    "John Doe",
		iban:      "DE04888888880087654321",
		bic:       "TESTDE88XXX",
		bankName:  "Test Bank",
		countryID: "DE",
	}

	if s.Holder() != s.holder {
		t.Fail()
	}

	if s.IBAN() != s.iban {
		t.Fail()
	}

	if s.BIC() != s.bic {
		t.Fail()
	}

	if s.BankName() != s.bankName {
		t.Fail()
	}

	if s.CountryID() != s.countryID {
		t.Fail()
	}
}
