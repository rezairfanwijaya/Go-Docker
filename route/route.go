package route

import (
	"go-docker/handler"
	"go-docker/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, r *gin.Engine) {
	// user
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewHandlerUser(userService)

	v1 := r.Group("v1")
	v1.POST("/login", userHandler.Login)
	v1.GET("/users", userHandler.GetAllUser)
}
