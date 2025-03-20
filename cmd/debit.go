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
		note, _ := cmd.Flags().GetString("note")
		dateArg, _ := cmd.Flags().GetString("date")
		if dateArg != "" {
			date, err := time.Parse(time.DateOnly, dateArg)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			amount, err := executeDebit(args[0], note, date)
			fmt.Printf("Debited %.2f on %v", amount, date.Format("2006-01-02"))
			if note != "" {
				fmt.Printf(" for %s", note)
			}
			fmt.Println()
			return
		}
		amount, err := executeDebit(args[0], note, time.Now().Local())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Debited %.2f", amount)
		if note != "" {
			fmt.Printf(" for %s", note)
		}
		fmt.Println()
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
	debitCmd.Flags().StringP("note", "n", "", "Include a note on the transaction")
	debitCmd.Flags().StringP("date", "d", "", "Set the date of the transaction (defaults to the current date)")
}

func executeDebit(amount, note string, date time.Time) (float64, error) {
	amountCents, err := getAmountCents(amount)
	if err != nil {
		return 0, err
	}

	amountCents, err = applyDebit(amountCents, note, date)
	if err != nil {
		return 0, err
	}

	return float64(amountCents) / 100, nil
}

func applyDebit(amount int64, note string, date time.Time) (int64, error) {
	err := withDbConn(func(db *sql.DB) error {
		noteParam := sql.NullString{String: note, Valid: true}
		params := database.CreateDebitParams{ID: uuid.New(), TransactedOn: date, AmountCents: amount, Note: noteParam}
		queries := database.New(db)
		_, err := queries.CreateDebit(context.Background(), params)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return amount, nil
}
