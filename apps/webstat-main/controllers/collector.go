package controllers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type CreateStatResponse struct {
	Message string `json:"message"`
}

func CollectStat(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	println("Body:", req.Body)
	return CreateResponse(http.StatusOK, CreateStatResponse{
		Message: "Created",
	})
}
