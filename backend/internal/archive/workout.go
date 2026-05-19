package archive

import (
	"strings"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

type WorkoutEntry struct {
	EntryDate types.DateTime `json:"entryDate"`
	Value     float64        `json:"value"`
}

type WorkoutInfo struct {
	Workout string         `json:"workout"`
	Type    string         `json:"type"`
	Metric  string         `json:"metric"`
	Data    []WorkoutEntry `json:"data"`
}

func getWorkoutData(app core.App, workout *core.Record) ([]WorkoutEntry, error) {
	workoutData := []WorkoutEntry{}

	dataCollection, err := app.FindCollectionByNameOrId("get_workout_entries_graph")
	if err != nil {
		return nil, err
	}

	dataRecords, err := app.FindRecordsByFilter(
		dataCollection,
		"workout = "+workout.GetString("workout"),
		"",
		0,
		0,
	)
	if err != nil {
		return nil, err
	}

	for _, r := range dataRecords {
		workoutData = append(workoutData, WorkoutEntry{EntryDate: r.GetDateTime("created"), Value: r.GetFloat("value")})
	}

	return workoutData, nil
}

func getWorkoutInformation(app core.App) ([]WorkoutInfo, error) {
	workoutCollection, err := app.FindCollectionByNameOrId("get_workout_entries_table")
	if err != nil {
		return nil, err
	}

	workoutRecords, err := app.FindRecordsByFilter(
		workoutCollection,
		"",
		"-workout_type",
		0,
		0,
	)
	if err != nil {
		return nil, err
	}

	workouts := []WorkoutInfo{}

	for _, workout := range workoutRecords {
		data, err := getWorkoutData(app, workout)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, WorkoutInfo{
			Workout: strings.Trim(workout.GetString("workout"), `"`),
			Type:    strings.Trim(workout.GetString("workout_type"), `"`),
			Metric:  strings.Trim(workout.GetString("metric"), `"`),
			Data:    data,
		})
	}

	return workouts, nil
}
