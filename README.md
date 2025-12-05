# Todo CLI App in Go

A very basic todo CLI app written in Go.

## Installation

1. **Run `install-cli.sh`**  
   This installs the todo CLI globally. (You need to have GoLang installed for this to work.)
2. **Uninstall**  
   To uninstall, run:
   ```bash
   rm /usr/local/bin/todo
   ```

## Usage

```
$ todo [command] [args]
```

Commands

- **add** <text> [due-date] Add new task (due-date: dd-mm-yyyy, default today)
- **list** List all tasks
- **update** <field> <value>... <id> Update task (fields: text, due, done)
- **done** <id> Complete task by ID
- **delete** <id> Delete task by ID
- **clear** Clear todo list

Examples:

```Bash
$ todo add "Buy groceries" "15-12-2025"
$ todo update text "New text" done true 1
$ todo delete 5
```
