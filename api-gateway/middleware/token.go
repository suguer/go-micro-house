package middleware

import (
	"api-gateway/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		Token := c.Query("token")
		// fmt.Printf("Token: %v\n", Token)
		UserId, err := utils.RDB.Get(c, fmt.Sprintf("session_key:%v", Token)).Result()
		// fmt.Printf("UserId: %v\n", UserId)
		if err == redis.Nil {
			c.JSON(200, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("UserId", UserId)
		c.Next()
	}
}
