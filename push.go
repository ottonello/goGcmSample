package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type PushApp struct {
	Config Config
	Gcm    GcmPushApp
}

type PushType int

const (
	GCM PushType = iota
)

type PushMessage struct {
	to        string
	message   string
	mechanism PushType
}

type Config struct {
	GcmApiKey string `json:"gcmApiKey"`
}

func NewPushApp() PushApp {
	config := loadConfig()
	gcmApp := NewGcmPushApp(config.GcmApiKey)
	return PushApp{config, gcmApp}
}

func (app *PushApp) SendPush(message PushMessage) {
	switch message.mechanism {
	case GCM:
		app.Gcm.Send(message)
	}
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
