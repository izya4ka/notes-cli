package cmd

import (
	"fmt"
	"syscall"
	"time"
)

func Add(filename, desc string) {
	tasks, err := ReadTasks(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	var last_id int
	if (len(tasks) == 0) {
		last_id = 1
	} else {
		last_id = tasks[len(tasks)-1].Id
		last_id++
	}
	new_task := Task{
		Id: last_id,
		UserId: syscall.Getuid(),
		Description: desc,
		Status: "todo",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	tasks = append(tasks, new_task)
	WriteTasks(filename, tasks)
	fmt.Println("Task added with ID: ", last_id)
}