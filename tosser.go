package main

import (
	"path/filepath"
)

type Application struct {
}

func IsNetmail(name string) bool {
	var ext string = filepath.Ext(name)
	return ext == ".pkt"
}

func IsArchmail(name string) bool {
	var ext string = filepath.Ext(name)
	return ext != ".pkt"
}

func (self *Application) Run() {
	self.ProcessInbound()
	self.ProcessOutbound()
}

func Tosser() {
	app := new(Application)
	app.Run()
}

