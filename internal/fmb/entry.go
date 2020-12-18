package fmb

import (
	"github.com/Haski007/fav-music-bot/internal/fmb/resource"
	"github.com/Haski007/fav-music-bot/pkg/factory"
	"github.com/Haski007/fav-music-bot/pkg/run"
	"github.com/sirupsen/logrus"
)

func Run(args *run.Args) error {

	botService, err := resource.NewFMBService()
	if err != nil {
		logrus.Fatalf("[NewFMBService] err: %s", err)
	}

	factory.InitLog(args.LogLevel)

	StartBot(botService)
	return nil
}
