package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/morr0/webstat/apps/webstat-main/src/controllers"
)

func main() {
	println("Started app")
	lambda.Start(handler)
}

func handler(req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	switch req.RawPath {
	case "/v1/stat":
		return controllers.HandleStat(req)

	default:
		headers := make(map[string]string)
		return &events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusOK,
			Headers:    headers,
			Body:       "",
		}, nil
	}
}
