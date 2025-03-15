/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mattr/xpns/internal/database"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// debitCmd represents the debit command
var debitCmd = &cobra.Command{
	Use:     "debit",
	Aliases: []string{"d", "out", "purchase"},
	Args:    cobra.ExactArgs(1),
	Short:   "Debits an amount from the balance",
	Long: `debit adds a 'debit' transaction to the transactions table, indicating money going out.
Example:
	xpns debit 17.50
`,
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := executeDebit(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Credited %.2f\n", amount)
	},
}

func init() {
	rootCmd.AddCommand(debitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// debitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func executeDebit(amount string) (float64, error) {
	amountCents, err := getAmountCents(amount)
	if err != nil {
		return 0, err
	}
	amountCents, err = applyDebit(amountCents)
	if err != nil {
		return 0, err
	}

	return float64(amountCents) / 100, nil
}

func applyDebit(amount int64) (int64, error) {
	db, err := sql.Open("sqlite", dbPath())
	if err != nil {
		return 0, err
	}
	defer db.Close()

	params := database.CreateDebitParams{ID: uuid.New(), TransactedOn: time.Now(), AmountCents: amount}
	queries := database.New(db)
	debit, err := queries.CreateDebit(context.Background(), params)
	if err != nil {
		return 0, err
	}
	return debit.AmountCents, nil
}
