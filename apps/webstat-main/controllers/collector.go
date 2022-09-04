package controllers

import (
	"github.com/aws/aws-lambda-go/events"
)

type CreateStatResponse struct {
	Message string `json:"message"`
}

func CollectStat(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return CreateResponse(200, CreateStatResponse{
		Message: "Created",
	})
}
