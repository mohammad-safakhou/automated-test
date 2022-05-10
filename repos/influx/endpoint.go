package influx

import (
	"context"
	"database/sql"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"strconv"
	"time"
)

type EndpointReportRepository interface {
	WriteEndpointReport(ctx context.Context, projectId int, pipelineId int, success int, responseTime float64) error
}

type endpointReportRepository struct {
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI

	db *sql.DB
}

func NewEndpointReportRepository(db *sql.DB) EndpointReportRepository {
	return &endpointReportRepository{db: db}
}

func (r *endpointReportRepository) WriteEndpointReport(ctx context.Context, projectId int, pipelineId int, success int, responseTime float64) error {
	p := influxdb2.NewPoint("endpoint",
		map[string]string{"project_id": strconv.Itoa(projectId), "endpoint_id": strconv.Itoa(pipelineId)},
		map[string]interface{}{"success": success, "response_time": responseTime},
		time.Now())
	err := r.writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}
