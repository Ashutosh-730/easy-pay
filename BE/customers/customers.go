// customer_routes.go

package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(customer *gin.RouterGroup) {
	customer.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Customer profile"})
	})

	// Add more customer page routes here
}
