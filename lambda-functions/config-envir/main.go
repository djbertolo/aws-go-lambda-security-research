package main

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "Hello"
	}

	name := req.QueryStringParameters["name"]
	if name == "" {
		name = "World"
	}

	body, _ := json.Marshal(Response{
		Message: greeting + ", " + name + "!",
	})

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
