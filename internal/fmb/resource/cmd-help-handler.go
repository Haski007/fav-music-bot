package resource

import (
	"github.com/Haski007/fav-music-bot/internal/fmb/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FMBService) commandHelpHandler(update tgbotapi.Update) {

	message := "Here is bot to subscribe on someone's playlist in Youtube\n" +
		"You can use such commands:\n" +
		" - /help - for help\n" +
		" - /reg_chat - for adding current chat to DataBase\n" +
		" - /del_chat - for removing chat\n\n" +
		"How to subscribe on playlist:\n" +
		"1) Open page in browser with playlist you want to subscribe\n" +
		"2) Copy playlist id an showed on image below\n" +
		"3) Use command in chat with bot\n" +
		"/reg_playlist [paste id here]"

	resp := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, nil)
	resp.FileID = config.GuideGetPlaylistIDImageURL
	resp.UseExisting = true
	resp.Caption = message
	bot.Bot.Send(resp)
}
