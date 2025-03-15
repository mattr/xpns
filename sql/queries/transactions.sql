-- name: CreateCredit :one
insert into transactions(id, created_at, updated_at, transaction_type, transacted_on, amount_cents, note)
values (?, datetime(), datetime(), "credit", ?, ?, ?)
returning id, created_at, updated_at, transaction_type, transacted_on, amount_cents, note;

-- name: CreateDebit :one
insert into transactions(id, created_at, updated_at, transaction_type, transacted_on, amount_cents, note)
values (?, datetime(), datetime(), "debit", ?, ?, ?)
returning id, created_at, updated_at, transaction_type, transacted_on, amount_cents, note;

-- name: GetAllTransactions :many
select id, created_at, updated_at, transaction_type, transacted_on, amount_cents, note
from transactions
order by transacted_on desc, created_at desc;
