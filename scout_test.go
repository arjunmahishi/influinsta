package main

import (
	"fmt"
	"testing"
)

func Test_scoutClient_ScoutVideos(t *testing.T) {
	testScout := scoutClient{
		hashtag: "instrumental",
	}

	testScout.ScoutVideos()

	if len(testScout.videoCollection) < 1 {
		t.Fail()
	}

	fmt.Println(len(testScout.videoCollection))
}
