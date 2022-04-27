package repos

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	models "test-manager/usecase_models/boiler"
)

type TraceRouteRepository interface {
	UpdateTraceRoute(ctx context.Context, TraceRoute models.TraceRoute) error
	GetTraceRoute(ctx context.Context, projectId int) (models.TraceRoute, error)
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

func (r *traceRouteRepository) GetTraceRoute(ctx context.Context, projectId int) (models.TraceRoute, error) {
	traceRoute, err := models.TraceRoutes(models.TraceRouteWhere.ProjectID.EQ(projectId)).One(ctx, r.db)
	if err != nil {
		return models.TraceRoute{}, err
	}
	return *traceRoute, nil
}
