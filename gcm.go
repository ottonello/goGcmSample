package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GcmPushApp an application supporting GCM push
type GcmPushApp struct {
	apiKey string
}

type gcmPush struct {
	To   string      `json:"to"`
	Data gcmPushData `json:"data"`
}

type gcmPushData struct {
	Message string `json:"message"`
}

func (app GcmPushApp) send(pushMessage PushMessage) {
	specificMessage := gcmPush{pushMessage.to, gcmPushData{Message: pushMessage.message}}
	app.sendGcm(specificMessage)
}

func (app GcmPushApp) sendGcm(gcmPush gcmPush) {
	b, _ := json.Marshal(gcmPush)

	if req, err := app.buildRequest(b); err == nil {
		sendRequest(req)
	} else {
		fmt.Println("Error building request ", err)
	}

}

func (app GcmPushApp) buildRequest(b []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", "https://android.googleapis.com/gcm/send", bytes.NewBuffer(b))
	req.Header.Add("Authorization", "key="+app.apiKey)
	req.Header.Add("Content-Type", "application/json")

	return req, err
}

func sendRequest(req *http.Request) {
	client := &http.Client{}

	resp, err := client.Do(req)

	if err == nil {
		fmt.Println("Success! ", resp)
	} else {
		fmt.Println("Error sending request", err)
	}

}
