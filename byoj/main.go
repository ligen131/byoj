package main

import (
	"byoj/model"
	"byoj/shared/server"
	"byoj/shared/yamlconfig"
)

func main() {
	configuration, err := yamlconfig.ConfigLoad("config.yml")
	if err != nil {
		panic(err)
	}

	err = model.Connect(configuration.Database)
	if err != nil {
		panic(err)
	}

	err = model.InitModel()
	if err != nil {
		panic(err)
	}

	err = server.Run(configuration.Server)
	if err != nil {
		panic(err)
	}
}
