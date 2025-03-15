/*
Copyright © 2025 Matt Redmond <me@mattredmond.com>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize xpns",
	Long:  `Initializes the xpns database and runs any pending migrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := executeInit(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func executeInit() error {
	if err := createWorkdir(); err != nil {
		return err
	}
	if err := createDatabase(); err != nil {
		return err
	}
	if err := runMigrations(); err != nil {
		return err
	}
	return nil
}

func createWorkdir() error {
	if _, err := os.Stat(workdir()); err == nil {
		return nil
	}

	if err := os.Mkdir(workdir(), 0755); err != nil {
		return err
	}
	return nil
}

func createDatabase() error {
	if _, err := os.Stat(dbPath()); err == nil {
		return nil
	}

	if _, err := os.Create(dbPath()); err != nil {
		return err
	}

	return nil
}

// //go:embed sql/schema/*.sql
//var embedMigrations embed.FS

func runMigrations() error {
	db, err := sql.Open("sqlite", dbPath())
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(nil)

	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}

	if err := goose.Up(db, "sql/schema"); err != nil {
		return err
	}

	return nil
}
