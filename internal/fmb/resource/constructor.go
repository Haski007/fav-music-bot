package resource

import (
	"github.com/Haski007/fav-music-bot/api"
	"github.com/Haski007/fav-music-bot/internal/fmb/config"
	"github.com/Haski007/fav-music-bot/internal/fmb/persistance/repository/mongodb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type FMBService struct {
	Bot       *tgbotapi.BotAPI
	CreatorID int64

	YoutubeService *api.YoutubeService
	ChatRepository *mongodb.ChatRepository
}

func NewFMBService() (*FMBService, error) {
	var err error

	bot := &FMBService{}

	var cfg config.Config
	if err := cfg.Parse(config.ConfigFile); err != nil {
		logrus.Fatalf("[NewFMBService] cfg.Parse | err: %s", err)
	}

	bot.CreatorID = cfg.Bot.CreatorID

	bot.ChatRepository = &mongodb.ChatRepository{}
	bot.ChatRepository.InitChatsConn(cfg.MongoDB)

	bot.Bot, err = tgbotapi.NewBotAPI(cfg.Bot.GetToken().String())
	if err != nil {
		return nil, err
	}

	bot.YoutubeService = api.NewYoutubeService()

	bot.Bot.Debug = true

	logrus.Printf("Authorized on account %s", bot.Bot.Self.UserName)

	return bot, nil
}
