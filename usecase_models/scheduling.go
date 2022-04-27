package usecase_models

type Scheduling struct {
	ProjectId      int    `json:"project_id"`
	Duration       int    `json:"duration"`
	EndAt          string `json:"end_at"`
	DataCentersIds []int  `json:"data_centers"` // datacenter id
}