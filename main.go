package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/izya4ka/notes-cli/cmd"
)

const filename = "tasks.json"

func main() {
    fmt.Print("notes-cli 1.0\n\n")
    if len(os.Args) == 1 {
        fmt.Println("Ошибка! Не представленно ни одного аргумента!")
        cmd.Help(os.Args[0])
        os.Exit(1)
    }
    
    switch os.Args[1] {
    case "help":
        cmd.Help(os.Args[0])
    case "add":
        if (len(os.Args) < 3) {
            fmt.Println(`Недостаточно аргументов к команде "add"!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        cmd.Add(filename, os.Args[2])
    case "list":
        if (len(os.Args) < 3) {
            cmd.List(filename, "all")
        } else {
            if (os.Args[2] != "done" && os.Args[2] != "todo" && os.Args[2] != "in-progress") {
                fmt.Println("Доступно только три типа статуса! См. помощь")
                os.Exit(1)
            } 
            cmd.List(filename, os.Args[2])
        }
    case "delete":
        if (len(os.Args) < 3) {
            fmt.Println(`Недостаточно аргументов к команде "delete"!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        id, cerr := strconv.Atoi(os.Args[2])
        if cerr != nil {
            fmt.Println(`Аргументом к команде "delete" должно быть число!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        cmd.Delete(filename, id)
    case "update":
        if (len(os.Args) < 4) {
            fmt.Println(`Недостаточно аргументов к команде "delete"!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        id, cerr := strconv.Atoi(os.Args[2])
        if cerr != nil {
            fmt.Println(`Аргументом к команде "update" должно быть число!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        cmd.Update(filename, id, os.Args[3])
    case "mark":
        if (len(os.Args) < 4) {
            fmt.Println(`Недостаточно аргументов к команде "mark"!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        id, cerr := strconv.Atoi(os.Args[2])
        if cerr != nil {
            fmt.Println(`Аргументом к команде "update" должно быть число!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        if (os.Args[3] != "done" && os.Args[3] != "todo" && os.Args[3] != "in-progress") {
            fmt.Println("Доступно только три типа статуса! См. помощь")
            os.Exit(1)
        } 
        cmd.Mark(filename, id, os.Args[3])

    default:
        fmt.Println("Ошибка! Такой команды нет")
        cmd.Help(os.Args[0])
        os.Exit(1)
    }

}