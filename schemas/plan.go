package schemas

import "github.com/vitorsoratto/csv-to-db/util"

type Plan struct {
	Name  string `json:"name"`
	Level string `json:"level"`
	ID    int    `json:"id"`
}

func InsertPlans(data map[int]Plan) error {
	defer util.Timer("InsertPlans")()
	for _, plan := range data {
		if err := db.Create(&plan).Error; err != nil {
			return err
		}
	}

	return nil
}
