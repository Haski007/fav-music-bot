package resource

import (
	"github.com/Haski007/fav-music-bot/internal/fmb/config"
	"github.com/Haski007/fav-music-bot/internal/fmb/persistance/repository/mongodb"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type FMBService struct {
	Bot       *tgbotapi.BotAPI
	CreatorID int64

	ChatRepository *mongodb.ChatRepository
}

func NewFMBService() (*FMBService, error) {
	var err error

	bot := &FMBService{}

	var cfg config.Config
	cfg.Parse(config.ConfigFile)

	bot.CreatorID = cfg.Bot.CreatorID

	bot.ChatRepository = &mongodb.ChatRepository{}
	bot.ChatRepository.InitChatsConn(cfg.MongoDB)

	bot.Bot, err = tgbotapi.NewBotAPI(cfg.Bot.GetToken().String())
	if err != nil {
		return nil, err
	}

	bot.Bot.Debug = true

	logrus.Printf("Authorized on account %s", bot.Bot.Self.UserName)

	return bot, nil
}
