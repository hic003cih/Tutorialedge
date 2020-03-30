package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading Environment Variable")
	var databasePass string
	databasePass = os.Getenv("DATABASE_PASS")
	fmt.Printf("Database Password: %s\n", databasePass)

	err := os.Setenv("DATABASE_PASS", "newunicorns222")
	if err != nil {
		fmt.Println(err)
	}
	databasePass = os.Getenv("DATABASE_PASS")
	fmt.Printf("Database Password: %s\n", databasePass)
}
