package model

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AuthClaims struct {
	User   int       `json:"user"`
	UID    uuid.UUID `json:"uuid"`
	Random string    `json:"rnd"`
	jwt.RegisteredClaims
}

func ValidateAccessToken(accessToken string) (bool, error) {
	_, _, err := ExtractUserIDFromToken(accessToken)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractUserIDFromToken(accessToken string) (int, uuid.UUID, error) {
	authClaims := &AuthClaims{}
	token, err := jwt.ParseWithClaims(accessToken, authClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ZWEB_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, uuid.Nil, err
	}

	claims, ok := token.Claims.(*AuthClaims)
	if !(ok && token.Valid) {
		return 0, uuid.Nil, err
	}

	return claims.User, claims.UID, nil
}

func CreateAccessToken(id int, uid uuid.UUID) (string, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(10000))

	claims := &AuthClaims{
		User:   id,
		UID:    uid,
		Random: vCode,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "ZWEB",
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 24 * 7),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(os.Getenv("ZWEB_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
