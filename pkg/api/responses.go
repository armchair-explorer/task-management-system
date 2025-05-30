package api

import "github.com/koushik/task-management-system/pkg/domain"

type TaskResponse struct {
	Task domain.Task `json:"task"`
}

type TasksResponse struct {
	Tasks []domain.Task `json:"tasks"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
	Total int           `json:"total"`
}

type PaginatedTasks struct {
    Tasks []domain.Task `json:"tasks"`
    Page  int           `json:"page"`
    Limit int           `json:"limit"`
}
