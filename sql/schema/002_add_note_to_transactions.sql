-- +goose Up
alter table transactions
    add column note text;

-- +goose Down
alter table transactions
    drop column note;
