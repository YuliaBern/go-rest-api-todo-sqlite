package todo

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(task Task) (Task, error) {
	query := `INSERT INTO tasks (title, description, completed) 
	          VALUES (?, ?, ?) 
	          RETURNING id, created_at, updated_at`
	
	var created Task
	err := r.db.QueryRow(query, task.Title, task.Description, task.Completed).
		Scan(&created.ID, &created.CreatedAt, &created.UpdatedAt)
	
	if err != nil {
		return Task{}, err
	}
	
	created.Title = task.Title
	created.Description = task.Description
	created.Completed = task.Completed
	return created, nil
}

func (r *Repository) GetAll() ([]Task, error) {
	rows, err := r.db.Query(`SELECT id, title, description, completed, created_at, updated_at FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}