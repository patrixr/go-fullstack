package core

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "myapp/internal/migrations"
)

type App struct {
	*pocketbase.PocketBase

	// @TODO Extend the main app struct with more fields as needed
}

// CreateApp is the main entry point for the app
// It initializes the app and configures it
// You may add more configurations here as needed
func CreateApp() *App {
	env := ReadEnv("ENV", "dev")

	base := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDev:     true,
		DefaultDataDir: "./.data/pb_data." + env + ".db",
	})

	app := &App{
		base,
	}

	migratecmd.MustRegister(base, base.RootCmd, migratecmd.Config{
		Automigrate: env == "dev",
		Dir:         "internal/migrations",
	})

	base.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		return MountRoutes(app, e.Router)
	})

	return app
}
