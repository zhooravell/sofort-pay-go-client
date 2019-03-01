package sofortpay

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Extract error from response body
func prepareError(r *http.Response) error {
	defer r.Body.Close()

	var data errorMessage
	if err := json.NewDecoder(r.Body).Decode(&data); err == nil {
		return errors.New(fmt.Sprintf("sofort pay client: %s", data.Message))
	} else {
		return err
	}
}
