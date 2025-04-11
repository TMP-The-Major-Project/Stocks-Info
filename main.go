package main

import (
	"log"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// User represents a user in the system
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest represents the registration request body
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ReturnsRequest represents the returns calculation request
type ReturnsRequest struct {
	InitialInvestment   float64 `json:"initialInvestment"`
	MonthlyContribution float64 `json:"monthlyContribution"`
	Years               int     `json:"years"`
	AnnualReturn        float64 `json:"annualReturn"`
}

// ReturnsResponse represents the returns calculation response
type ReturnsResponse struct {
	TotalInvestment float64 `json:"totalInvestment"`
	TotalReturns    float64 `json:"totalReturns"`
	FinalAmount     float64 `json:"finalAmount"`
}

// UserPreferences represents a user's investment preferences
type UserPreferences struct {
	FirstName         string   `json:"firstname"`
	LastName          string   `json:"lastname"`
	Email             string   `json:"email"`
	ContactNumber     string   `json:"contactnumber"`
	Gender            string   `json:"gender"`
	Address           string   `json:"address"`
	Occupation        string   `json:"occupation"`
	Income            string   `json:"income"`
	InvestmentReason  string   `json:"investmentreason"`
}

// In-memory storage (replace with database in production)
var users = make(map[string]User)

// In-memory storage for user preferences
var userPreferences = make(map[string]UserPreferences)

func init() {
	// Add default admin user
	users["admin"] = User{
		Username: "admin",
		Password: "admin",
		Email:    "admin@stockmarket.com",
	}
}

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Stock Market API v1.0",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// Serve static files from the src directory
	app.Static("/", "./src")

	// API routes
	api := app.Group("/api")

	// Authentication routes
	api.Post("/register", handleRegister)
	api.Post("/login", handleLogin)

	// Returns calculator
	api.Post("/calculate-returns", handleCalculateReturns)

	// Add the personalInfo route
	api.Post("/personalInfo", handlePersonalInfo)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

func handleRegister(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if username already exists
	if _, exists := users[req.Username]; exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Username already exists",
		})
	}

	// Directly convert req (RegisterRequest) to User
	user := User(req)
	users[req.Username] = user

	log.Printf("User registered: %+v", users[req.Username]) // Debug log

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func handleLogin(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Debug logging
	log.Printf("Login attempt for user: %s", req.Username)
	log.Printf("Stored users: %+v", users)

	// Check if user exists and password matches
	user, exists := users[req.Username]
	if !exists {
		log.Printf("User not found: %s", req.Username)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	if user.Password != req.Password {
		log.Printf("Password mismatch for user: %s", req.Username)
		log.Printf("Received password: %s", req.Password)
		log.Printf("Stored password: %s", user.Password)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	// In production, generate and return a JWT token
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func handleCalculateReturns(c *fiber.Ctx) error {
	var req ReturnsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Calculate monthly rate
	monthlyRate := req.AnnualReturn / 12 / 100
	months := req.Years * 12

	// Calculate future value with compound interest
	futureValue := req.InitialInvestment * math.Pow(1+monthlyRate, float64(months))

	// Calculate future value of monthly contributions
	monthlyFutureValue := req.MonthlyContribution * (math.Pow(1+monthlyRate, float64(months)) - 1) / monthlyRate

	// Total future value
	totalFutureValue := futureValue + monthlyFutureValue

	// Calculate total investment
	totalInvestment := req.InitialInvestment + (req.MonthlyContribution * float64(months))

	// Calculate total returns
	totalReturns := totalFutureValue - totalInvestment

	response := ReturnsResponse{
		TotalInvestment: totalInvestment,
		TotalReturns:    totalReturns,
		FinalAmount:     totalFutureValue,
	}

	return c.JSON(response)
}

// handlePersonalInfo handles the personal information form submission
func handlePersonalInfo(c *fiber.Ctx) error {
	var preferences UserPreferences
	if err := c.BodyParser(&preferences); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Store preferences in memory using email as key
	userPreferences[preferences.Email] = preferences

	return c.JSON(fiber.Map{
		"message": "Preferences saved successfully",
	})
}
