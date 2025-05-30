package repository

import (
	"database/sql"
	"github.com/koushik/task-management-system/pkg/domain"
)


type TaskRepository interface {
	Create(task *domain.Task) error
	Update(task *domain.Task) error
	Delete(task *domain.Task) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *domain.Task) error {
	query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, task.Title, task.Description, task.Status).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (r *taskRepository) Update(task *domain.Task) error {
	query := `UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=NOW() WHERE id=$4`
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.ID)
	return err
}

func (r *taskRepository) Delete(task *domain.Task) error {
	query := `UPDATE tasks SET is_active=FALSE, updated_at=NOW() WHERE id=$1`
	_, err := r.db.Exec(query,task.ID)
	return err
}

