package didar_test

import (
	"context"
	"go-arvanvod-sdk/go-didar-sdk"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateContact(t *testing.T) {

	didar.API_KEY = "<API_KEY>"
	request := &didar.ContactRequest{
		Contact: didar.Contact{
			FirstName:   "John",
			LastName:    "Smith",
			Email:       "test@example.com",
			MobilePhone: "0912345678",
			//Fields:      []didar.ContactField{{Key: "school_id", Value: 1}, {Key: "school_title", Value: "اکادمی داریک"}},
			//SegmentIds:  []string{"696d0306-22c1-4d38-a311-7bfa8a68d149"},
		},
	}

	resp, err := didar.SaveContact(context.Background(), request)

	if assert.NoError(t, err) {
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Response.Id)
	}
}
