package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	auth2 "goProject/internal/core/handler/auth"
	"goProject/internal/core/handler/middleware"
	"goProject/internal/core/handler/user"
	"goProject/internal/core/model"
	"goProject/internal/core/services/auth"
	"goProject/internal/core/services/userservice"
	"goProject/pkg/jwt"
	"goProject/pkg/postgres"
	"goProject/pkg/postgres/repositories"
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

	db, err := postgres.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Connected to database")
	}

	router := gin.Default()

	userService := userservice.NewUserService(repositories.NewUserRepository(db.GetInstance()))
	tokenConfig := jwt.NewJwt()

	authService := auth.NewAuthService(repositories.NewUserRepository(db.GetInstance()), tokenConfig)
	authHandler := auth2.NewAuthHandler(authService)
	userHandler := user.NewUserHandler(userService)
	refreshTokenHandler := auth2.NewRefreshToken(tokenConfig)

	authRouter := router.Group("/auth")
	authRouter.POST("/login", authHandler.Login)
	authRouter.POST("/refresh", refreshTokenHandler.RefreshAccessToken)

	userRouter := router.Group("/user")
	userRouter.Use(middleware.AuthMiddleware(userService))

	adminRouter := userRouter.Group("/")
	adminRouter.Use(middleware.RoleAuth(model.RoleAdmin))

	adminRouter.POST("/", userHandler.AddUser)
	adminRouter.PUT("/:id", userHandler.UpdateUser)
	adminRouter.DELETE("/:id", userHandler.DeleteUser)

	userRouter.GET("/", userHandler.GetAllUsers)
	userRouter.GET("/:id", userHandler.GetUserByID)
	userRouter.PUT("/profile", userHandler.UpdateProfile)
	userRouter.PUT("/password", userHandler.UpdatePassword)
	userRouter.GET("/profile", userHandler.GetProfile)

	fmt.Println("Starting the server")

	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

}
