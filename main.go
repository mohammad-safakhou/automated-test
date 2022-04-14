package main

//go:generate sqlboiler --wipe --no-tests psql -o models/boiler

import "test-manager/cmd"

func main() {
	cmd.Execute()
}
