package main

import (
	"os"

	"github.com/optique-dev/core"
	"github.com/optique-dev/template/config"
)

// @title Optique application TO CHANGE
// @version 1.0
// @description This is a sample application
// @contact.name Courtcircuits
// @contact.url https://github.com/Courtcircuits
// @contact.email tristan-mihai.radulescu@etu.umontpellier.fr
func main() {
	conf, err := config.LoadConfig()
	core.Init() // init logger

	if err != nil {
		config.HandleError(err)
	}
	cycle := NewCycle()

	if conf.Bootstrap {
		err := cycle.Setup()
		if err != nil {
			core.Error(err.Error())
			cycle.Stop()
			os.Exit(1)
		}
	}

	err = cycle.Ignite()
}
