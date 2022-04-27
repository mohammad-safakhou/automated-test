package repos

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	models "test-manager/usecase_models/boiler"
)

type PageSpeedRepository interface {
	UpdatePageSpeed(ctx context.Context, PageSpeed models.PageSpeed) error
	GetPageSpeed(ctx context.Context, projectId int) (models.PageSpeed, error)
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

func (r *pageSpeedRepository) GetPageSpeed(ctx context.Context, projectId int) (models.PageSpeed, error) {
	pageSpeed, err := models.PageSpeeds(models.PageSpeedWhere.ProjectID.EQ(projectId)).One(ctx, r.db)
	if err != nil {
		return models.PageSpeed{}, err
	}
	return *pageSpeed, nil
}
