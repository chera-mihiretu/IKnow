package main

import (
	"log"

	"github.com/chera-mihiretu/IKnow/delivery/controller"
	"github.com/chera-mihiretu/IKnow/delivery/router"
	"github.com/chera-mihiretu/IKnow/infrastructure/mongodb"
	"github.com/chera-mihiretu/IKnow/infrastructure/storage"
	"github.com/chera-mihiretu/IKnow/repository"
	"github.com/chera-mihiretu/IKnow/usecases"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	GothSetup()
	client, err := mongodb.NewMongoClient()

	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	// S3 client setup
	s3Client, err := storage.TestS3Client()
	if err != nil {
		log.Fatal("Failed to connect to S3:", err)
	}
	// Ensure the MongoDB client is closed when the application exits
	myDatabase := client.Database("lazyme")
	// storage dependencies

	blackBlaze := storage.NewBlackblazeStorage(s3Client)
	storageRepository := repository.NewStorageRepository(blackBlaze)
	storageUseCase := usecases.NewStorageUseCase(storageRepository)

	// auth dependecies
	authRepository := repository.NewAuthRepository(myDatabase)
	authUseCase := usecases.NewAuthUseCase(authRepository)
	AuthController := controller.NewAuthController(authUseCase)
	// user dependencies
	userRepository := repository.NewUserRepository(myDatabase)
	userUseCase := usecases.NewUserUseCase(userRepository)
	// follow dependencies
	connectionRepository := repository.NewConnectRepository(myDatabase)
	connectionUsecase := usecases.NewConnectUsecase(connectionRepository)
	connectionController := controller.NewConnectController(connectionUsecase, userRepository)
	// department dependencies
	departmentRepository := repository.NewDepartmentRepository(myDatabase)
	departmentUsecase := usecases.NewDepartmentUseCase(departmentRepository)
	// post dependencies
	postRepository := repository.NewPostRepository(myDatabase,
		departmentRepository,
		connectionRepository,
		*userRepository)
	postUseCase := usecases.NewPostUseCase(postRepository)
	PostController := controller.NewPostController(postUseCase, userUseCase, departmentUsecase, storageUseCase)

	router := router.SetupRoutes(
		AuthController,
		PostController,
		connectionController,
	)

	if err := router.Run(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
