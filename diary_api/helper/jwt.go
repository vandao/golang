package helper

import (
	"diary_api/model"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user model.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(context *gin.Context) (jwt.MapClaims, error) {
	token, err := getToken(context)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token provided")
}

func CurrentUser(context *gin.Context) (model.User, error) {
	claims, err := ValidateJWT(context)
	if err != nil {
		return model.User{}, err
	}

	userId := uint(claims["id"].(float64))
	user, err := model.FindUserById(userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func getTokenFromRequest(context *gin.Context) string {
	bearToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return privateKey, nil
	})
	return token, err
}
