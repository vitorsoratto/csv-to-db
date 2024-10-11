package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vitorsoratto/csv-to-db/config"
	"github.com/vitorsoratto/csv-to-db/schemas"
)

var csvPath string

func execute() {
	var err error

	fmt.Println("Reading CSV file...")
	if csvPath == "" {
		fmt.Println("Please provide a CSV file path")
		return
	}
	plans, workouts, exercises, err := schemas.ReadCSV(csvPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = schemas.InitSchema()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Inserting data into database...")
	schemas.InsertWorkouts(workouts)
	schemas.InsertExercises(exercises)
	schemas.InsertPlans(plans)
	fmt.Println("Done!")
}

func CreateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "csvtodb",
		Short: "A tool to import Workouts CSV into the SQLite database",
		Long:  "A tool to import Workouts CSV into the SQLite database",
		Run: func(cmd *cobra.Command, args []string) {
		  if config.DBPath == "" || csvPath == "" {
		    fmt.Println("Please provide a database path and a CSV file path")
		    cmd.Usage()
		    return
		  }
			execute()
		},
	}

	command.PersistentFlags().StringVarP(&config.DBPath, "db", "d", "", "Path to the SQLite database")
	command.PersistentFlags().StringVar(&csvPath, "csv", "", "Path to the CSV file")

	return command
}
