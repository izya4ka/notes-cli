package cmd

import (
	"fmt"
	"strings"
)

func Help(filename string) {
	message := 
	`Help:
	%s help - Display this message
	%s add "example" - Add task with "example" description
	%s update "id" "example" - Change task description with ID "id" to "example"
	%s delete "id" - Delete task with ID "id"
	%s list [status] - List all tasks [List with status (in-progress, done, todo)]
	%s mark "id" done/in-progress/todo - Mark task with ID`
	formatted := strings.ReplaceAll(message, "%s", filename)
	fmt.Println(formatted)
}