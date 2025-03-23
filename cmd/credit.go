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
	Aliases: []string{"c", "in", "payment"},
	Args:    cobra.ExactArgs(1),
	Short:   "Credit an amount to your balance",
	Long: `credit adds a 'credit' transaction to the database, indicating money coming in.
Example:
	xpns credit 5
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
			amount, err := executeCredit(args[0], note, date)
			fmt.Printf("Credited %.2f on %v", amount, date.Format("2006-01-02"))
			if note != "" {
				fmt.Printf(" for %s", note)
			}
			fmt.Println()
			return
		}
		amount, err := executeCredit(args[0], note, time.Now().Local())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Credited %.2f", amount)
		if note != "" {
			fmt.Printf(" for %s", note)
		}
		fmt.Println()
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
	creditCmd.Flags().StringP("note", "n", "", "Include a note on the transaction")
	creditCmd.Flags().StringP("date", "d", "", "Specify the date of the transaction (default: today)")
}

func executeCredit(amount, note string, date time.Time) (float64, error) {
	amountCents, err := getAmountCents(amount)
	if err != nil {
		return 0, err
	}

	amountCents, err = applyCredit(amountCents, note, date)
	if err != nil {
		return 0, err
	}

	return float64(amountCents) / 100, nil
}

func applyCredit(amount int64, note string, date time.Time) (int64, error) {
	err := withDbConn(func(db *sql.DB) error {
		noteParam := sql.NullString{String: note, Valid: true}
		params := database.CreateCreditParams{ID: uuid.New(), TransactedOn: date, AmountCents: amount, Note: noteParam}
		queries := database.New(db)
		_, err := queries.CreateCredit(context.Background(), params)
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
