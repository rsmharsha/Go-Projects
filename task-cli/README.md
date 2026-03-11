# Task Tracker CLI

A simple command-line task manager written in Go. Tasks are persisted locally in a `tasks.json` file.

## Installation

```bash
git clone https://github.com/your-username/task-tracker.git
cd task-tracker
go build -o task-cli .
```

## Usage

```
task-cli <command> [arguments]
```

## Commands

| Command                       | Description                   |
| ----------------------------- | ----------------------------- |
| `add "<description>"`         | Add a new task                |
| `list`                        | List all tasks                |
| `list <status>`               | List tasks filtered by status |
| `update <id> "<description>"` | Update a task's description   |
| `delete <id>`                 | Delete a task                 |
| `mark-in-progress <id>`       | Mark a task as `in-progress`  |
| `mark-done <id>`              | Mark a task as `done`         |

## Examples

```bash
# Add tasks
task-cli add "Buy groceries"
task-cli add "Write unit tests"

# List all tasks
task-cli list

# List by status
task-cli list todo
task-cli list in-progress
task-cli list done

# Update a task
task-cli update 1 "Buy groceries and cook dinner"

# Mark tasks
task-cli mark-in-progress 2
task-cli mark-done 1

# Delete a task
task-cli delete 3
```

## Task Statuses

- `todo` — default status when a task is created
- `in-progress` — task is currently being worked on
- `done` — task is completed

## Data Storage

Tasks are stored in a `tasks.json` file in the current working directory, created automatically on first use.

```json
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "done",
    "createdAt": "2025-01-01T10:00:00Z",
    "updatedAt": "2025-01-01T12:00:00Z"
  }
]
```

## Requirements

- Go 1.18 or higher
