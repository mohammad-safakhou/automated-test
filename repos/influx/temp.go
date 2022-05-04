package influx

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func Temp() {
	url := "http://localhost:8086"
	token := "I0B3yPC-XsAPDT35wxQgKX6GOKFIPaGdfz8t8kfELkqHjZz-yszZ_IaSyBTRMj5Ew8cijrOjp8_KU9VfNz71Tw=="
	org := "test"
	bucket := "my-bucket"

	client := influxdb2.NewClient(url, token)

	writeAPI := client.WriteAPIBlocking(org, bucket)
	queryAPI := client.QueryAPI(org)

	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	writeAPI.WritePoint(context.Background(), p)
	client.Close()

	result, err := queryAPI.Query(context.Background(), `from(bucket:"my-bucket")
    |> range(start: -1h) 
    |> filter(fn: (r) => r._measurement == "stat")`)

	if err == nil {
		for result.Next() {
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}