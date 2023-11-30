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
	return events.APIGatewayProxyResponse{
		Body:       res.Body,
		StatusCode: res.Status,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "http://dust.jamesmoessis.com",
			"Access-Control-Allow-Methods": "GET, PUT",
			"Access-Control-Allow-Headers": "Content-Type",
			"Content-Type":                 "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
