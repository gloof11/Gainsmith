package archive

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
)

type ArchiveEntry struct {
	WeightInfo  []WeightInfo  `json:"weightInfo"`
	WorkoutInfo []WorkoutInfo `json:"workoutInfo"`
}

func createArchive(app core.App, metadata ProgramData, bodyweight []WeightInfo, workouts []WorkoutInfo) (string, error) {
	collection, err := app.FindCollectionByNameOrId("archived_programs")
	if err != nil {
		return "", err
	}

	entry := ArchiveEntry{
		bodyweight,
		workouts,
	}

	// Create the data JSON object
	data, err := json.Marshal(entry)
	if err != nil {
		return "", err
	}

	// Initalize the new collection
	record := core.NewRecord(collection)

	// Populate the record
	record.Set("name", metadata.Name)
	record.Set("from", metadata.From)
	record.Set("to", metadata.To)
	record.Set("data", data)
	record.Set("notes", metadata.Notes)

	// Save the record
	err = app.Save(record)
	if err != nil {
		return "", err
	}

	return record.Id, nil
}
