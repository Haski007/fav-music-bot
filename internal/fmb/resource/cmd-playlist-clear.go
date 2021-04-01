package resource

import (
	"fmt"

	"github.com/Haski007/fav-music-bot/pkg/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func (bot *FMBService) commandPlaylistClearHandler(update tgbotapi.Update) {
	defer func() {
		if recoveryErr := recover(); recoveryErr != nil {
			message := fmt.Sprintf("Panic [commandRegPlaylistHandler] err: %s", recoveryErr)
			bot.ReportToTheCreator(message)
			logrus.Errorf(message)
		}
	}()

	chatID := update.Message.Chat.ID

	if err := bot.ChatRepository.ClearPlaylist(chatID); err != nil {
		bot.Reply(chatID, "Internal Error! "+emoji.NoEntry+"\nWrite to @pdemian to get some help")
		bot.ReportToTheCreator(
			fmt.Sprintf("[AddPlaylist] chatID: %d | err: %s", chatID, err))
		return
	}

	bot.Reply(chatID, "Playlist successfully cleared "+emoji.Basket)
}
