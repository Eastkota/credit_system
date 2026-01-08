package main

import (
	schema "credit_system/auth_service/graph"
	"credit_system/auth_service/handlers"
	"credit_system/auth_service/helpers"
	"credit_system/auth_service/repositories"
	"credit_system/auth_service/resolvers"
	"credit_system/auth_service/services"

	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := helpers.GetGormDB()
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	resolver := resolvers.NewAuthResolver(authService)

	mutationType := schema.NewMutationType(resolver)
	queryType := schema.NewQueryType(resolver)

	schema.InitSchema(queryType, mutationType)
	schema.InitMiddleware(authService)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://10.10.8.203:3000",
		},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	e.POST("/graphql", handlers.Handler)
	e.Logger.Fatal(e.Start(":8090"))
}
