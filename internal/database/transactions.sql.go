// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: transactions.sql

package database

import (
	"context"
	"time"
)

const createCredit = `-- name: CreateCredit :one
insert into transactions(id, created_at, updated_at, transaction_type, transacted_on, amount_cents)
values (?, datetime(), datetime(), "credit", ?, ?)
returning id, created_at, updated_at, transaction_type, transacted_on, amount_cents
`

type CreateCreditParams struct {
	ID           interface{}
	TransactedOn time.Time
	AmountCents  int64
}

func (q *Queries) CreateCredit(ctx context.Context, arg CreateCreditParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createCredit, arg.ID, arg.TransactedOn, arg.AmountCents)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TransactionType,
		&i.TransactedOn,
		&i.AmountCents,
	)
	return i, err
}

const createDebit = `-- name: CreateDebit :one
insert into transactions(id, created_at, updated_at, transaction_type, transacted_on, amount_cents)
values (?, datetime(), datetime(), "debit", ?, ?)
returning id, created_at, updated_at, transaction_type, transacted_on, amount_cents
`

type CreateDebitParams struct {
	ID           interface{}
	TransactedOn time.Time
	AmountCents  int64
}

func (q *Queries) CreateDebit(ctx context.Context, arg CreateDebitParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createDebit, arg.ID, arg.TransactedOn, arg.AmountCents)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TransactionType,
		&i.TransactedOn,
		&i.AmountCents,
	)
	return i, err
}
