package main

import (
	"io"
	"log"

	"gopkg.in/ahmdrz/goinsta.v2"
)

var client *instaClient

const (
	sessionPath = ".goinsta-session"
)

func initInstaClient() error {
	var insta *goinsta.Instagram
	// if sessionExists(sessionPath) {
	// 	var err error
	// 	insta, err = goinsta.Import(sessionPath)
	// 	if err != nil {
	// 		log.Printf("couldn't cache the session: %s", err.Error())
	// 		insta = goinsta.New(Config.Creds.Username, Config.Creds.Password)
	// 	}
	// } else {
	insta = goinsta.New(Config.Creds.Username, Config.Creds.Password)
	// }

	client = &instaClient{
		insta,
	}
	err := client.Login()
	if err != nil {
		log.Panicf("couldn't login to instagram: %s", err.Error())
		return err
	}

	err = client.Export(sessionPath)
	if err != nil {
		log.Printf("couldn't export login session: %s", err.Error())
		return err
	}

	log.Printf("[instagram-client] instagram client initiated for user: %s", client.Account.Username)
	return nil
}

// Instagram is an interface to interact with instagram
type Instagram interface {
	SearchHashtagForImages(hashtag string) []goinsta.Item
	SearchHashtagForVideos(hashtag string) []goinsta.Item
	SearchHashtagForAll(hashtag string) []goinsta.Item
	Upload(imageFile io.ReadCloser, caption string) error
	MyPosts(options ...interface{}) []goinsta.Item
}

type instaClient struct {
	*goinsta.Instagram
}

// GetInstagram return the single instagram client
func GetInstagram() Instagram {
	if client == nil {
		initInstaClient()
	}
	return client
}

func (ic *instaClient) SearchHashtagForImages(hashtag string) []goinsta.Item {
	log.Printf("[instagram-client] looking for images with hashtag %s", hashtag)
	items := ic.searchHashtags(hashtag, "image")
	log.Printf("[instagram-client] Found %d images for the hashtag %s", len(items), hashtag)
	return items
}

func (ic *instaClient) SearchHashtagForVideos(hashtag string) []goinsta.Item {
	log.Printf("[instagram-client] looking for videos with hashtag %s", hashtag)
	items := ic.searchHashtags(hashtag, "video")
	log.Printf("[instagram-client] Found %d videos for the hashtag %s", len(items), hashtag)
	return items
}

func (ic *instaClient) SearchHashtagForAll(hashtag string) []goinsta.Item {
	log.Printf("[instagram-client] looking for all posts with hashtag %s", hashtag)
	items := ic.searchHashtags(hashtag, "all")
	log.Printf("[instagram-client] Found %d posts for the hashtag %s", len(items), hashtag)
	return items
}

func (ic *instaClient) Upload(imageFile io.ReadCloser, caption string) error {
	log.Printf("[instagram-client] uploading post")
	defer imageFile.Close()
	_, err := ic.UploadPhoto(imageFile, caption, 1, 1)
	if err != nil {
		return err
	}
	log.Printf("[instagram-client] post uploaded")
	return nil
}

func (ic *instaClient) MyPosts(options ...interface{}) []goinsta.Item {
	feed := ic.Account.Feed()
	feed.Next()
	if len(options) == 1 && len(feed.Items) >= options[0].(int) {
		return feed.Items[:options[0].(int)]
	}
	return feed.Items
}

func (ic *instaClient) searchHashtags(hashtag, postType string) []goinsta.Item {
	var items []goinsta.Item
	res := ic.NewHashtag(hashtag)
	for res.Next() {
		for _, section := range res.Sections {
			for _, media := range section.LayoutContent.Medias {
				if postType == "video" && len(media.Item.Videos) > 0 {
					items = append(items, media.Item)
				} else if postType == "image" && len(media.Item.Videos) < 1 {
					items = append(items, media.Item)
				} else if postType == "all" {
					items = append(items, media.Item)
				}
			}
		}
		break
	}
	return items
}
