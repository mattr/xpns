package main

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mattr/xpns/internal/database"
	"github.com/pressly/goose/v3"
	"os"
	"strconv"
	"time"
)

//go:embed sql/schema/*.sql
var embedMigrations embed.FS

type command struct {
	name        string
	description string
	callback    func([]string, *config) error
}

func commandHelp(args []string, cfg *config) error {
	fmt.Println("xpns usage:")
	fmt.Println("  xpns <command> [<args>]")
	fmt.Println("")
	fmt.Println("  commands:")
	fmt.Println("    help - print this help message")
	fmt.Println("    init - initialize the application")
	return nil
}

func commandInit(args []string, cfg *config) error {
	err := os.MkdirAll(cfg.workDir, 0755)
	if err != nil {
		return err
	}
	_, _ = os.Create(cfg.workDir + "/xpns.db")

	var db *sql.DB
	db, err = sql.Open("sqlite", cfg.workDir+"/xpns.db")
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}

	if err := goose.Up(db, "sql/schema"); err != nil {
		return err
	}

	return nil
}

func commandCredit(args []string, cfg *config) error {
	if len(args) == 0 {
		return errors.New("credit requires an amount")
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return err
	}
	amountCents := int64(amount * 100)

	db, err := sql.Open("sqlite", cfg.workDir+"/xpns.db")
	if err != nil {
		return err
	}
	defer db.Close()

	params := database.CreateCreditParams{ID: uuid.New(), TransactedOn: time.Now(), AmountCents: amountCents}
	queries := database.New(db)
	credit, err := queries.CreateCredit(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("Credited $%.2f\n", float64(credit.AmountCents)/100)
	return nil
}

func commandDebit(args []string, cfg *config) error {
	if len(args) == 0 {
		return errors.New("debit requires an amount")
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return err
	}
	amountCents := int64(amount * 100)

	db, err := sql.Open("sqlite", cfg.workDir+"/xpns.db")
	if err != nil {
		return err
	}
	defer db.Close()

	params := database.CreateDebitParams{ID: uuid.New(), TransactedOn: time.Now(), AmountCents: amountCents}
	queries := database.New(db)
	credit, err := queries.CreateDebit(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("Debited $%.2f\n", float64(credit.AmountCents)/100)
	return nil
}
