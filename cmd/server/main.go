package main

import (
	//"context"
	"fmt"

	"github.com/HimanS-sys/go-rest-api-demo/internal/db"
)

// Run - responsible for
// initialization and startup
// of our go application
func Run() error {
	fmt.Println("starting our application...")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	fmt.Println("successfully connected and migrated database")
	return nil
}
func main() {
	fmt.Println("Go Rest-API Demo")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
