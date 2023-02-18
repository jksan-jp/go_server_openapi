package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// CreatePets - Create a pet
	e.POST("/v1/pets", c.CreatePets)

	// GetPetsPetID - 
	e.GET("/v1/pets/:petID", c.GetPetsPetID)

	// ListPets - List all pets
	e.GET("/v1/pets", c.ListPets)

	// PatchPetsPetID - 
	e.PATCH("/v1/pets/:petID", c.PatchPetsPetID)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}