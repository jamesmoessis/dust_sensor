package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jamesmoessis/dust_sensor/backend/handlers"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lambdaAgnosticReq := &handlers.Request{
		Path:   request.Path,
		Body:   request.Body,
		Method: request.HTTPMethod,
	}
	res, err := handlers.RouterHandler(lambdaAgnosticReq)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	return events.APIGatewayProxyResponse{Body: res.Body, StatusCode: res.Status}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
