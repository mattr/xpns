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

// creditCmd represents the credit command
var creditCmd = &cobra.Command{
	Use:     "credit [amount]",
	Aliases: []string{"in", "payment"},
	Args:    cobra.ExactArgs(1),
	Short:   "Credit an amount to your balance",
	Long: `credit adds a 'credit' transaction to the database, indicating money coming in.
Example:
	xpns credit 5
`,
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := executeCredit(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Debited %.2f\n", amount)
	},
}

func init() {
	rootCmd.AddCommand(creditCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// creditCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// creditCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func executeCredit(amount string) (float64, error) {
	amountCents, err := getAmountCents(amount)
	if err != nil {
		return 0, err
	}
	amountCents, err = applyCredit(amountCents)
	if err != nil {
		return 0, err
	}

	return float64(amountCents) / 100, nil
}

func applyCredit(amount int64) (int64, error) {
	db, err := sql.Open("sqlite", dbPath())
	if err != nil {
		return 0, err
	}
	defer db.Close()

	params := database.CreateCreditParams{ID: uuid.New(), TransactedOn: time.Now(), AmountCents: amount}
	queries := database.New(db)
	credit, err := queries.CreateCredit(context.Background(), params)
	if err != nil {
		return 0, err
	}
	return credit.AmountCents, nil
}
