package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/integer00/telegrambotz/pkg/telegram"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	telegram.HandleTelegramHookString(request.Body)
	log.Printf("%+v", request.Body)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil

}
func main() {

	lambda.Start(handleRequest)

}
