package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bpradana/devcamp_pe_bintangpradanaerlanggaputra/cmd/db"
	"github.com/bpradana/devcamp_pe_bintangpradanaerlanggaputra/pkg/products"
	"github.com/bpradana/devcamp_pe_bintangpradanaerlanggaputra/pkg/variants"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file, err: ", err.Error())
	}

	// Connect to database
	db, err := db.ConnectDB()
	if err != nil {
		log.Println("error connecting to database, err: ", err.Error())
	}

	// Create a new instance of Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	productRepository := products.NewRepository(db)
	variantRepository := variants.NewRepository(db)
	productService := products.NewUsecase(productRepository, variantRepository)
	variantService := variants.NewUsecase(variantRepository)

	v1 := e.Group("/api/v1")
	products.NewHandler(v1, productService)
	variants.NewHandler(v1, variantService)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
