package usecase_models

import (
	"github.com/volatiletech/null/v8"
	"time"
)

type Datacenter struct {
	ID             int       `json:"id"`
	Baseurl        string    `json:"baseurl"`
	Title          string    `json:"title"`
	ConnectionRate null.Int  `json:"connection_rate"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
	DeletedAt      null.Time `json:"deleted_at"`
}

type CreateDatacenterResponse struct {
	DatacenterId int `json:"datacenter_id"`
}
