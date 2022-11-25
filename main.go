package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-mastodon"
)

var (
	MASTODON_SERVER        = os.Getenv("MASTODON_SERVER")
	MASTODON_CLIENT_ID     = os.Getenv("MASTODON_CLIENT_ID")
	MASTODON_CLIENT_SECRET = os.Getenv("MASTODON_CLIENT_SECRET")
	MASTODON_USERNAME      = os.Getenv("MASTODON_USERNAME")
	MASTODON_PASSWORD      = os.Getenv("MASTODON_PASSWORD")

	VIDEO_PATH     = "./media/video.mp4"
	THUMBNAIL_PATH = "./media/thumbnail.png"
)

func main() {
	// make sure all needed environment variables are set
	if MASTODON_SERVER == "" || MASTODON_CLIENT_ID == "" || MASTODON_CLIENT_SECRET == "" || MASTODON_USERNAME == "" || MASTODON_PASSWORD == "" {
		log.Fatal(`you need to set the environment vars:
	MASTODON_SERVER - the mastodon server url. ex: https://mastodon.social
	MASTODON_CLIENT_ID - your application's client id
	MASTODON_CLIENT_SECRET - your application's client secret
	MASTODON_USERNAME - your mastodon account username
	MASTODON_PASSWORD - your mastodon account password`)
	}

	// create a new mastodon client
	c := mastodon.NewClient(&mastodon.Config{
		Server:       MASTODON_SERVER,
		ClientID:     MASTODON_CLIENT_ID,
		ClientSecret: MASTODON_CLIENT_SECRET,
	})

	// authenticate with mastodon server
	err := c.Authenticate(context.Background(), MASTODON_USERNAME, MASTODON_PASSWORD)
	if err != nil {
		log.Fatal("authentication failed: ", err)
	}

	// create media file handler
	mediaFile, err := os.Open(VIDEO_PATH)
	if err != nil {
		log.Fatal("unable to open media file: ", err)
	}

	// create thumbnail file handler
	thumbnail, err := os.Open(THUMBNAIL_PATH)
	if err != nil {
		log.Fatal("unable to open thumbnail file: ", err)
	}

	// create mastodon.Media struct
	media := &mastodon.Media{
		File:        bufio.NewReader(mediaFile),
		Thumbnail:   thumbnail,
		Description: "Daniel Craig wishing everyone a happy weekend",
	}

	// upload and get the media reference
	uploadedMedia, err := c.UploadMediaFromMedia(context.Background(), media)
	if err != nil {
		log.Fatal("failed while uploading media: ", err)
	}

	// create mastodon.Toot struct
	toot := &mastodon.Toot{
		MediaIDs: []mastodon.ID{uploadedMedia.ID},
	}

	// publish the toot
	status, err := c.PostStatus(context.Background(), toot)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("toot sent:", status.URL)
}
