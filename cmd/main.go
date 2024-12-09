package main

import (
	"20241209/infra"
	"20241209/routes"
	"flag"
	"log"
)

func main() {
	migrateDb, seedDb := readFlags()

	ctx, err := infra.NewServiceContext(migrateDb, seedDb)
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	if shouldNotLaunchServer(migrateDb, seedDb) {
		return
	}

	r := routes.NewRoutes(*ctx)

	if err = r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func readFlags() (bool, bool) {
	migrateDb := flag.Bool("m", false, "use this flag to migrate database")
	seedDb := flag.Bool("s", false, "use this flag to seed database")
	flag.Parse()
	return *migrateDb, *seedDb
}

func shouldNotLaunchServer(migrateDb bool, seedDb bool) bool {
	if migrateDb {
		return true
	}

	if seedDb {
		return true
	}

	return false
}
