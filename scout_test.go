package main

import (
	"fmt"
	"testing"
)

func Test_scoutClient_ScoutVideos(t *testing.T) {
	testScout := NewScout("instrumental")

	testScout.ScoutVideos()

	best, err := testScout.GetBestVideo()
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Printf(best.User.Username)
}
