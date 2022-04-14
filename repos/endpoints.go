package repos

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	models "test-manager/usecase_models/boiler"
)

type EndpointRepository interface {
	UpdateEndpoint(ctx context.Context, endpoint models.Endpoint) error
	GetEndpoint(ctx context.Context, projectId int) (models.Endpoint, error)
	SaveEndpoint(ctx context.Context, endpoint models.Endpoint) (int, error)
}

type endpointRepository struct {
	db *sql.DB
}

func NewEndpointRepository(db *sql.DB) EndpointRepository {
	return &endpointRepository{db: db}
}

func (r *endpointRepository) SaveEndpoint(ctx context.Context, endpoint models.Endpoint) (int, error) {
	err := endpoint.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return 0, err
	}
	return endpoint.ID, nil
}

func (r *endpointRepository) UpdateEndpoint(ctx context.Context, endpoint models.Endpoint) error {
	_, err := endpoint.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r *endpointRepository) GetEndpoint(ctx context.Context, projectId int) (models.Endpoint, error) {
	endpoint, err := models.Endpoints(models.EndpointWhere.ProjectID.EQ(projectId)).One(ctx, r.db)
	if err != nil {
		return models.Endpoint{}, err
	}
	return *endpoint, nil
}
