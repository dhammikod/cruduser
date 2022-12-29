package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dhammikod/cruduser/initializers"
	"github.com/dhammikod/cruduser/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	//get the cookie ff req
	tokenString, err := c.Cookie("authorization")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Error found",
		})
	} else {
		//decode/validate cookie
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Error found",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//check expired
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Token Expired",
				})
			}

			//find user with token
			var user models.User
			initializers.DB.First(&user, claims["sub"])

			if user.ID == 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Invalid Token Credentials",
				})
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			//attach to req
			c.Set("user", user)
			//continue
			c.Next()
			fmt.Println(claims["foo"], claims["nbf"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "No Token Found",
			})
		}
	}

}
