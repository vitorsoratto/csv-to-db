package schemas

import "github.com/vitorsoratto/csv-to-db/util"

type Exercise struct {
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	Repetitions   int     `json:"repetitions"`
	AnimationMale string  `json:"animation_male"`
	Instructions  *string `json:"instructions"`
	ID            int     `json:"id"`
	WorkoutID     int     `json:"workout_id"`
}

func InsertExercises(data []Exercise) error {
	defer util.Timer("InsertExercises")()
	for _, exercise := range data {
		if err := db.Create(&exercise).Error; err != nil {
			return err
		}
	}

	return nil
}
