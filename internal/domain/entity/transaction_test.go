package entity

import (
	"testing"
	"time"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		testDescription string
		transaction     Transaction
		err             error
	}{
		{
			"testing userID = 0",
			Transaction{
				ID:          1,
				Value:       12.6,
				Description: "teste 1",
				Category:    "mercado",
				Date:        time.Now().UTC(),
				UserID:      0,
			},
			errInvalidUserID,
		},
		{
			"testing userID = 2",
			Transaction{
				ID:          1,
				Value:       12.6,
				Description: "teste 1",
				Category:    "mercado",
				Date:        time.Now().UTC(),
				UserID:      2,
			},
			errInvalidUserID,
		},
	}

	for _, test := range testCases {
		t.Run()
	}
}
