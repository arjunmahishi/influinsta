package main

import (
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
		if Liked < int(Config.LikeThreshold) {
			item.Like()
			Liked++
		}
	}
}

// LikeCollectedVideos likes all the videos the scout has collected
func (s *scoutClient) LikeCollectedVideos() {
	for _, item := range s.videoCollection {
		if Liked < int(Config.LikeThreshold) {
			item.Like()
			Liked++
		}
	}
}

// GetBestImage from the collection
func (s *scoutClient) GetBestImage() (*goinsta.Item, error) {
	return getBestItem(s.imageCollection)
}

// GetBestVideo from the collection
func (s *scoutClient) GetBestVideo() (*goinsta.Item, error) {
	return getBestItem(s.videoCollection)
}
