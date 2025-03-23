# xpns

Command-line expense tracker written in Go. Uses a SQLite database in the backend, so you might need drivers installed
depending on your operating system.

## Motivation

I wanted a way to track purchases and expenses throughout the day, and reference them without having to log in to my
banking each time. To solve this, I created `xpns`, a command line interface to add and view transactions. It's not 
intended as a full-bodies expense tracker, more as a way to keep track of expenses at the time of purchase, and as a 
reminder when you need to know what you did spend (or receive).

## Installation instructions

```bash
go install https://github.com/mattr/xpns
```

## Usage and commands

### Init

Configures the application and creates a database with the necessary tables if one does not already exist. Will run
migrations to update the database to the latest version if required.

```shell
xpns init
```

### List

Lists the transactions in the database. Accepts a `--date` (`-d`) flag to list transactions for a given date. Date is
expected to be in ISO format.

List all transactions:

```shell
xpns list
```

List transactions for a specific date:

```shell
xpns list -d 2025-03-01
```

Aliases: `l`

### Credit

Adds a credit transaction in the database. Requires the amount of the transaction in dollars and cents. Also accepts a 
`--note` (`-n`) flag to provide a note on the transaction, and a `--date` (`-d`) flag to specify a date (defaults to the
current date if not specified).

```shell
xpns credit 1500 -n salary
```

Specify an alternate date for the transaction:

```shell
xpns credit 1500 -n salary -d 2025-03-21
```

Aliases: `c`, `in`, `payment`

### Debit

Like credit, but adds a debit transaction in the database. Requires the amount and accepts the same flags.

```shell
xpns debit 17.50 -n "burger for lunch"
```

Aliases: `d`, `out`, `purchase`, `expense`

## Roadmap

- [ ] Allow the date flag to accept a range
- [ ] Balance command to show credits - debits (either total or for a given date)
- [ ] Configuration via the `init` command
- [ ] GitHub actions to build releases
- [ ] Recurring credits and debits
