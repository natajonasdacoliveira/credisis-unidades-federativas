package tools

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)


func CreateToken(userid uint64) (string, error) {
  
  atClaims := jwt.MapClaims{}
  atClaims["user_id"] = userid

  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

  token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  if err != nil {
     return "", err
  }

  return token, nil
}