package main

import (
	"credit_system/auth_service/graph"
	"credit_system/auth_service/helpers"
	"credit_system/core/handlers"
	payment_graph "credit_system/credit_payment_service/payment_graph"
	customer_graph "credit_system/customer_management/customer_graph"
	"log"

	"credit_system/auth_service/repositories"
	"credit_system/auth_service/resolvers"
	"credit_system/auth_service/services"
	"credit_system/core/schema"

	customer_repo "credit_system/customer_management/customer_repositories"
	customer_resolver "credit_system/customer_management/customer_resolvers"
	customer_service "credit_system/customer_management/customer_services"

	payment_repo "credit_system/credit_payment_service/payment_repositories"
	payment_resolver "credit_system/credit_payment_service/payment_resolvers"
	payment_service "credit_system/credit_payment_service/payment_services"

	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db, err := helpers.GetGormDB()
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authResolver := resolvers.NewAuthResolver(authService)

	// Initialize Customer Module
	customerRepository := customer_repo.NewCustomerRepository(db)
	customerService := customer_service.NewCustomerService(customerRepository)
	customerResolver := customer_resolver.NewCustomerResolver(customerService)

	// Initialize Payment Module
	paymentRepository := payment_repo.NewPaymentSubmissionRepository(db)
	paymentService := payment_service.NewPaymentService(paymentRepository)
	paymentResolver := payment_resolver.NewPaymentResolver(paymentService)

	// Build GraphQL Schema
	queryFields := schema.MergeFields(
		graph.AuthQueries(authResolver),
		customer_graph.CustomerQueries(customerResolver),
		payment_graph.NewPaymentQueryType(paymentResolver),
	)

	mutationFields := schema.MergeFields(
		graph.AuthMutations(authResolver),
		customer_graph.CustomerMutations(customerResolver),
		payment_graph.NewPaymentMutationType(paymentResolver),
	)

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: queryFields,
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: mutationFields,
	})

	// Initialize schema with error handling
	schema.InitSchema(queryType, mutationType)
	schema.InitMiddleware(authService)

	startServer()
}

func startServer() {
	e := echo.New()

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			// "http://localhost:8081",
			// "http://10.10.9.197:8081",
			"*",
		},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	// Add request logging
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// GraphQL endpoint
	e.POST("/graphql", handlers.Handler)

	e.POST("/upload_screenshot", handlers.SaveImageToLocal)

	e.Logger.Fatal(e.Start(":8080"))
}
