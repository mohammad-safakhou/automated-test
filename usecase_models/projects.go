package usecase_models

import (
	"time"
)

type Project struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	IsActive  bool      `json:"is_active"`
	Notifications
	ExpireAt  time.Time `json:"expire_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CreateProjectResponse struct {
	ProjectId int `json:"project_id"`
}

type Notifications struct {

}
