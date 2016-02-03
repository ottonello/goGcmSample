package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 2 {
		fmt.Println("Supply arguments:\ngoGcmSample [to] [message]")
		return
	}
	to := argsWithoutProg[0] // /topics/global
	text := argsWithoutProg[1]

	msg := PushMessage{to, text, GCM}

	pushApp := NewPushApp()
	pushApp.SendPush(msg)
}
