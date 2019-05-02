package main

import (
	"io"
	"log"

	"gopkg.in/ahmdrz/goinsta.v2"
)

var client *instaClient

func init() {
	client = &instaClient{
		goinsta.New(Config.Creds.Username, Config.Creds.Password),
	}
	err := client.Login()
	if err != nil {
		panic("couldn't login to instagram")
	}
	log.Printf("instagram client initiated for user: %s", client.Account.Username)
}

// Instagram is an interface to interact with instagram
type Instagram interface {
	SearchHashtagForImages(hashtag string) ([]goinsta.Item, error)
	SearchHashtagForVideos(hashtag string) ([]goinsta.Item, error)
	Upload(imageFile io.ReadCloser, caption string) error
}

type instaClient struct {
	*goinsta.Instagram
}

// GetInstagram return the single instagram client
func GetInstagram() Instagram {
	return client
}

func (ic *instaClient) SearchHashtagForImages(hashtag string) ([]goinsta.Item, error) {
	var items []goinsta.Item
	res := ic.NewHashtag(hashtag)
	for res.Next() {
		for _, section := range res.Sections {
			for _, media := range section.LayoutContent.Medias {
				if len(media.Item.Videos) < 1 {
					items = append(items, media.Item)
				}
			}
		}
		break
	}
	return items, nil
}

func (ic *instaClient) SearchHashtagForVideos(hashtag string) ([]goinsta.Item, error) {
	var items []goinsta.Item
	res := ic.NewHashtag(hashtag)
	for res.Next() {
		for _, section := range res.Sections {
			for _, media := range section.LayoutContent.Medias {
				if len(media.Item.Videos) > 0 {
					items = append(items, media.Item)
				}
			}
		}
		break
	}
	return items, nil
}

func (ic *instaClient) Upload(imageFile io.ReadCloser, caption string) error {
	defer imageFile.Close()
	_, err := ic.UploadPhoto(imageFile, caption, 1, 1)
	if err != nil {
		return err
	}
	return nil
}
