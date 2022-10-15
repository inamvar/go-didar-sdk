package didar

import (
	"context"
	"encoding/json"
	"net/http"
)

func SaveDeal(context context.Context, request *DealRequest) (*DealResponse, error) {
	req, err := makeRequest(context, "deal/save", http.MethodPost, request)
	if err != nil {
		return nil, err
	}
	content, err := callApi(req)
	if err != nil {
		return nil, err
	}

	response := new(DealResponse)
	err = json.Unmarshal(content, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
