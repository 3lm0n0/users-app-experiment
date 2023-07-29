package main

import (
	"fmt"
	"log"
	database "user/database"
	handler "user/handler"
	repository "user/repository"
	service "user/service"

	"github.com/joho/godotenv"
)


func main() {
	// godotenv package to get .env variables.
	envFile, err := loadEnv()
	if err != nil {
		log.Fatal(err)
		log.Fatalf("Error loading environment variables file")
		return
	}
	// config data.
	port := envFile["port"]

	// Instantiates the database
	db, err := database.NewDatabaseConnection(true)
	if err != nil {
		log.Fatal(err)
		log.Fatalf("Error connecting to database")
		return
	}
	
	// Initialize the user repository.
	ru := repository.NewUserRepository(db)

	// Initialize the user service params.
	sus := service.UserService{
		Repository: ru,
	}

	// Initialize the user service.
	us := service.NewUserService(sus)

	// Initialize the user handlers.
	uh := handler.NewUserHandler(us)
	uh.Handlers()

	// Initialize the server.
	server := NewServer(port)
	fmt.Println("server runnig ar port", port)

	// log.
	log.Fatal(server.Start())
}

func loadEnv() (envMap map[string]string, err error) {
	envFile, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}
	return envFile, nil
}