package fmb

import (
	"fmt"
	"time"

	"github.com/Haski007/fav-music-bot/internal/fmb/resource"
	"github.com/Haski007/fav-music-bot/pkg/graceshut"
	"github.com/Haski007/go-errors"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func StartBot(bot *resource.FMBService) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.Bot.GetUpdatesChan(u)
	if err != nil {
		errors.Println(err)
		return
	}
	defer func() {
		if errR := recover(); errR != nil {
			_, err := bot.Bot.Send(
				tgbotapi.NewMessage(
					bot.CreatorID,
					fmt.Sprintf("[Main panic] err: %+v\n", errR)))
			if err != nil {
				logrus.Fatalf("[defer panic] err: %s", err)
			}
		}
	}()

	go bot.HandleRoutes(updates)
	go tiktokLoop(bot)

	graceshut.Loop()
}

func tiktokLoop(bot *resource.FMBService) {
	for {
		bot.CheckNewMusic()
		fmt.Println("Loop passed!")
		time.Sleep(5 * time.Second)
	}
}
