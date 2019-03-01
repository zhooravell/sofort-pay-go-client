package sofortpay

import (
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type APIKey string

// check is api key string valid
func (rcv APIKey) Valid() error {
	s := strings.TrimSpace(string(rcv))

	if s == "" {
		return errors.New("api key should not be blank")
	}

	if _, err := uuid.Parse(s); err != nil {
		return errors.Wrap(err, "api key should be valid UUID")
	}

	return nil
}
