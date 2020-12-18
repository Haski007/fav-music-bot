package main

import (
	"fmt"
	"os"

	"github.com/Haski007/fav-music-bot/internal/fmb"
	"github.com/Haski007/fav-music-bot/pkg/run"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var Version string

func main() {
	app := cli.App{
		Name:    "fav-videos-fmb",
		Usage:   "Favourite tiktok videos tg fmb entry point",
		Version: Version,
		Action: func(c *cli.Context) error {
			if err := fmb.Run(&run.Args{
				LogLevel: run.LogLevel(c.String("info")),
			}); err != nil {
				return fmt.Errorf("run: %w", err)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
