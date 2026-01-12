package main

import (
	"credit_system/auth_service/graph"
	"credit_system/auth_service/helpers"
	"credit_system/core/handlers"
	customer_graph "credit_system/customer_management/customer_graph"
	"fmt"

	"credit_system/auth_service/repositories"
	"credit_system/auth_service/resolvers"
	"credit_system/auth_service/services"
	"credit_system/core/schema"

	customer_repo "credit_system/customer_management/customer_repositories"
	customer_resolver "credit_system/customer_management/customer_resolvers"
	customer_service "credit_system/customer_management/customer_services"

	"log"

	"github.com/graphql-go/graphql"
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
	authResolver := resolvers.NewAuthResolver(authService)

	customerRepository := customer_repo.NewCustomerRepository(db)
	fmt.Println("Customer repository initialized:", customerRepository)
	customerService := customer_service.NewCustomerService(customerRepository)
	customerResolver := customer_resolver.NewCustomerResolver(customerService)
	fmt.Println("Customer resolver initialized:", customerResolver)

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: schema.MergeFields(
			graph.AuthMutations(authResolver),
			customer_graph.CustomerMutations(customerResolver),
			// graph.CreditMutations(creditResolver),
		),
	})
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: schema.MergeFields(
			graph.AuthQueries(authResolver),
			customer_graph.CustomerQueries(customerResolver),
			// graph.CreditQueries(creditResolver),
		),
	})

	schema.InitSchema(queryType, mutationType)
	fmt.Println("GraphQL schema initialized")
	graph.InitMiddleware(authService)

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
