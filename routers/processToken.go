package routers

import (
	"errors"
	"strings"

	"github.com/Tsuryu/tiwttor/db/userdao"
	"github.com/Tsuryu/tiwttor/model"
	"github.com/dgrijalva/jwt-go"
)

// Email : Logged in user email
var Email string

// UserID : Logged in user id
var UserID string

// ProcessToken : validates token
func ProcessToken(token string) (*model.Claim, bool, string, error) {
	myKey := []byte("mySecretKey")
	claims := &model.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, exists, _ := userdao.FindBy(claims.Email)
		if exists {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, exists, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}

	return claims, false, string(""), err
}
