package cmd

import (
	"encoding/json"
	"io"
	"os"
)

type Task struct {
    Id int `json:"id"`
	UserId int `json:"user_id"`
    Description string `json:"description"`
    Status string `json:"status"`
    CreatedAt int64 `json:"created_at"`
    UpdatedAt int64 `json:"updated_at"`
} 

func WriteTasks(filename string, tasks []Task) error {

	tasks_serialized, serr := json.Marshal(tasks)
	if serr != nil { 
		return serr
	}

	f, ferr := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if ferr != nil {
		return ferr
	}
	defer f.Close()

	_, werr := f.Write(tasks_serialized)
	if werr != nil {
		return werr
	}
	return nil
}

func ReadTasks(filename string) ([]Task, error) {
	
	tasks_file, ferr := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0600)
	if ferr != nil {
		return []Task{}, ferr
	}
	byte_value, rerr := io.ReadAll(tasks_file)
	if rerr != nil {
		return []Task{}, rerr
	}

	var tasks []Task
	json.Unmarshal(byte_value, &tasks)
	return tasks, nil
}