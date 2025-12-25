package main

import (
	"log"
	"os"
	"path/filepath"

	jimo "github.com/jimo-go/framework"
	"github.com/jimo-go/framework/database"
	"github.com/jimo-go/jimo/app/models"
	"github.com/jimo-go/jimo/routes"
)

func main() {
	app := jimo.New()
	app.Views(filepath.Join("resources", "views"))
	app.MustWeb()

	db := database.NewMemoryConnection()
	database.Use(db)
	if err := models.SeedUsers(); err != nil {
		log.Fatal(err)
	}
	routes.Web(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	addr := ":" + port

	log.Printf("JIMO app running at http://localhost%s\n", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
