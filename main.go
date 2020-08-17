package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vela/types/yaml"
	"net/http"
	"strings"
)

func main() {
	router := gin.Default()
	router.Use(authMiddleware("217758e80e8db1b597d2c52740a35580"))
	router.POST("/config", config)

	router.Run(":8080")
}

// authMiddleware ensures the authorization header is present and matches
func authMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the token from the `Authorization` header
		authToken := c.GetHeader("Authorization")
		if len(authToken) > 0 {
			if strings.Contains(authToken, "Bearer") {
				authToken = strings.Split(authToken, "Bearer ")[1]
			}
		}

		if authToken != token {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"result":  "error",
				"message": "not authorized",
			})
			return
		}

		c.Next()
	}
}

func config(c *gin.Context) {
	var payload yaml.Build
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var steps []*yaml.Step
	steps = append(steps, payload.Steps...)
	steps = append(steps, &yaml.Step{
			Name: "config",
			Image: "alpine",
			Commands: []string{"echo hello from config"},
	})
	payload.Steps = steps

	c.JSON(http.StatusOK, payload)
}