package productHandler

import (
	productDomain "test-majoo/src/domain/product"
	userDomain "test-majoo/src/domain/user"
	helpers "test-majoo/src/helper"
	"test-majoo/src/middleware"

	// Mail "github.com/Santara-Engineering-Team/admin-santara-api/src/pkg/mail"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	Product productDomain.Entity
	User    userDomain.Entity
}

func InitProductHandler(r *gin.RouterGroup, p productDomain.Entity, u userDomain.Entity) {
	handler := &AppHandler{
		Product: p,
		User:    u,
	}

	merchants := r.Group("/merchant")
	{
		merchants.GET("/", middleware.Auth(handler.User), handler.GetListMerchantByUserId)
		merchants.GET("/id/:id", middleware.Auth(handler.User), handler.GetMerchantById)
		merchants.POST("/", middleware.Auth(handler.User), handler.CreateMerchant)
	}

	outlets := r.Group("/outlet")
	{
		outlets.GET("/", middleware.Auth(handler.User), handler.GetListOutletByMerchantId)
		outlets.GET("/id/:id", middleware.Auth(handler.User), handler.GetOutletById)
		outlets.POST("/", middleware.Auth(handler.User), handler.CreateOutlet)
	}

	products := r.Group("/product")
	{
		products.GET("/", middleware.Auth(handler.User), handler.GetListProductByOutletId)
		products.GET("/id/:id", middleware.Auth(handler.User), handler.GetProductById)
		products.POST("/", middleware.Auth(handler.User), handler.CreateProduct)
		products.PUT("/id/:id", middleware.Auth(handler.User), handler.UpdateProduct)
		products.DELETE("/id/:id", middleware.Auth(handler.User), handler.DeleteProduct)
	}
}

func (a *AppHandler) Home(c *gin.Context) {
	params := map[string]interface{}{
		"payload": gin.H{"message": "OK", "version": "1"},
		"meta":    gin.H{"message": "OK"},
	}
	c.JSON(200, helpers.OutputAPIResponseWithPayload(params))
}
