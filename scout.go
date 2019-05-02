package main

import (
	"gopkg.in/ahmdrz/goinsta.v2"
)

// Scout is responsible for performing actions on instagram
type Scout struct {
	hashtag         string
	imageCollection []goinsta.Item
	videoCollection []goinsta.Item
}

// ScoutImages for the hashtag
func (s *Scout) ScoutImages() {
	s.imageCollection = GetInstagram().SearchHashtagForImages(s.hashtag)
}

// ScoutVideos for the hashtag
func (s *Scout) ScoutVideos() {
	s.imageCollection = GetInstagram().SearchHashtagForVideos(s.hashtag)
}

// LikeCollectedImages likes all the images the scout has collected
func (s *Scout) LikeCollectedImages() {
	for _, item := range s.imageCollection {
		item.Like()
	}
}

// LikeCollectedVideos likes all the videos the scout has collected
func (s *Scout) LikeCollectedVideos() {
	for _, item := range s.videoCollection {
		item.Like()
	}
}

// GetBestImage from the collection
func (s *Scout) GetBestImage() (goinsta.Item, error) {
	mostLikes := -1
	var bestItem goinsta.Item

	for _, item := range s.imageCollection {
		if item.Likes > mostLikes {
			bestItem = item
			mostLikes = item.Likes
		}
	}
	return bestItem, nil
}

// GetBestVideo from the collection
func (s *Scout) GetBestVideo() (goinsta.Item, error) {
	mostLikes := -1
	var bestItem goinsta.Item

	for _, item := range s.videoCollection {
		if item.Likes > mostLikes {
			bestItem = item
			mostLikes = item.Likes
		}
	}
	return bestItem, nil
}
