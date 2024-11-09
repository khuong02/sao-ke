package server

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ipSubnetMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		allowed := false
		for _, subnet := range allowedSubnets {
			_, cidrNet, _ := net.ParseCIDR(subnet)
			if cidrNet.Contains(net.ParseIP(clientIP)) {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Your IP is not allowed"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (s *Server) recoverPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError,
					gin.H{"message": "Internal server error"})
				c.Abort()
			}
		}()
		c.Next()
	}
}
