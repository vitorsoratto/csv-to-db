package schemas

import "github.com/vitorsoratto/csv-to-db/util"

type Workout struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	PlanID int    `json:"plan_id"`
}

func InsertWorkouts(data map[int]Workout) error {
	defer util.Timer("InsertWorkouts")()
	for _, workout := range data {
		if err := db.Create(&workout).Error; err != nil {
			return err
		}
	}

	return nil
}
