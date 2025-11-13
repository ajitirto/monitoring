package main

import (
	"monitoring/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	internalServerError = "Internal Server Error"
	notFound            = "Not Found"
)

func main() {
	r := gin.Default()

	middleware.PrometheusInit()

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Use(middleware.TrackMetrics())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Prometheus!")
	})

	r.GET("/get-user", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") 

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Users")
	})

	r.GET("/get-role", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") 

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Roles")
	})

	r.GET("/get-level", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") 

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Levels")
	})

	r.Run(":8080")
}
