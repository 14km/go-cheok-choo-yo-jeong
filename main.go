package main

import (
	"fmt"
	"log"
	_ "os"

	_ "github.com/K-Connor/go-cheok-choo-yo-jeong/controllers/slack"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//os.Getenv("")
	//
	//slack.NewSlack(event.text)
	//
	//slack.SendMessage()

	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
