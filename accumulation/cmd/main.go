package main

import (
	"accumulation_producer/internal/app"
	"accumulation_producer/internal/domain"
	"accumulation_producer/internal/infrastructure/sqs"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaHandler struct {
	pointHandler *app.PointHandler
}

func NewLambdaHandler() *LambdaHandler {
	queueURL := os.Getenv("QUEUE_URL")
	if queueURL == "" {
		log.Fatal("La variable de entorno QUEUE_URL es requerida")
	}

	sqsHandler := sqs.NewSQSHandler(queueURL)
	pointApp := app.NewPointApplication(sqsHandler)
	pointHandler := app.NewPointHandler(pointApp)
	return &LambdaHandler{
		pointHandler,
	}
}

func (h *LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	b := []byte(request.Body)
	var body domain.Point
	err := json.Unmarshal(b, &body)
	if err != nil {
		log.Fatal("Error Unmarshal:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	err = h.pointHandler.HandlePointCreation(&body)
	if err != nil {
		log.Fatal("Error HandlePointCreation:", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: fmt.Sprintf("%v", domain.Response{
			Code:    "200",
			Message: "point: success create",
		}),
	}, nil
}

func main() {
	handler := NewLambdaHandler()
	lambda.Start(handler.HandleRequest)
}
