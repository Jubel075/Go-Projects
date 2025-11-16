# GEMINI Code Assistant

This document provides a comprehensive overview of the Go projects in this repository, designed to give the Gemini AI a deep understanding of the codebase.

## Project Overview

This repository is a collection of Go projects created for the purpose of learning and mastering the Go programming language. The projects are diverse, ranging from a command-line chatbot with AI integration to a to-do list application built with the Cobra framework.

The main projects in this repository are:

*   **`cli-chatbot-go`**: A command-line chatbot with a focus on a beautiful and modern user interface. It can be run with or without AI integration (using the Gemini API).
*   **`CLI-cobra`**: A command-line to-do list application built with the popular Cobra framework. It demonstrates how to build a structured and extensible CLI application in Go.

## Building and Running

### `cli-chatbot-go`

**Prerequisites:**

*   Go 1.24 or higher
*   A Gemini API key (for AI functionality)

**Running the application:**

1.  Navigate to the `cli-chatbot-go` directory:
    ```bash
    cd cli-chatbot-go
    ```
2.  Create a `.env` file in the `cli-chatbot-go` directory with the following content:
    ```
    GEMINI_API_KEY=your_api_key
    AI_ENABLED=true
    ```
3.  Run the application:
    ```bash
    go run main.go
    ```

### `CLI-cobra`

**Prerequisites:**

*   Go 1.24 or higher

**Running the application:**

1.  Navigate to the `CLI-cobra` directory:
    ```bash
    cd CLI-cobra
    ```
2.  Run the application:
    ```bash
    go run main.go
    ```

**Available commands:**

*   `go run main.go add "My new task"`: Add a new task.
*   `go run main.go list`: List all tasks.
*   `go run main.go done 1`: Mark task #1 as done.

## Development Conventions

### `cli-chatbot-go`

*   **Dependencies:** Dependencies are managed with Go modules. The only external dependency is `github.com/joho/godotenv`.
*   **AI Integration:** The AI integration is handled in the `ai` package, which uses the `google.golang.org/genai` package to interact with the Gemini API.
*   **Configuration:** The application is configured through a `.env` file.
*   **Structure:** The project is structured into `main`, `ai`, and `responses` packages.

### `CLI-cobra`

*   **Framework:** The application is built with the Cobra framework, which provides a structured way to create CLI applications.
*   **Commands:** Each command is implemented in its own file in the `cmd` directory.
*   **Data Storage:** The to-do list is stored in a JSON file, which defaults to `~/.todo.json`.
*   **Configuration:** The application can be configured through a `.cli-cobra.yaml` file or environment variables, using the Viper library.
