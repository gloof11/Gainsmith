package server

import (
	"log/slog"
	"os"

	"github.com/pocketbase/pocketbase/core"
)

func CreateSuperuser(app core.App) {
	// load the superusers collection
	superusers, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
	if err != nil {
		app.Logger().Warn("Cannot find superusers collection", slog.Any("error", err))
		return
	}

	// only create a superuser if one does not exist
	count, err := app.CountRecords(superusers)
	if err != nil {
		app.Logger().Error("Cannot get count of superusers", slog.Any("error", err))
		os.Exit(1)
	}

	if count > 0 {
		app.Logger().Info("Existing superuser(s) detected")
		return
	}

	// Create the superuser
	email := os.Getenv("SUPERUSER")
	password := os.Getenv("SUPERUSER_PASSWORD")
	if email == "" || password == "" {
		return
	}

	if existing, _ := app.FindAuthRecordByEmail(superusers, email); existing != nil {
		return
	}

	record := core.NewRecord(superusers)
	record.SetEmail(email)
	record.SetPassword(password)
	if err := app.Save(record); err != nil {
		app.Logger().Warn("could not create PB superuser", slog.Any("error", err))
	} else {
		app.Logger().Info("Created superuser", slog.String("email", email))
	}
}
