package middleware

import (
	"fmt"
	"net/http"
	"os"

	userDomain "test-majoo/src/domain/user"
	helpers "test-majoo/src/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Auth ...
func Auth(a userDomain.Entity) gin.HandlerFunc {
	return func(c *gin.Context) {
		errorParams := map[string]interface{}{}
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			errorParams["meta"] = map[string]interface{}{
				"status":  401,
				"message": "Unauthorized",
			}
			errorParams["code"] = 401
			c.JSON(http.StatusUnauthorized, helpers.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}
		runes := []rune(header)
		tokenString := string(runes[7:])
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			errorParams["meta"] = map[string]interface{}{
				"status":  401,
				"message": "Unauthorized",
			}
			errorParams["code"] = 401
			c.JSON(http.StatusUnauthorized, helpers.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}
		if !token.Valid {
			errorParams["meta"] = map[string]interface{}{
				"status":  401,
				"message": "Unauthorized",
			}
			errorParams["code"] = 401
			c.JSON(http.StatusUnauthorized, helpers.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}

		id := claims["id"].(float64)
		_, err = a.GetUserById(c, int(id))
		fmt.Println(err)
		if err != nil {
			errorParams["meta"] = map[string]interface{}{
				"status":  401,
				"message": "Unauthorized",
			}
			errorParams["code"] = 401
			c.JSON(http.StatusUnauthorized, helpers.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}

		// claims["user_uuid"] = res
		c.Set("User", claims)
		c.Next()

		//End of Auth Service
	}
}

// GetUserCustom ...
func GetUserCustom(c *gin.Context) map[string]interface{} {
	User := c.MustGet("User").(jwt.MapClaims)
	return User
}

// CreateToken ...
func CreateToken(data userDomain.UserLogin) (string, string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = data.ID
	atClaims["data"] = data
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["id"] = data.ID
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rtoken, err := rt.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", "", err
	}
	return token, rtoken, nil
}
