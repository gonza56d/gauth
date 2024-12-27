package app

import (
	"os"
	"time"

	apimodel "github.com/gonza56d/gauth/pkg"

	"github.com/golang-jwt/jwt/v5"
)

// Log in user and return a JWT token if successful.
func Login(request *apimodel.AuthRequest) string {
	var authenticated bool = login(request)
	if !authenticated {
		return ""
	}
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": request.Email,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}
	storeJWT(request.Email, tokenString)
	return tokenString
}

// Signup user and return true if successful or false if email is already taken.
func SignUp(request *apimodel.AuthRequest) bool {
	if isEmailTaken(request.Email) {
		return false
	}
	signUp(request)
	return true
}
