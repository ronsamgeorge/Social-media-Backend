package main

import (
	"fmt"
	"log"

	"github/ronsamgeorge/Social-media-Backend/internal/database"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database ensured!")
}