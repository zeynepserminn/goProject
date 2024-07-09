package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"goProject/internal/core/handler"
	"goProject/internal/core/services"
	"goProject/internal/core/validation"
	"goProject/pkg/db/repositories"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the server",
	Long:  "Runs the server for your application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the server")
		StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func StartServer() {

	router := gin.Default()
	v := validator.New()

	validation.RegisterValidation(v)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("maxlen", validation.ValidateMaxLen)
		v.RegisterValidation("emailcheck", validation.ValidateEmail)
		v.RegisterValidation("validId", validation.ValidateID)
	}

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.POST("/users", userHandler.AddUser)
	router.GET("/users/all", userHandler.GetAllUsers)
	router.PUT("/users", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	fmt.Println("Starting the server")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

}
