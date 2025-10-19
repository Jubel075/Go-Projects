package main

import (
	"database/sql"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID          int
	Description string
	Done        bool
	Priority    string // "Low", "Medium", "High"
	DueDate     string // Format: "2006-01-02"
}

var db *sql.DB

// Initialize database
func initDB() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dbPath := filepath.Join(homeDir, ".todoapp", "todos.db")
	os.MkdirAll(filepath.Dir(dbPath), 0755)

	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	db = database

	// Create table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		done BOOLEAN DEFAULT 0,
		priority TEXT DEFAULT 'Medium',
		due_date TEXT DEFAULT ''
	);`

	_, err = db.Exec(createTableSQL)
	if err == nil {
		// Add new columns if they don't exist (for existing databases)
		db.Exec("ALTER TABLE todos ADD COLUMN priority TEXT DEFAULT 'Medium'")
		db.Exec("ALTER TABLE todos ADD COLUMN due_date TEXT DEFAULT ''")
	}
	return nil
}

// Load todos from database
func loadTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT id, description, done, priority, due_date FROM todos ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Description, &todo.Done, &todo.Priority, &todo.DueDate)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// Save new todo to database
func saveTodo(description, priority, dueDate string) error {
	_, err := db.Exec("INSERT INTO todos (description, done, priority, due_date) VALUES (?, ?, ?, ?)",
		description, false, priority, dueDate)
	return err
}

// Update todo
func updateTodo(id int, done bool, priority, dueDate string) error {
	_, err := db.Exec("UPDATE todos SET done = ?, priority = ?, due_date = ? WHERE id = ?",
		done, priority, dueDate, id)
	return err
}

// Delete todo from database
func deleteTodo(id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

// Delete all completed todos
func deleteCompletedTodos() error {
	_, err := db.Exec("DELETE FROM todos WHERE done = ?", true)
	return err
}

// Delete all todos
func deleteAllTodos() error {
	_, err := db.Exec("DELETE FROM todos")
	return err
}

func NewTodoFromDataItem(item binding.DataItem) Todo {
	v, _ := item.(binding.Untyped).Get()
	return v.(Todo)
}

// Get color based on priority
func getPriorityColor(priority string) color.Color {
	switch priority {
	case "High":
		return color.NRGBA{R: 255, G: 67, B: 67, A: 255} // Red
	case "Medium":
		return color.NRGBA{R: 255, G: 193, B: 7, A: 255} // Amber
	case "Low":
		return color.NRGBA{R: 76, G: 175, B: 80, A: 255} // Green
	default:
		return color.NRGBA{R: 158, G: 158, B: 158, A: 255} // Gray
	}
}

// Check if due date is overdue
func isOverdue(dueDate string) bool {
	if dueDate == "" {
		return false
	}
	due, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		return false
	}
	return due.Before(time.Now())
}

func main() {
	// Initialize database
	err := initDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	myApp := app.New()
	myWindow := myApp.NewWindow("Todo List App")
	myWindow.Resize(fyne.NewSize(500, 600))

	// Set window icon
	if iconResource, err := fyne.LoadResourceFromPath("icon.png"); err == nil {
		myWindow.SetIcon(iconResource)
	}

	// Create data binding list for todos
	todos := binding.NewUntypedList()

	// Load existing todos from database
	savedTodos, err := loadTodos()
	if err == nil {
		for _, todo := range savedTodos {
			todos.Append(todo)
		}
	}

	// Input field for new todos
	newTodoInput := widget.NewEntry()
	newTodoInput.SetPlaceHolder("Enter new todo...")

	// Priority selector
	prioritySelect := widget.NewSelect([]string{"Low", "Medium", "High"}, func(s string) {})
	prioritySelect.SetSelected("Medium")

	// Due date picker
	dueDateInput := widget.NewEntry()
	dueDateInput.SetPlaceHolder("YYYY-MM-DD (optional)")

	// Add button
	addButton := widget.NewButton("Add Todo", nil)
	addButton.Disable()

	// Enable button only when input is valid
	newTodoInput.OnChanged = func(text string) {
		if len(text) >= 3 {
			addButton.Enable()
		} else {
			addButton.Disable()
		}
	}

	// Add button functionality
	addButton.OnTapped = func() {
		description := newTodoInput.Text
		priority := prioritySelect.Selected
		dueDate := dueDateInput.Text

		// Validate date format if provided
		if dueDate != "" {
			if _, err := time.Parse("2006-01-02", dueDate); err != nil {
				// Invalid date format
				dueDate = ""
			}
		}

		err := saveTodo(description, priority, dueDate)
		if err == nil {
			updatedTodos, _ := loadTodos()
			todos.Set(make([]interface{}, 0))
			for _, todo := range updatedTodos {
				todos.Append(todo)
			}
			newTodoInput.SetText("")
			dueDateInput.SetText("")
			prioritySelect.SetSelected("Medium")
		}
	}

	// Create dynamic list widget with enhanced styling
	todoList := widget.NewListWithData(
		todos,
		// Template for each list item
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewCheck("", nil),
				canvas.NewRectangle(color.White),
				container.NewVBox(
					widget.NewLabel(""),
					widget.NewLabel(""),
				),
			)
		},
		// Update function for each item
		func(di binding.DataItem, obj fyne.CanvasObject) {
			hbox := obj.(*fyne.Container)
			checkbox := hbox.Objects[0].(*widget.Check)
			priorityBox := hbox.Objects[1].(*canvas.Rectangle)
			rightContainer := hbox.Objects[2].(*fyne.Container)
			label := rightContainer.Objects[0].(*widget.Label)
			dateLabel := rightContainer.Objects[1].(*widget.Label)

			todo := NewTodoFromDataItem(di)
			label.SetText(todo.Description)
			checkbox.SetChecked(todo.Done)
			priorityBox.FillColor = getPriorityColor(todo.Priority)
			priorityBox.SetMinSize(fyne.NewSize(12, 12))

			// Format due date display
			if todo.DueDate != "" {
				dateStr := "Due: " + todo.DueDate
				if isOverdue(todo.DueDate) {
					dateStr += " ⚠️ OVERDUE"
				}
				dateLabel.SetText(dateStr)
			} else {
				dateLabel.SetText("No due date")
			}

			checkbox.OnChanged = func(checked bool) {
				updateTodo(todo.ID, checked, todo.Priority, todo.DueDate)
				todo.Done = checked
				di.(binding.Untyped).Set(todo)
			}
		},
	)

	// Clear completed button
	removeCompletedButton := widget.NewButton("Remove Completed", func() {
		err := deleteCompletedTodos()
		if err == nil {
			updatedTodos, _ := loadTodos()
			todos.Set(make([]interface{}, 0))
			for _, todo := range updatedTodos {
				todos.Append(todo)
			}
		}
	})

	// Clear all button
	clearButton := widget.NewButton("Clear All", func() {
		err := deleteAllTodos()
		if err == nil {
			todos.Set([]interface{}{})
		}
	})

	// Input form container
	inputForm := container.NewVBox(
		widget.NewLabel("Task Description:"),
		newTodoInput,
		widget.NewLabel("Priority:"),
		prioritySelect,
		widget.NewLabel("Due Date (YYYY-MM-DD):"),
		dueDateInput,
		addButton,
	)

	inputContainer := container.NewBorder(
		nil, nil, nil, nil,
		inputForm,
	)

	// Buttons container
	buttonContainer := container.NewGridWithColumns(2,
		removeCompletedButton,
		clearButton,
	)

	// Header
	headerBox := canvas.NewRectangle(color.NRGBA{R: 63, G: 81, B: 181, A: 255})
	headerBox.SetMinSize(fyne.NewSize(0, 60))
	headerText := canvas.NewText("📝 My Todo List", color.White)
	headerText.TextSize = 24
	headerText.TextStyle.Bold = true
	headerContainer := container.NewCenter(
		container.NewVBox(
			headerBox,
			headerText,
		),
	)

	// Assemble the UI
	content := container.NewBorder(
		headerContainer,
		container.NewVBox(inputContainer, buttonContainer),
		nil, nil,
		todoList,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
