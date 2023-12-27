package sqs

import (
	"accumulation_producer/internal/domain"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSHandler struct {
	SQSClient *sqs.SQS
	QueueURL  string
}

func NewSQSHandler(queueURL string) *SQSHandler {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Cambia esto a tu región de AWS deseada
	})
	if err != nil {
		log.Fatal("Error creando la sesión:", err)
	}

	sqsClient := sqs.New(sess)

	return &SQSHandler{
		SQSClient: sqsClient,
		QueueURL:  queueURL,
	}
}

func (h *SQSHandler) CreatePoint(point *domain.Point) error {
	pointJSON, err := json.Marshal(point)
	if err != nil {
		return fmt.Errorf("fallo al convertir el punto a JSON: %w", err)
	}

	sendMessageInput := &sqs.SendMessageInput{
		MessageBody:  aws.String(string(pointJSON)),
		QueueUrl:     aws.String(h.QueueURL),
		DelaySeconds: aws.Int64(0),
	}

	_, err = h.SQSClient.SendMessage(sendMessageInput)
	if err != nil {
		return fmt.Errorf("fallo al enviar el mensaje a SQS: %w", err)
	}

	return nil
}
