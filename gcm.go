package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
)

type GcmPushApp struct {
	apiKey string
}

type GcmPush struct {
	To   string `json:"to"`
	Data GcmData `json:"data"`
}

type GcmData struct {
	Message string `json:"message"`
}

func (app GcmPushApp) send(pushMessage PushMessage) {
	specificMessage := GcmPush{pushMessage.to, GcmData{Message: pushMessage.message}}
	app.sendGcm(specificMessage)
}

func (app GcmPushApp) sendGcm(gcmPush GcmPush) {
	client := &http.Client{}

	b, err := json.Marshal(gcmPush)

	checkError(err)

	req, err := http.NewRequest("POST", "https://android.googleapis.com/gcm/send", bytes.NewBuffer(b))

	checkError(err)

	req.Header.Add("Authorization", "key=" + app.apiKey)
	req.Header.Add("Content-Type", "application/json");

	fmt.Println("Sending request ", req)

	resp, err := client.Do(req);

	checkError(err)

	fmt.Println("Success! ", resp)

}

func checkError(err error) {
	if (err != nil) {
		fmt.Println("Error", err)
		panic(fmt.Sprintf("%s", err))
	}
}