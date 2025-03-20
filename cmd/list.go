/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mattr/xpns/internal/database"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all transactions",
	Long: `Lists all the transactions in the database, ordered by the transaction date with most recent first.
List can accept a --date (-d) flag which lists transactions for a specific date. To avoid ambiguity, this is
passed in the format yyyy-mm-dd`,
	Run: func(cmd *cobra.Command, args []string) {
		dateArg, _ := cmd.Flags().GetString("date")
		if dateArg != "" {
			date, err := time.Parse(time.DateOnly, dateArg)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			executeListForDate(date)
			return
		}
		executeList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().StringP("date", "d", "", "List transactions for a specific date (defaults to all dates)")
}

func executeList() {
	_ = withDbConn(func(db *sql.DB) error {
		queries := database.New(db)
		transactions, err := queries.GetAllTransactions(context.Background())
		if err != nil {
			return err
		}
		fmt.Println("Listing all transactions")
		for _, txn := range transactions {
			printTransaction(txn)
		}
		return nil
	})
}

func executeListForDate(date time.Time) {
	_ = withDbConn(func(db *sql.DB) error {
		queries := database.New(db)
		transactions, err := queries.GetTransactionsForDate(context.Background(), date.Format("2006-01-02"))
		if err != nil {
			return err
		}
		fmt.Printf("Listing transactions for %v\n", date.Format("2006-01-02"))
		for _, txn := range transactions {
			printTransaction(txn)
		}
		return nil
	})
}

func printTransaction(txn database.Transaction) {
	amount := float64(txn.AmountCents) / 100
	year, month, day := txn.TransactedOn.Date()

	var note string
	if txn.Note.Valid {
		note = txn.Note.String
	}
	fmt.Printf("[%s] %v-%v-%v\t%s\t$%.2f\t%s\n", txn.ID, day, month, year, txn.TransactionType, amount, note)
}
