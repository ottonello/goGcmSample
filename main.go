package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	fmt.Printf("Config %v\n", config)

	gcmApp := GcmPushApp{config.GcmApiKey}
	msg := PushMessage{"/topics/global", "hello world!"}

	gcmApp.send(msg)
}

func loadConfig() Config {
	file := readConfigFile("./config.json")
	return unmarshalConfig(file)
}

func readConfigFile(configFile string) []byte {
	file, e := ioutil.ReadFile(configFile)
	if e != nil {
		panic(fmt.Sprintf("Could not open config file %v\n", file))
	}
	return file
}

func unmarshalConfig(file []byte) Config {
	config := Config{}
	err := json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	return config
}
