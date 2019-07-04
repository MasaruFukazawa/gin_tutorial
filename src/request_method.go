package main

import "github.com/gin-gonic/gin"

func main() {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "someGet",
		})
	})

	router.POST("/somePost", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "somePost",
		})
	})

	router.PUT("/somePut", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "somePut",
		})
	})

	router.DELETE("/someDelete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "someDelete",
		})
	})

	router.PATCH("/somePatch", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "somePatch",
		})
	})

	router.HEAD("/someHead", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "someHead",
		})
	})

	router.OPTIONS("/someOptions", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "someOptions",
		})
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
