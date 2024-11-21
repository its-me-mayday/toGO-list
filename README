# ToGO-list

A command-line application to manage a to-do list, written in Go, that uses SQLite for data persistence.

### Author
Luca (its-me-mayday)

---

## How to Run the Program

### Prerequisites

1. **Go**: Make sure you have Go installed. You can check by running `go version` in your terminal.
   - To install Go, visit the [official documentation](https://golang.org/doc/install).
2. **SQLite**: SQLite does not need to be installed separately, as the program uses a Go library that manages SQLite internally.
3. **SQLite Driver**: The program uses the `github.com/mattn/go-sqlite3` driver to interact with SQLite.

### Installation

1. Clone the repository:
   ```bash
   $ git clone https://github.com/its-me-mayday/todo-cli.git
   $ cd todo-cli
   ```
2. Install dependencies (if not already installed):

  ```bash
  $ go mod tidy
  ```

3. Run the program:

  ```bash
  $ go run main.go
  ```

## Usage
### Available Commands

- Add a Task:

  ```bash
  $ go run main.go -add "Buy milk"
  ```
This will add a task to the SQLite database. The task description should be provided within quotes.

- List Tasks:

  ```bash
  $ go run main.go -list
  ```
This command will list all tasks stored in the SQLite database, showing whether each task is completed or not.

- Complete a Task:

  ```bash
  $ go run main.go -complete 1
  ```
This will mark the task with ID 1 as completed in the database.

- Delete a Task:

```bash
$ go run main.go -delete 1
```
This will delete the task with ID 1 from the database.

### Design Patterns Used
1. Command Pattern
The Command Pattern is used to handle the various commands in the application. Each command (such as "Add Task", "Complete Task", etc.) is represented as an object that implements a Command interface with an Execute method. This pattern decouples the request for an action from its actual implementation, making the code easier to extend and test.

This pattern makes it easy to add new commands in the future without modifying the rest of the application.

2. Singleton Pattern (hypothetical)
If we add the possibility of maintaining a single shared database connection, the Singleton pattern can be introduced to manage the SQLite connection.

3. Repository Pattern (future improvement)
In the future, the code can be refactored to use the Repository Pattern, where interaction with the database is abstracted behind an interface that handles all CRUD operations, separating data access logic from the application logic.

## Roadmap

### Phase 1: Implemented Features
- **Add Task to the SQLite database (`-add`)**:
  - The command adds a task to the database and stores it in a `tasks` table.

### Phase 2: Future Features
- **List Tasks (`-list`)**:
  - Displays all tasks stored in the database with an indication of whether they are completed or not.
  
- **Complete a Task (`-complete`)**:
  - Implements the command to mark a task as completed in the database.

- **Delete a Task (`-delete`)**:
  - Command to remove a task from the database.

### Phase 3: Improvements and Optimizations
- **Better Error Handling**:
  - Improve error handling for commands, providing more descriptive error messages.
  
- **Automated Tests**:
  - Add unit tests to verify correct behavior of commands and database interactions.

- **Advanced User Interface**:
  - Implement filtering options to display only completed or incomplete tasks.

- **Advanced Persistence**:
  - Add features like editing tasks, searching by name, or sorting by date.