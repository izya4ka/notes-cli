package cmd

import (
	"fmt"
	"syscall"
	"time"
)

func List(filename, status string) {
	all_tasks, rerr := ReadTasks(filename)
	user_id := syscall.Getuid()
	if rerr != nil {
		fmt.Println("Ошибка: ", rerr)
	}
	for _, task := range all_tasks {
		if (task.UserId == user_id) && ((status == task.Status) || (status == "all")) {
			fmt.Printf("Задача [%d]:\n", task.Id)
			fmt.Println("Описание: ", task.Description)
			fmt.Println("Статус: ", task.Status)
			fmt.Println("Время создания: ", time.Unix(task.CreatedAt, 0).Format("15:04:05 02.01.2006"))
			fmt.Println("Время изменения: ", time.Unix(task.UpdatedAt, 0).Format("15:04:05 02.01.2006"))
			fmt.Println()
		}
	}

}