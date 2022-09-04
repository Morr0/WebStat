package controllers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func CreateResponse(statusCode int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	stringBody, _ := json.Marshal(body)
	return &events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: statusCode,
		Body:       string(stringBody),
	}, nil
}
