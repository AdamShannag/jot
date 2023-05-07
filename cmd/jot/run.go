package main

import (
	"os"
	"time"

	"github.com/AdamShannag/jot/internal/box"
	"github.com/AdamShannag/jot/internal/command"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func run(args []string) {
	app := &cli.App{
		Name:     "jot",
		Version:  VERSION,
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
		color.Red(err.Error())
		os.Exit(1)
	}
}
