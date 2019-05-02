package main

import (
	"encoding/json"
	"io/ioutil"
)

var (
	//Liked posts count
	Liked = 0
	// Followed count
	Followed = 0
)

// Config holds all the config values
var Config struct {
	Creds struct {
		Username string
		Password string
	}
	Hashtags        []string
	LikeThreshold   uint16
	FollowThreshold uint16
}

func init() {
	conts, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic("Can't read the config file")
	}

	json.Unmarshal(conts, &Config)
}
