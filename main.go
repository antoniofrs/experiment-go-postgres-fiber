package main

import module "github.com/antoniofrs/experiment-go-postgresql/pkg/modules"

func main() {
	module.InitConfigs()
	module.InitPostgres()
	module.InitRoutes()
}
