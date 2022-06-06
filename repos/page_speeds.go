package repos

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"test-manager/usecase_models"
	models "test-manager/usecase_models/boiler"
)

type PageSpeedRepository interface {
	UpdatePageSpeed(ctx context.Context, PageSpeed models.PageSpeed) error
	GetPageSpeed(ctx context.Context, projectId int) (pageSpeedUseCase []*usecase_models.PageSpeed, err error)
	SavePageSpeed(ctx context.Context, PageSpeed models.PageSpeed) (int, error)
}

type pageSpeedRepository struct {
	db *sql.DB
}

func NewPageSpeedRepository(db *sql.DB) PageSpeedRepository {
	return &pageSpeedRepository{db: db}
}

func (r *pageSpeedRepository) SavePageSpeed(ctx context.Context, pageSpeed models.PageSpeed) (int, error) {
	err := pageSpeed.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return 0, err
	}
	return pageSpeed.ID, nil
}

func (r *pageSpeedRepository) UpdatePageSpeed(ctx context.Context, pageSpeed models.PageSpeed) error {
	_, err := pageSpeed.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r *pageSpeedRepository) GetPageSpeed(ctx context.Context, projectId int) (pageSpeedUseCase []*usecase_models.PageSpeed, err error) {
	pageSpeeds, err := models.PageSpeeds(models.PageSpeedWhere.ProjectID.EQ(projectId)).All(ctx, r.db)
	if err != nil {
		return []*usecase_models.PageSpeed{}, err
	}

	for _, value := range pageSpeeds {
		var pageSpeed usecase_models.PageSpeed
		err := json.Unmarshal([]byte(value.Data.String), &pageSpeed)
		if err != nil {
			log.Println(err.Error())
		}
		pageSpeedUseCase = append(pageSpeedUseCase, &pageSpeed)
	}
	return pageSpeedUseCase, nil
}
