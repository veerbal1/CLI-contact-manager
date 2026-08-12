# CLI Contact Manager

A learning project built in Go. It manages contacts from the terminal and saves them in a local JSON file.

## Features

- Add, list, find, update, and delete contacts
- Validate required fields and basic email format
- Persist contacts in `contacts.json`
- Load saved contacts when the program starts
- Handle missing and malformed JSON safely
- Use packages, interfaces, JSON storage, context cancellation, and automated tests

## Requirements

- Go installed

## Run

From the project directory:

```bash
go run .
```

To show help:

```bash
go run . help
```

## Menu Actions

1. Add contact
2. List contacts
3. Find contact by ID
4. Update contact by ID
5. Delete contact by ID
6. Exit

## Data Storage

Contacts are saved in `contacts.json` in the project directory. The file is created automatically after the first successful add, update, or delete.

## Tests

Run all tests:

```bash
go test ./...
```

Useful quality checks:

```bash
gofmt -w .
go vet ./...
```

## Project Structure

```text
.
├── main.go            # Terminal menu and application wiring
├── contact/           # Contact model and CRUD logic
├── storage/           # JSON and in-memory storage implementations
├── contacts.json      # Saved contact data
└── ROADMAP.md         # Learning roadmap
```
