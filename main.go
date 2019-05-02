package main

import (
	"log"
	"sync"

	"gopkg.in/ahmdrz/goinsta.v2"
)

func main() {
	chosenVideos := []goinsta.Item{}

	var wg sync.WaitGroup
	for _, tag := range Config.Hashtags {
		wg.Add(1)
		go func(tag string) {
			scout := NewScout(tag)
			scout.GetBestVideo()
			scout.LikeCollectedVideos()
			best, err := scout.GetBestVideo()
			if err != nil {
				log.Printf("Couldn't get best video: %s", err.Error())
			}
			chosenVideos = append(chosenVideos, *best)
			wg.Done()
		}(tag)
	}
	wg.Wait()

	video, err := getBestItem(chosenVideos)
	if err != nil {
		log.Printf("couldn't get best video from final list: %s", err.Error())
	}

	handleChosenVideo(*video)
}

func handleChosenVideo(chosen goinsta.Item) {
	log.Printf("handling the chosen video")
	imageFile, err := downloadImage(chosen.Images.GetBest())
	if err != nil {
		log.Printf("Couldn't download image: %s", err.Error())
	}
	defer imageFile.Close()

	currentCaption := chosen.Caption.Text
	authorUsername := chosen.User.Username

	err = Reposter.Publish(imageFile, authorUsername, currentCaption)
	if err != nil {
		log.Printf("Couldn't publish post: %s", err.Error())
	}
}
