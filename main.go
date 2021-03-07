package main

import (
	"fmt"
	"log"
	"os"

	"github.com/K-Connor/go-cheok-choo-yo-jeong/controllers/slack"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type MessageRequestEvent struct {
	Text string `json:"text"`
}

type MessageResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MessageRequestEvent) (MessageResponse, error) {
	slack.SetSlackUrl(goDotEnvVariable("SLACK_URL"))

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("CheokChooYoJeong Error : ", err)
		}
	}()

	body, err := slack.SendCheokChooYoJeong()

	if err != nil {
		panic(err)
	}

	return MessageResponse{Message: fmt.Sprint(string(body))}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
