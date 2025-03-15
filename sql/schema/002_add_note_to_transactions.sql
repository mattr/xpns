-- +goose Up
alter table transactions
    add column note text;

-- +goose Down
alter table transactions
    remove column note;
