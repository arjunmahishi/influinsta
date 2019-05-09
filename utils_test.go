package main

import (
	"fmt"
	"testing"
)

func Test_selectBestItem(t *testing.T) {
	items := GetInstagram().SearchHashtagForVideos("instrumental")
	best, err := selectBestItem(items)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(best.User.Username)
	fmt.Println(len(best.Hashtags()))
}
