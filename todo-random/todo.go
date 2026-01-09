package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Todo struct {
	ID int
	Task string
	Completed bool
	Created time.Time
	CompletedAt *time.Time
}

func storePath() (string, error) {
	return filepath.Abs("todo.json")
}

func checkFileExist(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return !info.IsDir(), nil
	}
	if errors.Is(err, os.ErrNotExist){
		return false, nil
	}
	return false, err
}

func saveTodos(todos []Todo) error {
	p, err := storePath()
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, b, 0o644)
}

func loadTodos() ([]Todo, error) {
    // TODO: open p := storePath(); read file; if missing return empty slice
    // TODO: json.Unmarshal into []Todo
	p, err := storePath()
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(p)
	if errors.Is(err, os.ErrNotExist) {
		return []Todo{}, nil
	}

	if err != nil {
		return nil, err
	}

	if len(b) == 0{
		return []Todo{}, nil
	}

	var todos []Todo
	if err := json.Unmarshal(b, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}


func nextID(todos []Todo) int {
    // TODO: scan max id; return max+1
	max := 0 
	for _, t := range todos {
		if t.ID > max{
			max = t.ID
		}
	}
	return max + 1
}

func findIndexByID(todos []Todo, id int) int {
    // TODO: linear search; return index or -1
	for i, t := range todos {
		if t.ID == id {
			return i
		}
	}
	return -1 
}

func ensureStoreFile(p string) error {
	ok, err := checkFileExist(p)
	if err != nil {
		return fmt.Errorf("checkFileExist: %w", err)
	}
	if ok {
		return nil
	}

	f, err := os.OpenFile(p, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil
		}
		return fmt.Errorf("create: %w", err)
	}
	defer f.Close()
	if _, err := f.WriteString("[]"); err != nil {
		return fmt.Errorf("init write: %w", err)
	}
	fmt.Printf("created: %s\n", p)
	return nil
}

func addTodo(text string) error {
	text = strings.TrimSpace(text)
	if text == "" {
		return errors.New("todo text required")
	}

	p, err := storePath()
	if err != nil {
		return err
	}
	if err := ensureStoreFile(p); err != nil {
		return err
	}

	todos, err := loadTodos()
	if err != nil {
		return err
	}

	t := Todo{
		ID:      nextID(todos),
		Task:    text,
		Created: time.Now(),
	}
	todos = append(todos, t)
	return saveTodos(todos)
}

func listTodos() error {
    // TODO: load; print each with status [ ]/[x] and id
	p, err := storePath()
	if err != nil {
		return nil
	}
	if err := ensureStoreFile(p); err != nil {
		return err
	}
	todos,err := loadTodos()
	if err != nil {
		return nil
	}
	if len(todos) == 0 {
		fmt.Println("No Todos")
	}
	for _, t := range todos {
		status := ""
		if t.Completed {
			status = "x"
		}
		fmt.Printf("[%s] %-3d %s  (%s)\n", status, t.ID, t.Task, t.Created.Format("2006-01-02"))
	}
	return nil
}

func markDone(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	p, err := storePath()
	if err != nil {
		return err
	}
	if err := ensureStoreFile(p); err != nil {
		return err
	}

	todos, err := loadTodos()
	if err != nil {
		return err
	}

	i := findIndexByID(todos, id)
	if i == -1 {
		return fmt.Errorf("no todo with id %d", id)
	}

	if !todos[i].Completed {
		now := time.Now()
		todos[i].Completed = true
		todos[i].CompletedAt = &now
	}
	return saveTodos(todos)
}

func removeTodo(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	p, err := storePath()
	if err != nil {
		return err
	}
	if err := ensureStoreFile(p); err != nil {
		return err
	}

	todos, err := loadTodos()
	if err != nil {
		return err
	}

	i := findIndexByID(todos, id)
	if i == -1 {
		return fmt.Errorf("no todo with id %d", id)
	}

	todos = append(todos[:i], todos[i+1:]...)
	return saveTodos(todos)
}

func clearTodos() error {
    // TODO: save empty slice
	return saveTodos([]Todo{})
}

func usage() {
    fmt.Fprintln(os.Stderr, "usage: todo add <text> | todo list | todo done <id> | todo rm <id> | todo clear")
}

func dispatch(args []string) error {
    if len(args) == 0 { usage(); return nil }
    switch args[0] {
    case "add":
        return addTodo(strings.Join(args[1:], " "))
    case "list":
        return listTodos()
    case "done":
        if len(args) < 2 { return errors.New("need id") }
        id, err := strconv.Atoi(args[1]); if err != nil { return err }
        return markDone(id)
    case "rm":
        if len(args) < 2 { return errors.New("need id") }
        id, err := strconv.Atoi(args[1]); if err != nil { return err }
        return removeTodo(id)
    case "clear":
        return saveTodos([]Todo{})
    default:
        usage(); return nil
    }
}
func main() {
	p, err := storePath(); if err != nil { log.Fatal(err) }
	if err := ensureStoreFile(p); err != nil { log.Fatal(err) }
	if err := dispatch(os.Args[1:]); err != nil { log.Fatal(err) }
}
