package main

import (
	_ "time/tzdata"
)

//go:generate go run scripts/static.go

func main() {
	app := NewApplication()
	app.Run()
}
