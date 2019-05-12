package main

import (
	"fmt"
	"log"
	"sync"

	"gopkg.in/ahmdrz/goinsta.v2"
)

// Action type def
type Action func(args ...interface{}) error

// Actions maps strings with functions
var Actions = map[string]Action{
	"reshare-video":   reshareBestVideos,
	"random-comments": makeRandomComments,
	"random-follow":   followRandomPeople,
}

func performAction(actionName string, args ...interface{}) error {
	action, err := getAction(actionName)
	if err != nil {
		return err
	}
	log.Printf("performing action '%s'", actionName)
	return action(args)
}

func getAction(actionName string) (Action, error) {
	if action, ok := Actions[actionName]; ok {
		return action, nil
	}
	return nil, fmt.Errorf("the action '%s' does not exist", actionName)
}

func reshareBestVideos(args ...interface{}) error {
	chosenVideos := []goinsta.Item{}

	var wg sync.WaitGroup
	for _, tag := range Config.Hashtags {
		wg.Add(1)
		go func(tag string) {
			log.Printf("Creating a scout for the hashtag %s", tag)
			scout := NewScout(tag)
			scout.ScoutVideos()
			scout.LikeCollectedVideos()
			best, err := scout.GetBestVideo()
			if err != nil {
				log.Printf("Couldn't get best video: %s", err.Error())
			} else {
				chosenVideos = append(chosenVideos, *best)
			}
			wg.Done()
		}(tag)
	}
	wg.Wait()

	video, err := selectBestItem(chosenVideos)
	if err != nil {
		return fmt.Errorf("couldn't get best video from final list: %s", err.Error())
	}

	log.Printf("Chose a post by %s", video.User.Username)
	handleChosenVideo(*video)
	return nil
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

	chosen.User.Follow()
	log.Printf("followed %s", authorUsername)

	err = Reposter.Publish(imageFile, authorUsername, currentCaption)
	if err != nil {
		log.Printf("Couldn't publish post: %s", err.Error())
	}
}

func makeRandomComments(args ...interface{}) error {
	commetCount := 0
	var wg sync.WaitGroup
	for _, tag := range Config.Hashtags {
		wg.Add(1)
		go func(tag string) {
			posts := GetInstagram().SearchHashtagForVideos(tag)
			for _, post := range posts {
				err := post.Comments.Add(getRandomGenericComment())
				if err == nil {
					commetCount++
				}
			}
			wg.Done()
		}(tag)
	}
	wg.Wait()
	log.Printf("commented on %d post(s)", commetCount)
	return nil
}

func followRandomPeople(args ...interface{}) error {
	return nil
}
