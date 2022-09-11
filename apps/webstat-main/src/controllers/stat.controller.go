package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/morr0/webstat/apps/webstat-main/src/handlers"
)

type CreateStatResponse struct {
	Message string `json:"message"`
}

func collectStat(req events.APIGatewayV2HTTPRequest, awsSession session.Session) (*events.APIGatewayV2HTTPResponse, error) {
	var request handlers.CollectStatRequest
	parseErr := json.Unmarshal([]byte(req.Body), &request)
	if parseErr != nil {
		return CreateResponse(http.StatusBadRequest, CreateStatResponse{
			Message: "{}",
		})
	}

	statusCode := handlers.CollectStat(request, awsSession)

	return CreateResponse(statusCode, CreateStatResponse{
		Message: "{}",
	})
}

func HandleStat(req events.APIGatewayV2HTTPRequest, awsSession session.Session) (*events.APIGatewayV2HTTPResponse, error) {
	switch req.RequestContext.HTTP.Method {
	case http.MethodPost:
		return collectStat(req, awsSession)
	default:
		return CreateResponse(http.StatusMethodNotAllowed, CreateStatResponse{
			Message: "Not Allowed",
		})
	}
}
