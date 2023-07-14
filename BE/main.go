package main

import (
	"fmt"
	"net"

	adminRoutes "github.com/Ashutosh-730/easy-pay/BE/admins"
	customerRoutes "github.com/Ashutosh-730/easy-pay/BE/customers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Enable CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Update with your frontend domain
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	// Admin panel routes
	admin := router.Group("/admin")
	{
		adminRoutes.AdminRoutes(admin)
	}

	// Customer page routes
	customer := router.Group("/customer")
	{
		customerRoutes.CustomerRoutes(customer)
	}

	runRouter(router)

}

func runRouter(router *gin.Engine) {
	// if port 8080 is not available then use any other port
	port := 8080
	for {
		addr := fmt.Sprintf(":%d", port)
		ln, err := net.Listen("tcp", addr)
		if err == nil {
			// Port is available
			ln.Close()
			break
		}
		port++ // Try the next port
	}

	fmt.Printf("\n Available port: http://localhost:%d\n \n", port)
	// fmt.Printf("Server Address:  http://localhost:%s\n", cast.ToString(port))

	// Start server with available port
	addr := fmt.Sprintf(":%d", port)

	// // Set the environment variable
	// err := os.Setenv("BACKEND_PORT", cast.ToString(port))
	// if err != nil {
	// 	fmt.Println("Failed to set environment variable:", err)
	// 	return
	// }

	router.Run(addr)
}
