package slack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type MessageRequest struct {
	Text string `json:"text"`
}

// Slack Send Url
var slackUrl string

func SetSlackUrl(url string) {
	slackUrl = url
}

func SendCheokChooYoJeong() (body []byte, err error) {
	jsonBytes, _ := json.Marshal(MessageRequest{"여러분, 아프니까 척추입니다! 모두들 척추를 펴고 일을 하세요!"})
	buff := bytes.NewBuffer(jsonBytes)

	req, _ := http.NewRequest(
		http.MethodPost,
		slackUrl,
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

	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}
