package cmd

import (
	"fmt"
	"syscall"
	"time"
)

func Update(filename string, id int, description string) {
	all_tasks, rerr := ReadTasks(filename)
	if rerr != nil {
		println("Ошибка: ", rerr)
	}

	user_id := syscall.Getuid()
	temp_index := -1

	if (cap(all_tasks) == 0) {
		fmt.Println("В данный момент, у вас нет задач")
		return 
	}
	for index, value := range all_tasks {
		if (value.Id == id) && (user_id == value.UserId) {
			all_tasks[index].Description = description
			all_tasks[index].UpdatedAt = time.Now().Unix()
			temp_index = index
			break
		}
	}
	if temp_index == -1 {
		fmt.Println("Нет задач с таким ID!")
		return 
	}

	WriteTasks(filename, all_tasks)
	fmt.Println("Задача успешно изменена!")
}