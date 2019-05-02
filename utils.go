package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/ahmdrz/goinsta.v2"
)

func downloadImage(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func getBestItem(items []goinsta.Item) (*goinsta.Item, error) {
	mostLikes := -1
	var bestItem goinsta.Item

	if len(items) == 0 {
		return nil, fmt.Errorf("the video collection is empty")
	}

	for _, item := range items {
		if item.Likes > mostLikes {
			bestItem = item
			mostLikes = item.Likes
		}
	}
	return &bestItem, nil
}

func sessionExists(sessionPath string) bool {
	info, err := os.Stat(sessionPath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
