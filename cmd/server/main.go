package main

import (
	"context"
	"fmt"

	"github.com/HimanS-sys/go-rest-api-demo/internal/db"
)

// Run - responsible for
// initialization and startup
// of our go application
func Run() error {
	fmt.Println("Starting our application...")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Successfully connected and pinged database")
	return nil
}
func main() {
	fmt.Println("Go Rest-API Demo.")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
