# CLI-Cobra

A lightweight and elegant command-line to-do manager built in Go, powered by the Cobra library.  
This tool allows quick task tracking directly from your terminal — no web interfaces, no clutter.

---

## Features

- Simple task management with add, list, and complete commands  
- Persistent JSON storage for offline use  
- Clean tabular output for clear task visualization  
- Built using [Cobra](https://github.com/spf13/cobra) for a structured CLI design  
- Cross-platform support (Windows, macOS, Linux)

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/jubel075/cli-cobra.git
   cd cli-cobra
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. To install it globally:
   ```bash
   go install
   ```

---

## Usage

### Add a Task
Add one or more tasks to your to-do list.
```bash
cli-cobra add "Finish project report"
cli-cobra add "Buy groceries" "Call client"
```

### List Tasks
Display all stored tasks in a clean, tabular format.
```bash
cli-cobra list
```

Example output:
```
You have 3 tasks in your to-do list:

LABEL   PRIORITY    TASK
-----   --------    ----
1.      1           Finish project report
2.      2           Buy groceries
3.      1           Call client
```

### Complete a Task
Mark a task as completed by its label or index.
```bash
cli-cobra done 2
```

---

## Configuration

By default, tasks are stored in:
```
<project_root>/.todo.json
```

You can specify a custom data file:
```bash
cli-cobra --datafile /path/to/custom.json add "Example task"
```

---

## Project Structure

```
cli-cobra/
├── cmd/                 # Cobra command definitions
│   ├── root.go
│   ├── add.go
│   └── list.go
├── todo/                # Core logic for reading/writing tasks
│   └── todo.go
├── go.mod
├── main.go
└── README.md
```

---

## Built With

- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra)
- Standard Library (encoding/json, os, fmt, text/tabwriter)

---

## License

This project is licensed under the MIT License.  
See the [LICENSE](LICENSE) file for details.

---

## Author

Developed by Guilian Kasandikromo  
Repository: [github.com/jubel075/cli-cobra](https://github.com/jubel075/cli-cobra)
