package main

import (
	"net/http"
	"oapicodegen/lib"
	"oapicodegen/petstore"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type petstoreHandler struct{}

type HelloWorld struct {
	Message string `json:"message"`
}

// List all pets
// (GET /pets)
func (si petstoreHandler) ListPets(ctx echo.Context, params petstore.ListPetsParams) error {
	return ctx.JSON(http.StatusOK, HelloWorld{
		Message: "oapi-codegen",
	})
}

// Create a pet
// (POST /pets)
func (si petstoreHandler) CreatePets(ctx echo.Context) error {
	return nil
}

// (GET /pets/{petID})
func (si petstoreHandler) GetPetsPetID(ctx echo.Context, petID string) error {
	return nil
}

// (PATCH /pets/{petID})
func (si petstoreHandler) PatchPetsPetID(ctx echo.Context, petID string) error {
	return nil
}

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	petstoreHandler := petstoreHandler{}
	petstore.RegisterHandlers(e, petstoreHandler)

	lib.SAY()

	e.Logger.Fatal(e.Start(":8080"))
}
