package sofortpay

import "testing"

func TestRecipient(t *testing.T) {
	r := Recipient{
		holder:    "MUSTERMANN, HARTMUT",
		iban:      "AT04888888880087654321",
		bic:       "TESTAT88XXX",
		bankName:  "Testbank Austria",
		countryID: "AT",
		street:    "street",
		city:      "city",
		zip:       "zip",
	}

	if r.Holder() != r.holder {
		t.Fail()
	}

	if r.IBAN() != r.iban {
		t.Fail()
	}

	if r.BIC() != r.bic {
		t.Fail()
	}

	if r.BankName() != r.bankName {
		t.Fail()
	}

	if r.CountryID() != r.countryID {
		t.Fail()
	}

	if r.Street() != r.street {
		t.Fail()
	}

	if r.City() != r.city {
		t.Fail()
	}

	if r.Zip() != r.zip {
		t.Fail()
	}
}
