package main

type PushMessage struct {
	to      string
	message string
}

func main() {
	gcmApp := GcmPushApp{"API_KEY"}
	msg := PushMessage{"/topics/global", "hello world!"}

	gcmApp.send(msg)
}

