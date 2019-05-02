package main

import (
	"fmt"

	"gopkg.in/ahmdrz/goinsta.v2"
)

// Scout interface
type Scout interface {
	ScoutImage()
	ScoutVoice()
	LikeCollectedImages()
	LikeCollectedVideos()
	GetBestImage() (*goinsta.Item, error)
	GetBestVideo() (*goinsta.Item, error)
}

// scoutClient responsible for performing actions on instagram
type scoutClient struct {
	hashtag         string
	imageCollection []goinsta.Item
	videoCollection []goinsta.Item
}

// ScoutImages for the hashtag
func (s *scoutClient) ScoutImages() {
	s.imageCollection = GetInstagram().SearchHashtagForImages(s.hashtag)
}

// ScoutVideos for the hashtag
func (s *scoutClient) ScoutVideos() {
	s.imageCollection = GetInstagram().SearchHashtagForVideos(s.hashtag)
}

// LikeCollectedImages likes all the images the scout has collected
func (s *scoutClient) LikeCollectedImages() {
	for _, item := range s.imageCollection {
		item.Like()
	}
}

// LikeCollectedVideos likes all the videos the scout has collected
func (s *scoutClient) LikeCollectedVideos() {
	for _, item := range s.videoCollection {
		item.Like()
	}
}

// GetBestImage from the collection
func (s *scoutClient) GetBestImage() (*goinsta.Item, error) {
	mostLikes := -1
	var bestItem goinsta.Item

	if len(s.imageCollection) == 0 {
		return nil, fmt.Errorf("the image collection is empty")
	}

	for _, item := range s.imageCollection {
		if item.Likes > mostLikes {
			bestItem = item
			mostLikes = item.Likes
		}
	}
	return &bestItem, nil
}

// GetBestVideo from the collection
func (s *scoutClient) GetBestVideo() (*goinsta.Item, error) {
	mostLikes := -1
	var bestItem goinsta.Item

	if len(s.videoCollection) == 0 {
		return nil, fmt.Errorf("the video collection is empty")
	}

	for _, item := range s.videoCollection {
		if item.Likes > mostLikes {
			bestItem = item
			mostLikes = item.Likes
		}
	}
	return &bestItem, nil
}