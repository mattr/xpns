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
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all transactions",
	Long:  `Lists all the transactions in the database, ordered by the transaction date with most recent first`,
	Run: func(cmd *cobra.Command, args []string) {
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
}

func executeList() {
	_ = withDbConn(func(db *sql.DB) error {
		queries := database.New(db)
		transactions, err := queries.GetAllTransactions(context.Background())
		if err != nil {
			return err
		}
		for _, txn := range transactions {
			amount := float64(txn.AmountCents) / 100
			year, month, day := txn.TransactedOn.Date()
			fmt.Printf("[%s] %v-%v-%v\t%s\t$%.2f\n", txn.ID, day, month, year, txn.TransactionType, amount)
		}
		return nil
	})
}
