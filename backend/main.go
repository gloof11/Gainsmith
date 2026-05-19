package main

import (
	"log"

	"gainsmith/cmd/server"
	"gainsmith/internal/archive"
	_ "gainsmith/migrations"
	"gainsmith/ui"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/osutils"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: osutils.IsProbablyGoRun(),
	})

	// Bootstrap on serve
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Create the Pocketbase Superuser
		server.CreateSuperuser(se.App)

		// register archive program route
		se.Router.POST("/api/archive/create", archive.HandlerArchive)

		// Serve the frontend
		se.Router.GET("/{path...}", apis.Static(ui.DistDirFS, true))

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
