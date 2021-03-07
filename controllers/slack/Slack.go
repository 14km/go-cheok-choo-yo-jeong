package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type MessageRequest struct {
	Text string `json:"text"`
}

func NewSlack(text string) *MessageRequest {
	// 구조체 인스턴스를 생성한 뒤 포인터를 리턴
	return &MessageRequest{Text: text}
}

func SendCheokChooYoJeong() (err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("CheokChooYoJeong Error : ", err)
		}
	}()

	jsonBytes, _ := json.Marshal(MessageRequest{"여러분 척추를 피고 일을 하세요!"})
	buff := bytes.NewBuffer(jsonBytes)

	req, _ := http.NewRequest(
		http.MethodPost,
		"https://hooks.slack.com/services/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",
		buff,
	)

	//Content-Type 헤더 추가
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return err
}

func SendMessage(context echo.Context) (err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("SendMessage Error : ", err)
		}
	}()

	text := context.FormValue("text")
	Slack := NewSlack(text)

	jsonBytes, _ := json.Marshal(Slack)
	buff := bytes.NewBuffer(jsonBytes)

	req, _ := http.NewRequest(
		http.MethodPost,
		"https://hooks.slack.com/services/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",
		buff,
	)

	//Content-Type 헤더 추가
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return err
}
