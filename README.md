# Task Tracker CLI

A simple command-line application built in Golang to manage tasks. This application allows you to add, update, delete, list, and mark the status of tasks. Tasks are stored in a JSON file for persistence.
Also supported multi-user, by using SUID bit. (maybe unsafe, but i don't think so). JSON file stored in /etc/tasks.json

Roadmap Project Challenge: https://roadmap.sh/projects/task-tracker

## Features

- **Add Task**: Create a new task with a description.
- **Update Task**: Modify the description of an existing task.
- **Delete Task**: Remove a task from the list.
- **Mark Task Status**: Mark a task as in progress, done, or todo.
- **List Tasks**: View all tasks.

## Installation

Before installation you MUST(!) install Golang compiler

```bash
# Clone the repository
$ git clone https://github.com/izya4ka/notes-cli.git

# Navigate to the project directory
$ cd notes-cli
```
You have two ways of installing

### Root (support for multi-user)

```bash

# Build and install the application
$ sudo make install

# Run the application
$ notes-cli
```

### Rootless

before rootless installation please change distination file in main.go (because can't create file in /etc directory without root priveleges)

```bash
# Build application
$ go build

# Run application
$ ./notes-cli
```

### Uninstall
```bash
$ sudo make uninstall
```

## Usage
Here’s how you can use the various commands:

### Add a Task
```bash
./notes-cli add "Your task description"
```

### Update a Task
```bash
./notes-cli update <task_id> "Updated task description"
```

### Delete a Task
```bash
./notes-cli delete <task_id>
```

### Marking a task
```bash
./notes-cli mark <task_id> in-progress/done/todo
```

### Listing all tasks
```bash
./notes-cli list
```

### Listing all done/in-progress/todo tasks
```bash
./notes-cli list done/in-progress/todo
```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request for review.
