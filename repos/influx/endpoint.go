package influx

import (
	"context"
	"database/sql"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"strconv"
	"time"
)

type EndpointReportRepository interface {
	WriteEndpointReport(ctx context.Context, projectId int, pipelineId int, success int, responseTime float64) error
	ReadEndpointReportByProject(ctx context.Context, projectId int, pipelineId int, timeFrame string, fields []string) (err error, res []interface{})
	ReadEndpointReportByPipeline(ctx context.Context, projectId int, pipelineId int) (err error, res []interface{})
}

type endpointReportRepository struct {
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI

	db *sql.DB
}

func NewEndpointReportRepository(writeAPI api.WriteAPIBlocking, queryAPI api.QueryAPI, db *sql.DB) EndpointReportRepository {
	return &endpointReportRepository{writeAPI: writeAPI, queryAPI: queryAPI, db: db}
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

func (r *endpointReportRepository) ReadEndpointReportByProject(ctx context.Context, projectId int, pipelineId int, timeFrame string, fields []string) (err error, res []interface{}) {
	fieldsQuery := ""
	for _, value := range fields {
		fieldsQuery = fieldsQuery + fmt.Sprintf("or r[\"_field\" == \"%s\"", value)
	}
	query := ""
	if fieldsQuery == "" {
		fieldsQuery = fieldsQuery[2:]
		query = fmt.Sprintf(`from(bucket:"my-bucket")
		 |> range(start: -%s) 
		|> filter(fn: (r) => r["_measurement"] == "endpoint")
		|> filter(fn: (r) => %s)
		|> filter(fn: (r) => r.project_id == "%s")
		|> aggregateWindow(every: 10s, fn: last, createEmpty: false)
		|> yield(name: "last")`, timeFrame, fieldsQuery, strconv.Itoa(projectId))
	} else {
		query = fmt.Sprintf(`from(bucket:"my-bucket")
		 |> range(start: -%s) 
		|> filter(fn: (r) => r["_measurement"] == "endpoint")
		|> filter(fn: (r) => r.project_id == "%s")
		|> aggregateWindow(every: 10s, fn: last, createEmpty: false)
		|> yield(name: "last")`, timeFrame, strconv.Itoa(projectId))
	}

	//query := fmt.Sprintf(`from(bucket:"my-bucket")
	//  |> range(start: -%s)
	//|> filter(fn: (r) => r["_measurement"] == "endpoint")
	//|> filter(fn: (r) => r["_field"] == "success" or r["_field"] == "response_time")
	//|> filter(fn: (r) => r.project_id == "%s")
	//|> aggregateWindow(every: 10s, fn: last, createEmpty: false)
	//|> yield(name: "last")`, timeFrame, strconv.Itoa(projectId))
	fmt.Println(query)
	result, err := r.queryAPI.Query(context.Background(), query)

	if err == nil {
		for result.Next() {
			fmt.Printf("value: %v\n", result.Record().Value())
			res = append(res, result.Record().Value())
		}
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
			return err, res
		}
	} else {
		return err, res
	}
	return nil, res
}

func (r *endpointReportRepository) ReadEndpointReportByPipeline(ctx context.Context, projectId int, pipelineId int) (err error, res []interface{}) {
	result, err := r.queryAPI.Query(context.Background(), `from(bucket:"my-bucket")
    |> range(start: -1h) 
    |> filter(fn: (r) => r._measurement == "stat")`)

	if err == nil {
		for result.Next() {
			fmt.Printf("value: %v\n", result.Record().Value())
			res = append(res, result.Record().Value())
		}
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
			return err, res
		}
	} else {
		return err, res
	}
	return nil, res
}
