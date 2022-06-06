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

type EndpointRepository interface {
	UpdateEndpoint(ctx context.Context, endpoint models.Endpoint) error
	GetEndpoint(ctx context.Context, projectId int) (endpointsUseCase []*usecase_models.Endpoints, err error)
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

func (r *endpointRepository) GetEndpoint(ctx context.Context, projectId int) (endpointsUseCase []*usecase_models.Endpoints, err error) {
	endpoints, err := models.Endpoints(models.EndpointWhere.ProjectID.EQ(projectId)).All(ctx, r.db)
	if err != nil {
		return []*usecase_models.Endpoints{}, err
	}

	for _, value := range endpoints {
		var endpoint usecase_models.Endpoints
		err := json.Unmarshal([]byte(value.Data.String), &endpoint)
		if err != nil {
			log.Println(err.Error())
		}
		endpointsUseCase = append(endpointsUseCase, &endpoint)
	}
	return endpointsUseCase, nil
}
