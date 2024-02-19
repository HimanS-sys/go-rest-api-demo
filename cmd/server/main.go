package main

import "fmt"

// Run - responsible for 
// initialization and startup
// of our go application 
func Run() error {
	fmt.Println("Starting our application...")
	return nil
}
func main() {
	fmt.Println("Go Rest-API Demo.")
	if err := Run(); err!=nil {
		fmt.Println(err)
	}
}