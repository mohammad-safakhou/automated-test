package main

//go:generate sqlboiler --wipe --no-tests psql -o usecase_models/boiler

import (
	"test-manager/cmd"
	"test-manager/repos/influx"
)

func main() {
	influx.Temp()
	return
	cmd.Execute()
}
