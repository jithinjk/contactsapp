package contacts

import (
	"log"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetRequestID add request ID
func GetRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := uuid.NewV4()
		if err != nil {
			log.Println(err)
		}
		c.Writer.Header().Set("X-Request-ID", uid.String())
		c.Next()
	}
}
