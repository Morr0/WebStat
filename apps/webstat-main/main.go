package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/morr0/webstat/apps/webstat-main/controllers"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		return controllers.CollectStat(req)
	default:
		headers := make(map[string]string)
		return &events.APIGatewayProxyResponse{
			StatusCode: 405,
			Headers:    headers,
			Body:       "",
		}, nil
	}
}
