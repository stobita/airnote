package main

import (
	"log"

	"github.com/stobita/airnote/internal/command"
)

func main() {
	log.Print("arinote start")
	rootCmd := command.RootCommand
	rootCmd.AddCommand(
		command.StartCommand,
		command.MigrateCommand,
	)
	rootCmd.Execute()
}
