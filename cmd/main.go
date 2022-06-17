package main

import (
	"gateway/internal/adapters/app"
	"gateway/internal/adapters/core"
	"gateway/internal/adapters/framework/left"
	"gateway/internal/adapters/framework/right"
)

func main() {
	db := right.NewDBAdapter("dbname", "db@password")
	coreComs := app.NewCoreToOutsideAdapter(db)
	core := core.NewInAdapter(coreComs)
	app := app.NewAppAdapter(core, coreComs)
	svr := left.NewHttpAdapter(app)
	svr.StartListening(":9000")
}
