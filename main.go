package main

import (
	"UserLoginSystem/config"
	controllers "UserLoginSystem/controller"
	middlewares "UserLoginSystem/middleware"
	"UserLoginSystem/model"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.InitDB()
	if err := model.Migrate(config.DB); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login/username", controllers.LoginByUsername)
		v1.POST("/send-otp/email", controllers.SendOTPByEmail)
		v1.POST("/login/email", controllers.LoginByEmail)
		v1.POST("/send-otp/phone", controllers.SendOTPByPhone)
		v1.POST("/login/phone", controllers.LoginByPhone)
		auth := v1.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			auth.PUT("/change-password", controllers.ChangePassword)
		}
	}
	if err := r.Run(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
