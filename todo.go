package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
    // TODO: marshal pretty JSON; write to p; perms 0644
	p, err := storePath()
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(todos, "", "")
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
		fmt.Fprintln(os.Stderr, "checkFileExist error:", err)
		os.Exit(1)
	}

	if ok {
		return nil
	}

	f, err := os.OpenFile(p, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil
		}
		return fmt.Errorf("created: %v", err)
	}
	defer f.Close()
	if _, err := f.WriteString("[]"); err != nil {
	return fmt.Errorf("init write: %w", err)
	}
	fmt.Printf("crated: %v", p)
	return nil
}

func addTodo(text string) error {
	text = strings.TrimSpace(text)
	if text == "" {
		return errors.New("todo text required")
	}

	p, err := storePath()
	if err != nil {
		return nil
	}
	if err := ensureStoreFile(p); err != nil {
		return err
	}

	todos, err := loadTodos()
	if err != nil {
		return nil
	}

	t := Todo{
		ID: nextID(todos),
		Task: text,
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
    // TODO: load; find by id; set Completed=true; CompletedAt=now; save
	p, err := storePath()
	if err != nil {
		return nil
	}
	if err := ensureStoreFile(p); err == nil {
		return nil
	}

	todos, err := loadTodos()
	if err != nil {
		return nil
	}
	if strings.Contains(id){
		
	}

}

//func removeTodo(id int) error {
//    // TODO: load; filter out id; save
//}
//
//func clearTodos() error {
//    // TODO: save empty slice
//}
//
//func usage() {
//    // TODO: print commands
//}
//
//func dispatch(args []string) error {
//    // TODO: switch args[0]: add/list/done/rm/clear; parse rest; call above
//}
//
func main() {

}