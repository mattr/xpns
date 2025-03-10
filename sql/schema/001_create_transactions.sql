-- +goose Up
create table transactions
(
    id               uuid primary key,
    created_at       timestamp not null default current_timestamp,
    updated_at       timestamp not null default current_timestamp,
    transaction_type string    not null default "",
    transacted_on    date      not null default current_date,
    amount_cents     integer   not null
);

-- +goose Down
drop table transactions;
