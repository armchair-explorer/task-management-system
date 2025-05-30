package service

import (
	"github.com/koushik/task-management-system/pkg/domain"
	"github.com/koushik/task-management-system/pkg/repository"
)


type TaskService interface {
	CreateTask(task *domain.Task) error
	UpdateTask(task *domain.Task) error
	DeleteTask(task *domain.Task) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *domain.Task) error {
	return s.repo.Create(task)
}


func (s *taskService) UpdateTask(task *domain.Task) error {
	return s.repo.Update(task)
}

func (s *taskService) DeleteTask(task *domain.Task) error {
        return s.repo.Delete(task)
}
