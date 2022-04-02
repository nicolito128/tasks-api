package tasks

import (
	"log"

	"github.com/nicolito128/tasks-api/pkg/database"
)

// GetTasks() return all tasks saved in the database
func GetTasks() []Task {
	query := `SELECT * FROM tasks`
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	list := []Task{}
	for rows.Next() {
		curTask := Task{}
		err = rows.Scan(&curTask.ID, &curTask.Name, &curTask.Content)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, curTask)
	}

	return list
}

// 	CreateTask() insert a new Task in the database.
func CreateTask(t Task) error {
	query := `
			INSERT INTO 
				tasks (name, content)
				VALUES ($1, $2)
	`
	return database.Request(query, t.Name, t.Content)
}

// DeleteTaskById() delete a Task in the database by id.
func DeleteTaskById(id int) error {
	query := `DELETE FROM tasks WHERE id = ($1)`
	return database.Request(query, id)
}

// UpdateTask() update task in the database with a new task.
func UpdateTask(t Task) error {
	query := `
			UPDATE tasks
				SET name=($2),
				content=($3)
				WHERE id=($1);
	`
	return database.Request(query, t.ID, t.Name, t.Content)
}
