package internal

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserRepository struct {
}

var secretKey = []byte("secret-key")

func (obj *UserRepository) Login(email, password string) (*AuthDTO, error) {
	auth := &AuthDTO{}
	db := GetDB()
	defer db.Close()
	hash := md5.Sum([]byte(password))
	h := fmt.Sprintf("%x", hash)
	statement, err := db.Prepare("select id, email,fullname  from users where email=? and password = ?")
	if err != nil {
		panic("invalid credentails")
	}
	err = statement.QueryRow(email, h).Scan(&auth.Id, &auth.Email, &auth.Fullname)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("invalid creadentails " + err.Error())
	}

	auth.Token, err = createToken(auth.Email, auth.Fullname)

}

func createToken(email, fullname string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":    email,
			"fullname": fullname,
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
