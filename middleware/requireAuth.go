package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("in middleware")
	//get the cookie ff req
	//decode/validate cookie
	//check expired
	//find user with token
	//attach to req
	//continue

	c.Next()
}
