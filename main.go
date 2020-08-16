package main

import (
	"fmt"
	"os"

	"github.com/bantl23/goup/cmd"
	"github.com/mitchellh/cli"
)

func main() {
	app := cli.NewCLI("goup", cmd.Vers())
	app.Args = os.Args[1:]
	app.Commands = map[string]cli.CommandFactory{
		"show": func() (cli.Command, error) {
			return &cmd.Todo{}, nil
		},
		"update": func() (cli.Command, error) {
			return &cmd.Todo{}, nil
		},
		"check": func() (cli.Command, error) {
			return &cmd.Todo{}, nil
		},
		"list": func() (cli.Command, error) {
			return &cmd.Todo{}, nil
		},
		"version": func() (cli.Command, error) {
			return &cmd.Version{}, nil
		},
	}

	status, err := app.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(status)
}
