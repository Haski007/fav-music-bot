package fmb

import (
	"context"
	"github.com/Haski007/fav-music-bot/internal/fmb/resource"
	"github.com/Haski007/fav-music-bot/pkg/factory"
	"github.com/Haski007/fav-music-bot/pkg/run"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

func Run(args *run.Args) error {

	//b, err := ioutil.ReadFile("secrets/client_secret.json")
	//if err != nil {
	//	log.Fatalf("Unable to read client secret file: %v", err)
	//}
	//
	//// If modifying these scopes, delete your previously saved credentials
	//// at ~/.credentials/youtube-go-quickstart.json
	//config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	//if err != nil {
	//	logrus.Fatalf("Unable to parse client secret file to config: %v", err)
	//}
	//_ = api.GetClient(ctx, config)
	//
	//service, err := youtube.NewService(ctx)
	//if err != nil {
	//	logrus.Fatalf("[NewFMBService] err: %s", err)
	//}

	//api.HandleError(err, "Error creating YouTube client")

	//api.ChannelsListByUsername(service, "snippet,contentDetails,statistics", "GoogleDevelopers")

	//os.Exit(-1)

	botService, err := resource.NewFMBService()
	if err != nil {
		logrus.Fatalf("[NewFMBService] err: %s", err)
	}

	factory.InitLog(args.LogLevel)

	StartBot(botService)
	return nil
}
