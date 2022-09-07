package controllers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type CreateStatResponse struct {
	Message string `json:"message"`
}

func collectStat(req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	println("Body:", req.Body)
	return CreateResponse(http.StatusOK, CreateStatResponse{
		Message: "Created",
	})
}

func HandleStat(req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	switch req.RequestContext.HTTP.Method {
	case http.MethodPost:
		return collectStat(req)
	default:
		return CreateResponse(http.StatusMethodNotAllowed, CreateStatResponse{
			Message: "Not Allowed",
		})
	}
}
