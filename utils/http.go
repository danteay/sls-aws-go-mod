package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// Response wraps the base API Gateway response
type Response events.APIGatewayProxyResponse

// NewResponse -
func NewResponse(code int, body interface{}) (*Response, error) {
	j, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res := &Response{
		StatusCode:      code,
		Body:            string(j),
		IsBase64Encoded: false,
	}

	return res, nil
}

// SetHeaders define response headers
func (r *Response) SetHeaders(h map[string]string) {
	r.Headers = h
}

// WithCors add cors support to API Gatway response
func (r *Response) WithCors() {
	if r.Headers == nil {
		r.Headers = map[string]string{}
	}

	r.Headers["Access-Control-Allow-Origin"] = "*"
	r.Headers["Access-Control-Allow-Credentials"] = "true"
}
