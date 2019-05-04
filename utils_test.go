package main

import (
	"fmt"
	"testing"
)

func Test_getBestItem(t *testing.T) {
	items := GetInstagram().SearchHashtagForVideos("instrumental")
	best, err := getBestItem(items)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(best.User.Username)
}
