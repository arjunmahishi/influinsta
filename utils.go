package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ahmdrz/goinsta/v2"
)

func downloadImage(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func sessionExists(sessionPath string) bool {
	info, err := os.Stat(sessionPath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getItemScore(item goinsta.Item) int {
	return (item.Likes +
		item.CommentCount +
		int(item.DeviceTimestamp) +
		int(item.ViewCount) +
		len(item.Hashtags()) +
		len(item.Mentions)) / 6
}

func alreadyPosted(item goinsta.Item) bool {
	myPosts := GetInstagram().MyPosts(5)
	for _, post := range myPosts {
		if strings.Contains(post.Caption.Text, item.Caption.Text) {
			return true
		}
	}
	return false
}

func selectBestItem(items []goinsta.Item) (*goinsta.Item, error) {
	bestScore := -1
	var bestItem goinsta.Item

	if len(items) == 0 {
		return nil, fmt.Errorf("the video collection is empty")
	}

	for _, item := range items {
		score := getItemScore(item)
		if score > bestScore && !alreadyPosted(item) {
			bestItem = item
			bestScore = score
		}
	}
	return &bestItem, nil
}

func getRandomGenericComment() string {
	rand.Seed(time.Now().Unix())
	return genericComments[rand.Intn(len(genericComments))]
}

func likeItems(items []goinsta.Item) {
	tempLiked := 0

	for _, item := range items {
		if Liked < int(Config.LikeThreshold) {
			if err := item.Like(); err == nil {
				Liked++
				tempLiked++
			}
		}
	}
	log.Printf("liked %d posts", tempLiked)
}
