package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/izya4ka/notes-cli/cmd"
)

const filename = "/etc/tasks.json"

func main() {
    fmt.Print("notes-cli 1.0\n\n")
    if len(os.Args) == 1 {
        fmt.Println("Error! No arguments!")
        cmd.Help(os.Args[0])
        os.Exit(1)
    }
    
    switch os.Args[1] {
    case "help":
        cmd.Help(os.Args[0])
    case "add":
        if (len(os.Args) < 3) {
            fmt.Println(`Not enough arguments for "add" command!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        cmd.Add(filename, os.Args[2])
    case "list":
        if (len(os.Args) < 3) {
            cmd.List(filename, "all")
        } else {
            if (os.Args[2] != "done" && os.Args[2] != "todo" && os.Args[2] != "in-progress") {
                fmt.Println("Here is only three types of statys! Type help")
                os.Exit(1)
            } 
            cmd.List(filename, os.Args[2])
        }
    case "delete":
        if (len(os.Args) < 3) {
            fmt.Println(`Not enough arguments for "delete" command!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        id, cerr := strconv.Atoi(os.Args[2])
        if cerr != nil {
            fmt.Println(`Argument for "delete" must be a number!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        cmd.Delete(filename, id)
    case "update":
        if (len(os.Args) < 4) {
            fmt.Println(`Not enough arguments for "update" command!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        id, cerr := strconv.Atoi(os.Args[2])
        if cerr != nil {
            fmt.Println(`Argument for "update" must be a number!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        cmd.Update(filename, id, os.Args[3])
    case "mark":
        if (len(os.Args) < 4) {
            fmt.Println(`Not enough arguments for "mark" command!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        id, cerr := strconv.Atoi(os.Args[2])
        if cerr != nil {
            fmt.Println(`Argument for "mark" must be a number!`)
            cmd.Help(os.Args[0])
            os.Exit(1)
        }
        if (os.Args[3] != "done" && os.Args[3] != "todo" && os.Args[3] != "in-progress") {
            fmt.Println("Here is only three types of statuses! Type help")
            os.Exit(1)
        } 
        cmd.Mark(filename, id, os.Args[3])

    default:
        fmt.Println("Error! Here is no this command")
        cmd.Help(os.Args[0])
        os.Exit(1)
    }

}