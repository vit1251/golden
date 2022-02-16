package main

//go:generate go run scripts/static.go

func main() {
	app := NewApplication()
	app.Run()
}
