# Village Quest

## Run the application

Install Go: https://golang.org/doc/install
Install Sqlite3: https://www.sqlite.org/download.html

```bash
    go run cmd/villagequest/main.go
```

```bash
    go build
    go install
    villageQuest
```

### Run tests

```bash
    go test ./...
    go test -cover ./...
```

### Run linter

```bash
    golangci-lint run
```

## 📝 Game Description

The objective of the game is to build a village from scratch and manage resources to grow the village until it becomes a city.

### Game Mechanics

- The game is turn-based.
- Each turn represents a year.
- The player can only choose two actions per turn.

Turns are divided into 3 actions:

1. Collect resources.
2. Player action.
   a. Build.
   b. Upgrade.
   c. Allocate workers.
   d. Collect taxes.
3. Events.

### Resources

- Food
- People
- Wood
- Stone
- Gold

### Buildings

- House
- Mill
- Farm
- Lumber Mill
- Quarry
- Mine

### Events

A pool of events that can happen each turn. Events can be positive or negative.
Events can affect resources, buildings, or population.
Events can be affected by the player's actions and game's indicators.

## 🗂️ Folder Structure

- cmd/: Entry point for the CLI application (main.go).
- internal/domain/: Contains aggregates (Resources, Buildings, Population, Events).
- internal/services/: Business logic for managing resources, buildings, and events.
