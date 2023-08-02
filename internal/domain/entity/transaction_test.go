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
			Transaction{
				ID:          1,
				Value:       12.6,
				Description: "teste 1",
				Category:    "mercado",
				Date:        time.Now().UTC(),
				UserID:      0,
			},
			ErrInvalidUserID,
		},
		{
			Transaction{
				ID:          1,
				Value:       12.6,
				Description: "teste 1",
				Category:    "mercado",
				Date:        time.Now().UTC(),
				UserID:      2,
			},
			ErrInvalidUserID,
		},
	}

	for _, test := range testCases {
		t.Run()
	}
}
