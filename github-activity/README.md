# github-activity

A simple CLI tool that fetches and displays a GitHub user's recent public activity using the GitHub API.

## Usage

```
github-activity <username>
```

### Example

```
github-activity torvalds
```

Output:

```
Fetching GitHub user activity for torvalds
+-----+----------------------+--------------------------------+
| #   | Event                | Repository                     |
+-----+----------------------+--------------------------------+
| 1   | Pushed commits       | torvalds/linux                 |
| 2   | Starred              | some/repo                      |
| 3   | Forked repository    | another/repo                   |
+-----+----------------------+--------------------------------+
```

## Supported Event Types

| Event | Description |
|-------|-------------|
| `PushEvent` | Pushed commits |
| `WatchEvent` | Starred a repository |
| `IssuesEvent` | Opened an issue |
| `ForkEvent` | Forked a repository |
| `CreateEvent` | Created a repository |
| Other | Displays the raw event type |

## Installation

**Prerequisites:** Go 1.25+

```bash
git clone https://github.com/rsmharsha/Go-Projects
cd Go-Projects/github-activity
go build -o github-activity .
```

## How It Works

1. Accepts a GitHub username as a command-line argument
2. Queries `https://api.github.com/users/<username>/events`
3. Parses the JSON response and maps event types to human-readable descriptions
4. Displays results in a formatted table

## Project URL

https://roadmap.sh/projects/github-user-activity
