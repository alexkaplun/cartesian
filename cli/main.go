package cli

import (
	"log"
	"os"

	"github.com/alexkaplun/cartesian/model"
	"github.com/alexkaplun/cartesian/util"

	"github.com/alexkaplun/cartesian/api"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

const (
	DEFAULT_PORT = ":8080"
	DATA_FILE    = "data/points.json"
)

func Run(args []string) bool {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	app := cli.NewApp()

	var pointList *model.PointList

	// load the points from a file before starting the service
	before := func(_ *cli.Context) error {
		// parse the points.json file
		var err error
		pointList, err = util.LoadPointListFromCsv(DATA_FILE)
		if err != nil {
			return errors.Wrap(err, "can't parse the points data file")
		}
		log.Printf("Data file loaded successfully, total points: %v", len(pointList.Points()))
		return nil
	}

	app.Commands = cli.Commands{
		{
			Name: "run",
			Subcommands: cli.Commands{
				{
					Name:   "cartesian_api",
					Before: before,
					Action: func(_ *cli.Context) error {
						service := api.New(
							logger,
							pointList,
						)

						service.Run(DEFAULT_PORT)
						return errors.New("cartesian_api service died")
					},
				},
			},
		},
	}

	if err := app.Run(args); err != nil {
		logger.Printf("app failed with error: %v\n", err)
		return false
	}
	return true
}
