package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mohnaofal/go-clean-architecture/app/handler"
	"github.com/mohnaofal/go-clean-architecture/app/repository"
	"github.com/mohnaofal/go-clean-architecture/app/services"
	"github.com/mohnaofal/go-clean-architecture/config"
	"github.com/mohnaofal/go-clean-architecture/migration/migrate"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Init Configuration
	cfg := config.InitConfig()

	// Migration Table
	migrate.AutoMigration(cfg.DB().GormMysql())

	// init framework echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	productsRepo := repository.NewProductRepository(cfg.DB().GormMysql())
	productsService := services.NewProducts(productsRepo)
	productsHandler := handler.NewProductsHandler(productsService)
	productsGroups := e.Group("v1/products")
	productsHandler.Mount(productsGroups)

	if err := e.Start(fmt.Sprintf(`:%d`, cfg.PORT())); err != nil {
		e.Logger.Fatal(err)
	}

}
