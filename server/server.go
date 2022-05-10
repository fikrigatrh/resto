package server

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

}
