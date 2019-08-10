package command

import (
	"github.com/spf13/cobra"
	"github.com/stobita/airnote/internal/migration"
	"github.com/stobita/airnote/internal/server"
)

// RootCommand start web server
var RootCommand = &cobra.Command{
	RunE: runStart,
}

// StartCommand start web server
var StartCommand = &cobra.Command{
	Use:  "start",
	RunE: runStart,
}

func runStart(cmd *cobra.Command, args []string) error {
	return server.Run()
}

// MigrateCommand run db migration
var MigrateCommand = &cobra.Command{
	Use:  "migrate",
	RunE: runMigrate,
}

func runMigrate(cmd *cobra.Command, args []string) error {
	return migration.Run()
}
