package didar

import (
	"context"
	"encoding/json"
	"net/http"
)

func SaveContact(context context.Context, request *ContactRequest) (*ContactResponse, error) {
	req, err := makeRequest(context, "contact/save", http.MethodPost, request)
	if err != nil {
		return nil, err
	}
	content, err := callApi(req)
	if err != nil {
		return nil, err
	}

	response := new(ContactResponse)
	err = json.Unmarshal(content, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
