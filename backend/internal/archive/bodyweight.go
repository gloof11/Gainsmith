package archive

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

type WeightInfo struct {
	Date   types.DateTime `json:"date"`
	Weight float64        `json:"weight"`
}

func getWeightInformation(app core.App) ([]WeightInfo, error) {
	weightCollection, err := app.FindCollectionByNameOrId("weight_entries")
	if err != nil {
		return nil, err
	}

	weightRecords, err := app.FindRecordsByFilter(
		weightCollection,
		"",
		"date",
		0,
		0,
	)
	if err != nil {
		return nil, err
	}

	bodyWeights := []WeightInfo{}

	for _, v := range weightRecords {
		bodyWeights = append(bodyWeights, WeightInfo{Date: v.GetDateTime("date"), Weight: v.GetFloat("weight")})
	}

	return bodyWeights, nil
}
