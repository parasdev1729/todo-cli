# TODO CLI — Project Tasks

Follow these steps to turn your cobra scaffold into a real CLI app.

## Task 1: Data model & storage (`todo/todo.go`)

- Implement `Save(path string)` method on `Todo` that writes todos as JSON
- Implement `Load(path string)` function that reads JSON back into a slice
- Use `os.OpenFile`, `json.NewEncoder` / `json.NewDecoder`
- Store the JSON file at `~/.todo-cli/todos.json` (create the dir if missing)
- Add fields: `CreatedAt`, `Done`, `DoneAt`
- Make fields **exported** (capital letter) so JSON works

## Task 2: `add` command (`cmd/add.go`)

- Accept the todo title as an argument
- Parse a `--priority` / `-p` flag (high / mid / low)
- Load existing todos, append a new one, save back
- Print a confirmation message

## Task 3: `list` command (`cmd/list.go`)

- Generate with: `cobra-cli add list`
- Load todos and print them as a table (use `text/tabwriter` for alignment)
- Support a `--done` / `-d` flag to filter by completion status
- Support a `--priority` / `-p` flag to filter by priority level
- Show: index, priority, title, created date, status (done/pending)

## Task 4: `done` command (`cmd/done.go`)

- Accept a todo index (or multiple indices) as arguments
- Load todos, mark the selected ones as done (set `Done = true`, `DoneAt = time.Now()`)
- Save back and print a confirmation

## Task 5: `delete` command (`cmd/delete.go`)

- Accept a todo index (or multiple indices) as arguments
- Load todos, remove the selected items, save back
- Print a confirmation

## Task 6: Polish the root command (`cmd/root.go`)

- Remove the auto-generated `--toggle` flag
- Show a brief help / usage summary when run without subcommands

Try it out:

```bash
go build -o todo-cli && ./todo-cli add "Learn Go" -p high
./todo-cli add "Build a CLI" -p mid
./todo-cli list
./todo-cli done 1
./todo-cli delete 2
```
