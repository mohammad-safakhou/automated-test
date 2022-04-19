package repos

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	models "test-manager/usecase_models/boiler"
)

type DataCentersRepository interface {
	UpdateDataCenters(ctx context.Context, dataCenter models.Datacenter) error
	GetDataCenters(ctx context.Context, id int) (models.Datacenter, error)
	SaveDataCenters(ctx context.Context, dataCenter models.Datacenter) (int, error)
}

type dataCentersRepository struct {
	db *sql.DB
}

func NewDataCentersRepositoryRepository(db *sql.DB) DataCentersRepository {
	return &dataCentersRepository{db: db}
}

func (r *dataCentersRepository) SaveDataCenters(ctx context.Context, dataCenter models.Datacenter) (int, error) {
	err := dataCenter.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return 0, err
	}
	return dataCenter.ID, nil
}

func (r *dataCentersRepository) UpdateDataCenters(ctx context.Context, dataCenter models.Datacenter) error {
	_, err := dataCenter.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r *dataCentersRepository) GetDataCenters(ctx context.Context, id int) (models.Datacenter, error) {
	datacenter, err := models.Datacenters(models.DatacenterWhere.ID.EQ(id)).One(ctx, r.db)
	if err != nil {
		return models.Datacenter{}, err
	}
	return *datacenter, nil
}
