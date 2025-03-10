package main

import (
	"fmt"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

type config struct {
	//db      *database.Queries
	workDir string
}

func main() {
	commands := map[string]command{
		"credit": {
			name:        "credit",
			description: "Credit an amount to the balance",
			callback:    commandCredit,
		},
		"debit": {
			name:        "debit",
			description: "Debit an amount from the balance",
			callback:    commandDebit,
		},
		"help": {
			name:        "help",
			description: "print this help message",
			callback:    commandHelp,
		},
		"in": {
			name:        "in",
			description: "alias for `credit`",
			callback:    commandCredit,
		},
		"init": {
			name:        "init",
			description: "Initialise the xpns application",
			callback:    commandInit,
		},
		"out": {
			name:        "out",
			description: "alias for `debit`",
			callback:    commandDebit,
		},
		"payment": {
			name:        "payment",
			description: "alias for `credit`",
			callback:    commandCredit,
		},
		"purchase": {
			name:        "purchase",
			description: "alias for `debit`",
			callback:    commandDebit,
		},
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	cfg := &config{workDir: filepath.Join(dir, ".xpns")}

	args := os.Args[1:]
	if len(args) == 0 {
		_ = commandHelp(nil, cfg)
		os.Exit(1)
	}
	err = commands[args[0]].callback(args[1:], cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
