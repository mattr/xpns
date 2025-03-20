# xpns

Command-line expense tracker written in Go. Uses a SQLite database in the backend, so you might need drivers installed
depending on your operating system.

## Installation instructions

```bash
go install https://github.com/mattr/xpns
```

## Commands

### Init

Configures the application and creates a database with the necessary tables if one does not already exist. Will run
migrations to update the database to the latest version if required.

```shell
xpns init
```

### List

Lists the transactions in the database. Accepts a `--date` (`-d`) flag to list transactions for a given date. Date is
expected to be in ISO format.

```shell
xpns list -d 2025-03-01
```

### Credit

Adds a credit transaction in the database. Requires the amount of the transaction in dollars and cents. Also accepts a 
`--note` (`-n`) flag to provide a note on the transaction, and a `--date` (`-d`) flag to specify a date (defaults to the
current date if not specified).

```shell
xpns credit 1500 -n salary
```

Aliases: `in`, `payment`

### Debit

Like credit, but adds a debit transaction in the database. Requires the amount and accepts the same flags.

```shell
xpns debit 17.50 -n "burger for lunch"
```

Aliases: `out`, `purchase`

## Roadmap

- [ ] Allow the date flag to accept a range
- [ ] Balance command to show credits - debits (either total or for a given date)
- [ ] Configuration via the `init` command
- [ ] GitHub actions to build releases
- [ ] Recurring credits and debits
