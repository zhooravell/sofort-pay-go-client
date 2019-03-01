package sofortpay

import (
	"strings"
	"testing"
)

func TestAPIKey_Valid(t *testing.T) {
	ak := APIKey("40f2de22-5288-416c-a008-7b241e5be67f")

	if err := ak.Valid(); err != nil {
		t.Error(err)
	}
}

func TestAPIKey_Empty(t *testing.T) {
	ak := APIKey("")

	err := ak.Valid()

	if err == nil {
		t.Errorf("should bee error")
	}

	if err.Error() != "api key should not be blank" {
		t.Errorf("invalid error message")
	}
}

func TestAPIKey_InvalidUUID(t *testing.T) {
	ak := APIKey("test")

	err := ak.Valid()

	if err == nil {
		t.Errorf("should bee error")
	}

	if !strings.Contains(err.Error(), "api key should be valid UUID") {
		t.Errorf("invalid error message: %s", err.Error())
	}
}
