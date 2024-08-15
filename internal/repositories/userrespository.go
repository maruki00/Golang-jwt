package repositories

import (
	"Golang-jwt/internal/core"
	"Golang-jwt/internal/dtos"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserRepository struct {
}

var secretKey = []byte("golang-jwt")

func (obj *UserRepository) Login(email, password string) (*dtos.AuthDTO, error) {
	auth := &dtos.AuthDTO{}
	db := core.GetDB()
	defer db.Close()
	//hash := md5.Sum([]byte(password))
	//h := fmt.Sprintf("%x", hash)
	h := password
	statement, err := db.Prepare("select id, email,fullname  from users where email=? and password = ?")
	fmt.Println("data : ", password, email)
	if err != nil {
		return nil, errors.New("invalid credentails " + err.Error())
	}
	err = statement.QueryRow(email, h).Scan(&auth.Id, &auth.Email, &auth.Fullname)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("invalid creadentails " + err.Error())
	}
	auth.Token, err = createToken(auth.Email, auth.Fullname)
	if err != nil {
		panic("could not create token ")
	}
	return auth, nil
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
