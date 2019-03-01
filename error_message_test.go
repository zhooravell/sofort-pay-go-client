package sofortpay

import (
	"encoding/json"
	"testing"
)

func TestErrorMessageUnmarshal(t *testing.T) {
	j := `{"code":400,"message":"Bad Request"}`

	var e errorMessage

	if err := json.Unmarshal([]byte(j), &e); err != nil {
		t.Fail()
	}

	if e.Message != "Bad Request" {
		t.Fail()
	}

	if e.Code != 400 {
		t.Fail()
	}
}
