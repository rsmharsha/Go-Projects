# Contact Book

A simple command-line contact management application written in Go. Contacts are stored locally in a JSON file.

## Features

- Add contacts with name, email, and phone
- List all contacts
- Delete a contact by ID
- Search contacts by name or email

## Installation

```bash
git clone https://github.com/rsmharsha/Go-Projects
cd Go-Projects/contact-book
go build -o contact-book .
```

## Usage

```bash
contact-book <command> [arguments]
```

### Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `add` | Add a new contact | `contact-book add <name> <email> <phone>` |
| `list` | List all contacts | `contact-book list` |
| `delete` | Delete a contact by ID | `contact-book delete <id>` |
| `search` | Search by name or email | `contact-book search <query>` |

### Examples

```bash
# Add a contact
contact-book add "John Doe" john@example.com 555-1234

# List all contacts
contact-book list

# Search contacts
contact-book search john

# Delete a contact
contact-book delete 1
```

## Data Storage

Contacts are saved in `contacts.json` in the current working directory. Each contact stores:
- `id` — auto-incremented integer
- `name`, `email`, `phone`
- `createdAt`, `updatedAt` — timestamps
