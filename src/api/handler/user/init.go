package userHandler

import (
	userDomain "test-majoo/src/domain/user"
	helpers "test-majoo/src/helper"
	"test-majoo/src/middleware"

	// Mail "github.com/Santara-Engineering-Team/admin-santara-api/src/pkg/mail"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	Entity userDomain.Entity
}

func InitUserHandler(r *gin.RouterGroup, u userDomain.Entity) {
	handler := &AppHandler{
		Entity: u,
	}

	users := r.Group("/user")
	{
		users.GET("/home", handler.Home)
		users.GET("/", middleware.Auth(handler.Entity), handler.GetListUser)
		users.POST("/login", handler.Login)
		users.POST("/", handler.Create)
		users.PUT("/id/:id", handler.Update)
		users.DELETE("/id/:id", handler.Delete)
	}
}

func (a *AppHandler) Home(c *gin.Context) {
	params := map[string]interface{}{
		"payload": gin.H{"message": "OK", "version": "1"},
		"meta":    gin.H{"message": "OK"},
	}
	c.JSON(200, helpers.OutputAPIResponseWithPayload(params))
}
