package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Generate token from userID with user permissions
func GenerateToken(userID interface{}) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID

	//secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(config.SecretKey)
}

// Check is token passed on request is valid
func ValidateToken(r *http.Request) error {
	tokenString, err := extractToken(r)
	if err != nil {
		return err
	}

	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

// Extract token from bearer token
func extractToken(r *http.Request) (string, error) {
	tokenWithBearer := strings.Split(r.Header.Get("Authorization"), " ")

	if len(tokenWithBearer) == 2 {
		return tokenWithBearer[1], nil
	}

	return "", errors.New("invalid token, require Bearer prefix")
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("method of subscription inexpected %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

// Extract ID from jwt token
func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString, err := extractToken(r)
	if err != nil {
		return 0, err
	}

	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%0.f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New(("invalid token"))
}
