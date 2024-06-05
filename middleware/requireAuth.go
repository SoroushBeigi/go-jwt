package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/SoroushBeigi/go-jwt/initializers"
	"github.com/SoroushBeigi/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("In middleware!")
	providedToken := c.Request.Header.Get("Authorization")
	if providedToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized:": "Please provide a token"})
	}

	providedToken = providedToken[7:]
	print(providedToken)
	token, err := jwt.Parse(providedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil

	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized:": "Token invalid,"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized:": "Token Expired"})
		}
		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized:": "User not found"})
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized:": "Token invalid"})
	}

}
