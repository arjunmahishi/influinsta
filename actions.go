package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/ahmdrz/goinsta"
)

// Action type def
type Action func(args ...interface{}) error

// Actions maps strings with functions
var Actions = map[string]Action{
	"reshare-video":   reshareBestVideo,
	"reshare-image":   reshareBestImage,
	"random-comments": makeRandomComments,
	"random-follow":   followRandomPeople,
	"like-posts":      likePosts,
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

func reshareBestItem(itemType string) error {
	chosenItems := []goinsta.Item{}

	var wg sync.WaitGroup
	for _, tag := range Config.Hashtags {
		wg.Add(1)
		go func(tag string) {
			log.Printf("Creating a scout for the hashtag %s", tag)
			scout := NewScout(tag)

			switch itemType {
			case "video":
				scout.ScoutVideos()
				scout.LikeCollectedVideos()
				best, err := scout.GetBestVideo()
				if err != nil {
					log.Printf("Couldn't get best video: %s", err.Error())
				} else {
					chosenItems = append(chosenItems, *best)
				}
			case "image":
				scout.ScoutImages()
				scout.LikeCollectedImages()
				best, err := scout.GetBestImage()
				if err != nil {
					log.Printf("Couldn't get best video: %s", err.Error())
				} else {
					chosenItems = append(chosenItems, *best)
				}
			}

			wg.Done()
		}(tag)
	}
	wg.Wait()

	item, err := selectBestItem(chosenItems)
	if err != nil {
		return fmt.Errorf("couldn't get best item from final list: %s", err.Error())
	}

	log.Printf("Chose a post by %s", item.User.Username)
	handleChosenItem(*item)
	return nil
}

func reshareBestVideo(args ...interface{}) error {
	return reshareBestItem("video")
}

func reshareBestImage(args ...interface{}) error {
	return reshareBestItem("image")
}

func handleChosenItem(chosen goinsta.Item) {
	log.Printf("handling the chosen item")
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
			posts := GetInstagram().SearchHashtagForAll(tag)
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
	followCount := 0
	var wg sync.WaitGroup
	for _, tag := range Config.Hashtags {
		wg.Add(1)
		go func(tag string) {
			posts := GetInstagram().SearchHashtagForAll(tag)
			for _, post := range posts {
				err := post.User.Follow()
				if err == nil {
					followCount++
				}
			}
			wg.Done()
		}(tag)
	}
	wg.Wait()
	log.Printf("followed %d people", followCount)
	return nil
}

func likePosts(args ...interface{}) error {
	var wg sync.WaitGroup
	for _, tag := range Config.Hashtags {
		wg.Add(1)
		go func(tag string) {
			posts := GetInstagram().SearchHashtagForAll(tag)
			count := 0
			for _, post := range posts {
				if err := post.Like(); err != nil {
					continue
				}
				count++
			}
			log.Printf("Liked %d posts for the tag '%s'", count, tag)
			wg.Done()
		}(tag)
	}
	wg.Wait()
	return nil
}
