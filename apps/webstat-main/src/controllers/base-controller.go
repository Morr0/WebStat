package controllers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func CreateResponse(statusCode int, body interface{}) (*events.APIGatewayV2HTTPResponse, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Access-Control-Allow-Origin"] = "*"

	stringBody, _ := json.Marshal(body)
	return &events.APIGatewayV2HTTPResponse{
		Headers:    headers,
		StatusCode: statusCode,
		Body:       string(stringBody),
	}, nil
}
