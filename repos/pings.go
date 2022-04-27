package repos

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	models "test-manager/usecase_models/boiler"
)

type PingRepository interface {
	UpdatePing(ctx context.Context, Ping models.Ping) error
	GetPing(ctx context.Context, projectId int) (models.Ping, error)
	SavePing(ctx context.Context, Ping models.Ping) (int, error)
}

type pingRepository struct {
	db *sql.DB
}

func NewPingRepository(db *sql.DB) PingRepository {
	return &pingRepository{db: db}
}

func (r *pingRepository) SavePing(ctx context.Context, ping models.Ping) (int, error) {
	err := ping.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return 0, err
	}
	return ping.ID, nil
}

func (r *pingRepository) UpdatePing(ctx context.Context, ping models.Ping) error {
	_, err := ping.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r *pingRepository) GetPing(ctx context.Context, projectId int) (models.Ping, error) {
	ping, err := models.Pings(models.PingWhere.ProjectID.EQ(projectId)).One(ctx, r.db)
	if err != nil {
		return models.Ping{}, err
	}
	return *ping, nil
}
