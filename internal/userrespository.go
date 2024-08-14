package internal

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserRepository struct {
}

var secretKey = []byte("secret-key")

func (obj *UserRepository) Login(email, password string) (error, *UserModel) {
	db := GetDB()
	defer db.Close()
	hash := md5.Sum([]byte(password))
	h := fmt.Sprintf("%x", hash)
	statement, err := db.Prepare("select * from users where email=? and password = ?")
	if err != nil {
		panic("invalid credentails")
	}
	rows := statement.QueryRow(email, h).Scan()

}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
