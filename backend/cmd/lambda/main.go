package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jamesmoessis/dust_sensor/backend/handlers"
	"github.com/jamesmoessis/dust_sensor/backend/storage"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	agnosticReq := &handlers.Request{
		Path:   request.Path,
		Body:   request.Body,
		Method: request.HTTPMethod,
	}
	b, _ := json.Marshal(request)
	fmt.Printf(string(b) + "\n")
	ddb := storage.NewDynamoSettingsDb(ctx)
	handler := handlers.Handler{
		DB: ddb,
	}
	res, err := handler.RouterHandler(ctx, agnosticReq)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	return events.APIGatewayProxyResponse{Body: res.Body, StatusCode: res.Status}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
