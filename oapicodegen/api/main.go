package main

import (
	"fmt"
	"oapicodegen/lib"
	"oapicodegen/petstore"

	"github.com/labstack/echo/v4"
)

type Server struct{}

func main() {
	e := echo.New()

	lib.SAY()
	fmt.Println(petstore.Pet{})
	e.Logger.Fatal(e.Start(":8080"))
}
