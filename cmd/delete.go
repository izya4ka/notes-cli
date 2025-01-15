package cmd

import (
	"fmt"
	"syscall"
)

func Delete(filename string, id int) {
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
			temp_index = index
			break
		}
	}
	if temp_index == -1 {
		fmt.Println("Нет задач с таким ID!")
		return 
	}

	all_tasks = append(all_tasks[:temp_index], all_tasks[temp_index+1:]...)
	WriteTasks(filename, all_tasks)
	fmt.Println("Задача успешно удалена!")
}