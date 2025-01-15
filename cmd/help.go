package cmd

import (
	"fmt"
	"strings"
)

func Help(filename string) {
	message := 
	`Помощь:
	%s help - Вывести это сообщение
	%s add "example" - Добавить задачу "example"
	%s update "id" "example" - Изменить текст задачи с ID "id" на "example"
	%s delete "id" - Удалить задачу с ID "id"
	%s list [status] - Посмотреть все задачи и их ID [Вывести задачи с определённым статусом (in-progress, done, todo)]
	%s mark done/in-progress/todo - Пометить задачи`
	formatted := strings.ReplaceAll(message, "%s", filename)
	fmt.Println(formatted)
}