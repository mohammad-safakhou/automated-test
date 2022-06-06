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

type PingRepository interface {
	UpdatePing(ctx context.Context, Ping models.Ping) error
	GetPing(ctx context.Context, projectId int) (pingUseCase []*usecase_models.Pings, err error)
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

func (r *pingRepository) GetPing(ctx context.Context, projectId int) (pingUseCase []*usecase_models.Pings, err error) {
	pings, err := models.Pings(models.PingWhere.ProjectID.EQ(projectId)).All(ctx, r.db)
	if err != nil {
		return []*usecase_models.Pings{}, err
	}

	for _, value := range pings {
		var ping usecase_models.Pings
		err := json.Unmarshal([]byte(value.Data.String), &ping)
		if err != nil {
			log.Println(err.Error())
		}
		pingUseCase = append(pingUseCase, &ping)
	}
	return pingUseCase, nil
}
