package main

import (
	"log"

	"gopkg.in/ahmdrz/goinsta.v2"
)

// Scout interface
type Scout interface {
	ScoutImages()
	ScoutVideos()
	LikeCollectedImages()
	LikeCollectedVideos()
	GetBestImage() (*goinsta.Item, error)
	GetBestVideo() (*goinsta.Item, error)
}

// NewScout constructor
func NewScout(hashtag string) Scout {
	return &scoutClient{
		hashtag: hashtag,
	}
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
	s.videoCollection = GetInstagram().SearchHashtagForVideos(s.hashtag)
}

// LikeCollectedImages likes all the images the scout has collected
func (s *scoutClient) LikeCollectedImages() {
	tempLiked := 0
	for _, item := range s.imageCollection {
		if Liked < int(Config.LikeThreshold) {
			item.Like()
			Liked++
			tempLiked++
		}
	}
	log.Printf("liked %d posts", tempLiked)
}

// LikeCollectedVideos likes all the videos the scout has collected
func (s *scoutClient) LikeCollectedVideos() {
	tempLiked := 0
	for _, item := range s.videoCollection {
		if Liked < int(Config.LikeThreshold) {
			item.Like()
			Liked++
			tempLiked++
		}
	}
	log.Printf("liked %d posts", tempLiked)
}

// GetBestImage from the collection
func (s *scoutClient) GetBestImage() (*goinsta.Item, error) {
	return getBestItem(s.imageCollection)
}

// GetBestVideo from the collection
func (s *scoutClient) GetBestVideo() (*goinsta.Item, error) {
	return getBestItem(s.videoCollection)
}
