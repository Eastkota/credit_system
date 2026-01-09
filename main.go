package main

import (
	"credit_system/auth_service/graph"
	"credit_system/auth_service/handlers"
	"credit_system/auth_service/helpers"
	customer_graph "credit_system/customer_management/customer_graph"

	"credit_system/auth_service/repositories"
	"credit_system/auth_service/resolvers"
	"credit_system/auth_service/services"
	"credit_system/core/schema"

	customer_repo "credit_system/customer_management/customer_repositories"
	customer_resolver "credit_system/customer_management/customer_resolvers"
	customer_service "credit_system/customer_management/customer_services"

	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/graphql-go/graphql"
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

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: schema.MergeFields(
			graph.AuthMutations(authResolver),
			// graph.UserMutations(userResolver),
			// graph.CreditMutations(creditResolver),
		),
	})
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: schema.MergeFields(
			graph.AuthQueries(authResolver),
			// graph.UserQueries(userResolver),
			// graph.CreditQueries(creditResolver),
		),
	})

	customerRepository := customer_repo.NewCustomerRepository(db)
	customerService := customer_service.NewCustomerService(customerRepository)
	customerResolver := customer_resolver.NewCustomerResolver(customerService)

	customerMutationType := customer_graph.NewCustomerMutationType(customerResolver)
	customerQueryType := customer_graph.NewCustomerQueryType(customerResolver)

	schema.InitSchema(queryType, mutationType)
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
