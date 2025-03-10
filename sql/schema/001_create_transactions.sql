-- +goose Up
create table transactions
(
    id               uuid primary key,
    created_at       timestamp not null,
    updated_at       timestamp not null,
    transaction_type string    not null default "",
    transacted_on    date      not null,
    amount_cents     integer   not null
);

-- +goose Down
drop table transactions;
