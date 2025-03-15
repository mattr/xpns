package cmd

import (
	"database/sql"
	"os"
	"strconv"
)

func workdir() string {
	homedir, _ := os.UserHomeDir()
	workdir := homedir + "/.xpns"
	return workdir
}

func dbPath() string {
	workdir := workdir()
	return workdir + "/xpns.db"
}

func getAmountCents(s string) (int64, error) {
	amount, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return int64(amount * 100), nil
}

func withDbConn(fn func(db *sql.DB) error) error {
	db, err := sql.Open("sqlite", dbPath())
	if err != nil {
		return err
	}
	defer db.Close()
	if err = fn(db); err != nil {
		return err
	}
	return nil
}
