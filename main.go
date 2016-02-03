package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PushMessage struct {
	to      string
	message string
}

type Config struct {
	GcmApiKey string `json:"gcmApiKey"`
}

func main() {
	config := loadConfig()
	gcmApp := GcmPushApp{config.GcmApiKey}
	msg := PushMessage{"/topics/global", "hello world!"}

	fmt.Printf("Config %v\n", config)

	gcmApp.send(msg)
}

func loadConfig() Config {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	config := Config{}
	json.Unmarshal(file, &config)
	return config
}
