package main

import (
	"fmt"
	"log"
	"os"
	database "user/database"
	handler "user/handler"
	repository "user/repository"
	service "user/service"

	"github.com/joho/godotenv"
)

// App encapsulates the application dependencies.
type App struct {
	Port        string
}

func main() {
	// Load environment variables
	envConfig, err := loadEnv()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Initialize dependencies
	db, err := database.NewDatabaseConnection(true)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(service.UserService{Repository: userRepo})
	
	uh := handler.NewUserHandler(userService)
	uh.Handlers()

	app := App{
		Port:      envConfig["PORT"],
	}

	// Initialize the server
	app.startServer()
}

// startServer initializes the server and starts listening for requests.
func (app *App) startServer() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)
	server := NewServer(app.Port, logger)
	defer func() {
		if err := server.Shutdown(); err != nil {
			logger.Printf("Error during shutdown: %v", err)
		}
	}()
	
	err := server.Start()
	if err != nil {
		logger.Printf("Error during startup: %v", err)
	}

	// Wait for the shutdown signal
	<-server.ShutdownChannel()
	fmt.Println("Application has shut down.")
}

func loadEnv() (envMap map[string]string, err error) {
	envFile, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}
	return envFile, nil
}