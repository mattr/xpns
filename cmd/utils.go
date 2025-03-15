package cmd

import (
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
