# Task Manager CLI

A simple yet powerful command-line tool for managing your tasks.

## Description

Task Manager CLI is a lightweight, cross-platform application that helps you keep track of your tasks and stay organized. You can add, list, update, delete, and get details of your tasks right from your terminal.

## Project Url
https://github.com/Dacostasolo/task-manager-cli
project source url [Roadmap.sh](https://roadmap.sh/projects/task-tracker)

## Getting Started

### Prerequisites

- Go (version 1.15 or higher)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/Dacostasolo/task-manager-cli.git
   ```
2. Navigate to the project directory:
   ```sh
   cd task-manager-cli
   ```
3. Build the application:
   ```sh
   go build -o task-cli ./cmd
   ```

## Usage

You can run the application using the following command:

```sh
./task-cli [command] [options]
```

## Commands

The following commands are available:

- `add`: Add a new task
- `list`: List all tasks
- `update`: Update an existing task
- `delete`: Delete a task
- `get`: Get details of a specific task
- `save`: Save tasks to persistent storage

### Options

- `-t`, `-title`: Title of the task (required for `add`)
- `-d`, `-description`: Description of the task (optional for `add`)
- `-s`, `-status`: Status of the task (0 = todo, 1 = in-progress, 2 = done) (optional for `add`/`update`)
- `-id`: ID of the task (required for `update`, `delete`, `get`)
- `-f`, `-filter`: Filter tasks by status when listing (optional for `list`)

### Examples

- Add a new task:
  ```sh
  ./task-cli add -t "Buy groceries" -d "Milk, Bread, Eggs" -s "Pending"
  ```
- List all tasks:
  ```sh
  ./task-cli list
  ```
- Update a task:
  ```sh
  ./task-cli update -id 123456 -s "Completed"
  ```
- Delete a task:
  ```sh
  ./task-cli delete -id 123456
  ```
- Get details of a task:
  ```sh
  ./task-cli get -id 123456
  ```
- Save all tasks:
  ```sh
  ./task-cli save
  ```
