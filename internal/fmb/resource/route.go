package resource

import (
	"fmt"

	"github.com/Haski007/fav-music-bot/pkg/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (bot *FMBService) HandleRoutes(updates tgbotapi.UpdatesChannel) {
	botCreds, err := bot.Bot.GetMe()
	if err != nil {
		bot.ReportToTheCreator(
			fmt.Sprintf("[bot GetMe] err: %s", err))
		return
	}

	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		// Check if someone removes bot
		if update.Message.LeftChatMember != nil &&
			update.Message.LeftChatMember.UserName == botCreds.UserName {
			bot.ChatRepository.RemoveChat(update.Message.Chat.ID)
		}

		if command := update.Message.CommandWithAt(); command != "" {
			switch {
			case command == "help" || command == "help"+"@"+botCreds.UserName:
				go bot.commandHelpHandler(update)
			case command == "reg_chat" || command == "reg_chat"+"@"+botCreds.UserName:
				go bot.commandRegNewChatHandler(update)
			case command == "del_chat" || command == "del_chat"+"@"+botCreds.UserName:
				go bot.commandDelChatHandler(update)
			case command == "reg_playlist" || command == "reg_playlist"+"@"+botCreds.UserName:
				go bot.commandPlaylistRegHandler(update)
			case command == "clear_playlist" || command == "clear_playlist"+"@"+botCreds.UserName:
				go bot.commandPlaylistClearHandler(update)
			default:
				bot.Reply(update.Message.Chat.ID, "Such command does not exist! "+emoji.NoEntry)
			}
		}
	}
}
