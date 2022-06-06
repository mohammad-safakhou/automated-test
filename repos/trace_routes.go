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

type TraceRouteRepository interface {
	UpdateTraceRoute(ctx context.Context, TraceRoute models.TraceRoute) error
	GetTraceRoute(ctx context.Context, projectId int) (traceRouteUseCase []*usecase_models.TraceRoutes, err error)
	SaveTraceRoute(ctx context.Context, TraceRoute models.TraceRoute) (int, error)
}

type traceRouteRepository struct {
	db *sql.DB
}

func NewTraceRouteRepository(db *sql.DB) TraceRouteRepository {
	return &traceRouteRepository{db: db}
}

func (r *traceRouteRepository) SaveTraceRoute(ctx context.Context, traceRoute models.TraceRoute) (int, error) {
	err := traceRoute.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return 0, err
	}
	return traceRoute.ID, nil
}

func (r traceRouteRepository) UpdateTraceRoute(ctx context.Context, traceRoute models.TraceRoute) error {
	_, err := traceRoute.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r *traceRouteRepository) GetTraceRoute(ctx context.Context, projectId int) (traceRouteUseCase []*usecase_models.TraceRoutes, err error) {
	traceRoutes, err := models.TraceRoutes(models.TraceRouteWhere.ProjectID.EQ(projectId)).All(ctx, r.db)
	if err != nil {
		return []*usecase_models.TraceRoutes{}, err
	}

	for _, value := range traceRoutes {
		var traceRoute usecase_models.TraceRoutes
		err := json.Unmarshal([]byte(value.Data.String), &traceRoute)
		if err != nil {
			log.Println(err.Error())
		}
		traceRouteUseCase = append(traceRouteUseCase, &traceRoute)
	}
	return traceRouteUseCase, nil
}
