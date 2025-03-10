-- name: CreateCredit :one
insert into transactions(id, created_at, updated_at, transaction_type, transacted_on, amount_cents)
values (?, datetime(), datetime(), "credit", ?, ?)
returning id, created_at, updated_at, transaction_type, transacted_on, amount_cents;

-- name: CreateDebit :one
insert into transactions(id, created_at, updated_at, transaction_type, transacted_on, amount_cents)
values (?, datetime(), datetime(), "debit", ?, ?)
returning id, created_at, updated_at, transaction_type, transacted_on, amount_cents;
