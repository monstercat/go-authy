package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode int
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
}

func NewResponse(response *http.Response) (*Response, error) {
	request := &Response{StatusCode: response.StatusCode}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Logger.Println("Error reading from API:", err)
		return request, err
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return request, err
	}
	return request, nil
}

// Valid returns true if the SMS was sent
func (r *Response) Valid() bool {
	return r.StatusCode == 200
}
