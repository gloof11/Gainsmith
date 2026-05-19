package archive

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

type ProgramData struct {
	Name  string `json:"name" form:"name"`
	From  string `json:"from" form:"from"`
	To    string `json:"to" form:"to"`
	Notes string `json:"notes" form:"notes"`
}

func HandlerArchive(e *core.RequestEvent) error {
	var data ProgramData

	if err := e.BindBody(&data); err != nil {
		return e.InternalServerError("Failed to read request data", err)
	}

	// Check that the archive request is correct
	if data.Name == "" || data.From == "" || data.To == "" {
		return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid archive data"})
	}

	app := e.App

	weightInformation, err := getWeightInformation(app)
	if err != nil {
		return e.InternalServerError("Could not build weight information", err)
	}

	workoutInformation, err := getWorkoutInformation(app)
	if err != nil {
		return e.InternalServerError("Could not build weight information", err)
	}

	// Create the archive
	id, err := createArchive(app, data, weightInformation, workoutInformation)
	if err != nil {
		return e.InternalServerError("Failed to archive program", err)
	}

	return e.JSON(http.StatusOK, map[string]string{"id": id})
}
