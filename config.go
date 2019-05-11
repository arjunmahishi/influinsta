package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var (
	//Liked posts count
	Liked = 0
	// Followed count
	Followed = 0

	configPath = flag.String("config", "./config.json", "path of the config file")
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
	Actions         []struct {
		Name string
		Args []string
	}
}

func init() {
	flag.Parse()
	conts, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic("Can't read the config file")
	}

	json.Unmarshal(conts, &Config)
}
