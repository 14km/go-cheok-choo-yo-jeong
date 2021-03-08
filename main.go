package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"

	"github.com/K-Connor/go-cheok-choo-yo-jeong/controllers/slack"
)

type MessageRequestEvent struct {
	Text string `json:"text"`
}

type MessageResponse struct {
	Message string `json:"Answer:"`
}

func main() {
	lambda.Start(HandleLambdaEvent)
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

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	return os.Getenv(key)
}
