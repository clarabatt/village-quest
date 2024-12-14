# Village Quest

## Run the application

```bash
    go run cmd/*.go

    CompileDaemon -command="go run cmd/*.go"
```

## 🗂️ Folder Structure

- cmd/: Entry point for the CLI application (main.go).
- internal/domain/: Contains aggregates (Resources, Buildings, Population, Events).
- internal/services/: Business logic for managing resources, buildings, and events.
