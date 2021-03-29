package resource

import (
	"fmt"
	"log"
	"time"

	"github.com/Haski007/fav-music-bot/internal/fmb/resource/transform"

	"github.com/Haski007/fav-music-bot/internal/fmb/persistance/model"
	"github.com/Haski007/fav-music-bot/pkg/emoji"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FMBService) CheckNewMusic() {
	//defer func() {
	//	if recoveryErr := recover(); recoveryErr != nil {
	//		message := fmt.Sprintf("Panic [CheckNewMusic] err: %s", recoveryErr)
	//		bot.ReportToTheCreator(message)
	//		logrus.Errorf(message)
	//	}
	//}()

	var chats []model.Chat
	bot.ChatRepository.GetAllChats(&chats)

	for _, chat := range chats {
		likes, err := bot.YoutubeService.GetLikedIDs(3)
		if err != nil {
			logrus.Printf("[YoutubeService.GetLikedIDs] err: %s", err)
			return
		}

		music := transform.DecodeYoutubeVideos(likes)

		for _, m := range music {
			if bot.ChatRepository.PostedMusicExists(chat.ID, m.ID) {
				continue
			}

			log.Println("Posting", m.ID)

			var message string

			message += fmt.Sprintf("%s *Liked music* %s\n", emoji.Heart, emoji.Heart)
			message += fmt.Sprintf(
				"Title: *%s*\n"+
					"Author: *%s*\n",
				m.Title,
				m.Author)

			//imgURL, err := url.Parse(m.Image)
			//if err != nil {
			//	logrus.Errorf("[CheckNewMusic] url.Parse | err: %s", err)
			//}
			resp := tgbotapi.NewPhotoUpload(chat.ID, nil)
			resp.FileID = m.Image
			resp.UseExisting = true
			resp.ReplyMarkup = model.NewOriginalURLMarkup(m.URL)
			resp.Caption = message
			resp.ParseMode = "MarkDown"

			if message, err := bot.Bot.Send(resp); err != nil {
				logrus.Errorf("[bot Send] message: %+v | err: %s", message, err)
			}
			if err := bot.ChatRepository.PushPostedMusic(chat.ID, m.ID); err != nil {
				bot.Reply(chat.ID, "Internal Error! "+emoji.NoEntry+"\nWrite to @pdemian to get some help")
				bot.ReportToTheCreator(
					fmt.Sprintf("[PushPostedMusic] chatID: %d | videoID: %s", chat.ID, m.ID))
			}
			time.Sleep(time.Second * 3)
		}
	}
}
