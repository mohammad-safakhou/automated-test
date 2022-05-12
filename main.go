package main

//go:generate sqlboiler --wipe --no-tests psql -o usecase_models/boiler

import (
	"test-manager/cmd"
)

func main() {
	cmd.Execute()
}
