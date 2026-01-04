package route

import (
	"flag"
	"fmt"
	"go-fiber-minimal/app/console"
	"go-fiber-minimal/config"
	"go-fiber-minimal/lib"
	"os"
)

var Console consoleManager

type consoleManager struct{}

func (consoleManager) Init() {
	if config.Env.APP_ENV == "production" {
		fmt.Println("Cannot access cmd while mode production")
		os.Exit(1)
	}

	dropFlag := flag.Bool("drop", false, "Drop the database tables")
	seedFlag := flag.Bool("seed", false, "Seed the database with initial data")
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")

	flag.Parse()
	status := false

	if *dropFlag {
		console.DB.Drop(lib.DB.Run)
		status = true
	}

	if *migrateFlag {
		console.DB.Migrate(lib.DB.Run)
		status = true
	}

	if *seedFlag {
		console.DB.Seed(lib.DB.Run)
		status = true
	}

	if status {
		os.Exit(0)
	}
}
