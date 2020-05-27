package jwt

import (
	"time"

	"github.com/Tsuryu/tiwttor/model"
	"github.com/dgrijalva/jwt-go"
)

// CreateJWT : creates a jwt to manage the season
func CreateJWT(user model.User) (string, error) {
	myKey := []byte("mySecretKey")

	payload := jwt.MapClaims{
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"birth_date": user.BirthDate,
		"website":    user.WebSite,
		"_id":        user.ID.Hex(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
		// "biography": user.biography,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(myKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
