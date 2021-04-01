package resource

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Haski007/fav-music-bot/internal/fmb/persistance/repository"

	"github.com/Haski007/fav-music-bot/internal/fmb/persistance/model"
	"google.golang.org/api/youtube/v3"

	"github.com/Haski007/fav-music-bot/pkg/emoji"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func (bot *FMBService) commandPlaylistRegHandler(update tgbotapi.Update) {
	defer func() {
		if recoveryErr := recover(); recoveryErr != nil {
			message := fmt.Sprintf("Panic [commandRegPlaylistHandler] err: %s", recoveryErr)
			bot.ReportToTheCreator(message)
			logrus.Errorf(message)
		}
	}()

	chatID := update.Message.Chat.ID

	arguments := strings.Fields(update.Message.CommandArguments())
	if len(arguments) == 0 {
		bot.Reply(chatID, "Please use playlist id as your command argument!\n/reg_playlist [playlist_id]")
		return
	}

	playlistID := arguments[0]

	found, err := bot.YoutubeService.GetPlaylist(playlistID)
	if err != nil {
		bot.Reply(chatID, "Internal Error! "+emoji.NoEntry+"\nWrite to @pdemian to get some help")
		bot.ReportToTheCreator(
			fmt.Sprintf("[GetPlaylist] playlist_id: %s | err: %s", playlistID, err))
	}

	var data *youtube.Playlist
	for _, p := range found {
		data = p
		break
	}

	playlist := &model.Playlist{
		ID:    playlistID,
		Title: data.Snippet.Title,
	}

	if err := bot.ChatRepository.AddPlaylist(chatID, playlist); err != nil {
		if errors.Is(err, repository.ErrPlaylistAlreadyExists) {
			bot.Reply(chatID, "You already have registered playlist! "+emoji.NoEntry)
			return
		}
		bot.Reply(chatID, "Internal Error! "+emoji.NoEntry+"\nWrite to @pdemian to get some help")
		bot.ReportToTheCreator(
			fmt.Sprintf("[AddPlaylist] chatID: %d | err: %s", chatID, err))
		return
	}

	bot.Reply(chatID, "Playlist [*"+playlist.Title+"*] successfully added "+emoji.Check)
}
