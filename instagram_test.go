package main

import (
	"fmt"
	"testing"
)

func Test_instaClient_SearchHashtagForImages(t *testing.T) {
	items := GetInstagram().SearchHashtagForImages("instrumental")
	fmt.Println(len(items))
}

func Test_instaClient_SearchHashtagForVideos(t *testing.T) {
	items := GetInstagram().SearchHashtagForVideos("instrumental")
	fmt.Println(len(items))
}
