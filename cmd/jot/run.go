package main

import (
	"os"
	"time"

	"github.com/AdamShannag/jot/internal/box"
	"github.com/AdamShannag/jot/internal/command"
	"github.com/AdamShannag/jot/internal/command/log"
	"github.com/AdamShannag/jot/internal/config"
	"github.com/urfave/cli/v2"
)

func run(args []string) {
	app := &cli.App{
		Name:     "jot",
		Version:  config.VERSION,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Adam Shannaq",
				Email: "adamsgtrs@gmail.com",
			},
			{
				Name:  "Mohammad Yassen",
				Email: "mohammad.t.yaseen1@gmail.com",
			},
		},
		Copyright: "(c) 2023 jot",
		HelpName:  "help",
		Usage:     "jot - quickly generate microservices and related components",
		UsageText: "help - displays information about available commands.",
		Commands:  command.Commands(),
		Action: func(*cli.Context) error {
			box.Welcome()
			return nil
		},
	}

	if err := app.Run(args); err != nil {
		log.Info(err.Error(), log.FAILED)
		os.Exit(1)
	}
}
