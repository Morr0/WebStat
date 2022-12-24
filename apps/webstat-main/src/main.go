package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/morr0/webstat/apps/webstat-main/src/controllers"
)

func main() {
	println("Started app")
	lambda.Start(handler)
}

func handler(req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	awsSession, awsError := session.NewSession()
	if awsError != nil {
		log.Fatal("Failed to create an AWS session")

		headers := make(map[string]string)
		return &events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       "",
		}, nil
	}

	switch req.RawPath {
	case "/v1/stat":
		return controllers.HandleStat(req, *awsSession)

	default:
		var resObj struct{}
		return controllers.CreateResponse(http.StatusMethodNotAllowed, resObj)
	}
}
