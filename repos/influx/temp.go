package influx

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func Temp() {
	url := "http://localhost:8086"
	token := "EEPjW1onUpQlhPy5bAL-SwQkE_AkI57KY4RBtNak13qk5ODuhjH9zuabMMGy7GPBhZw383eplNXoy3j5HfpArg=="
	org := "test"
	bucket := "my-bucket"

	client := influxdb2.NewClient(url, token)

	writeAPI := client.WriteAPIBlocking(org, bucket)
	queryAPI := client.QueryAPI(org)

	p := influxdb2.NewPoint("endpoint",
		map[string]string{"project_id": "1"},
		map[string]interface{}{"success": 1},
		time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Println(err.Error())
	}
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
