package didar

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var API_KEY = "<your api key>"
var BASE_URL = "https://app.didar.me/api"

func makeRequest(path string, method string, data any) (*http.Request, error) {

	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	str := string(jsonBody)
	fmt.Println(str)
	bodyReader := bytes.NewReader(jsonBody)

	url := fmt.Sprintf("%s/%s?apikey=%s", BASE_URL, path, API_KEY)

	req, err := http.NewRequest(method, url, bodyReader)

	if err != nil {
		return nil, err
	}

	if API_KEY == "" {
		return nil, errors.New("no api key provided")
	}

	header := make(http.Header)
	header["Content-Type"] = []string{"application/json"}
	//	header["Accept"] = []string{"application/json"}
	//	header["api-key"] = []string{API_KEY}
	req.Header = header
	return req, nil
}

func callApi(req *http.Request) ([]byte, error) {
	resposne, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resposne.Body.Close()
	content, _ := ioutil.ReadAll(resposne.Body)

	fmt.Println(string(content))

	err = checkStatus(resposne)
	return content, err
}

func checkStatus(resposne *http.Response) error {
	var err error
	if resposne != nil {
		switch resposne.StatusCode {
		case http.StatusOK, http.StatusAccepted, http.StatusCreated:
			err = nil
		default:
			err = errors.New(http.StatusText(resposne.StatusCode))
		}
	}

	return err
}
