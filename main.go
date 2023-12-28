package main

import module "github.com/antoniofrs/experiment-go-postgresql/pkg/module"

func main() {
	module.InitConfigs()
	module.InitPostgres()
	module.InitRoutes()
}
