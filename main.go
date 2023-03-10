package main

import (
	"fmt"
	"io"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var privateKeyBytes = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`

var privateKeyBytes2 = []byte(privateKeyBytes)

var privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes2)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func getMachineToken(c *gin.Context) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["machine_id"] = "random id"
	claims["serial_number"] = "12345"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(privateKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		c.JSON(http.StatusBadRequest, "Error")
		return
	}

	response := TokenResponse{
		AccessToken:  tokenString,
		RefreshToken: "random token",
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.Discard
	router := gin.Default()
	router.POST("/machine", getMachineToken)
	router.Run("localhost:8888")
}
