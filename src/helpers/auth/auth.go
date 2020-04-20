package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tohidplus/url_project/src/app/models"
	"github.com/tohidplus/url_project/src/database"
	"github.com/tohidplus/url_project/src/exception"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

//Todo:: move to environment variable

var jwtKey = []byte("ede83948343d7d306c3d21895b2ddedd")

type Token struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Claims struct {
	User models.User `json:"user"`
	jwt.StandardClaims
}

func AttemptLogin(request *http.Request) (bool, Token) {
	user := models.User{}
	database.DB.Where("email = ?", request.PostFormValue("email")).First(&user)
	if user.ID == 0 {
		return false, Token{}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.PostFormValue("password"))); err != nil {
		return false, Token{}
	}
	token, _ := GetToken(user)
	return true, token
}

func GetToken(user models.User) (Token, error) {
	expiredAt := time.Now().Add(24 * time.Hour)
	claims := getClaim(user, expiredAt)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString(jwtKey)
	if err != nil {
		return Token{}, err
	}
	token := Token{
		Token:     tokenString,
		ExpiredAt: expiredAt,
	}
	return token, nil
}

func GetUser(request *http.Request) (bool, models.User) {

	tknStr := request.Header.Get("Authorization")
	if tknStr == "" {
		return false, models.User{}
	}
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		exception.LogPrint(err)
		return false, models.User{}
	}
	if !tkn.Valid {
		return false, models.User{}
	}
	return true, claims.User
}

func getClaim(user models.User, expiredAt time.Time) *Claims {
	return &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
		},
	}
}
