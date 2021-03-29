package api

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"

	"github.com/Haski007/fav-music-bot/internal/fmb/config"
	"google.golang.org/api/option"

	"google.golang.org/api/youtube/v3"
)

type YoutubeService struct {
	SRV *youtube.Service
}

func NewYoutubeService() *YoutubeService {
	ctx := context.Background()

	b, err := ioutil.ReadFile(config.GoogleAPIFile)
	if err != nil {
		logrus.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/youtube-go-quickstart.json
	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))

	handleError(err, "Error creating YouTube client")
	channelsListByUsername(service, "snippet,contentDetails,statistics", "GoogleDevelopers")
	return &YoutubeService{SRV: service}
}
