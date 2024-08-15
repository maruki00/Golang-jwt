package repositories

import (
	"Golang-jwt/internal/core"
	"Golang-jwt/internal/dtos"
	"Golang-jwt/internal/models"
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserRepository struct {
}

var secretKey = []byte("golang-jwt")

func getPass(pass string) string {
	hash := md5.Sum([]byte(pass))
	return fmt.Sprintf("%x", hash)
}

func (obj *UserRepository) Login(email, password string) (*dtos.AuthDTO, error) {
	auth := &dtos.AuthDTO{}
	db := core.GetDB()
	defer db.Close()
	statement, err := db.Prepare("select id, email,fullname  from users where email=? and password = ?")
	if err != nil {
		return nil, errors.New(err.Error())
	}
	err = statement.QueryRow(email, getPass(password)).Scan(&auth.Id, &auth.Email, &auth.Fullname)
	if err != nil {
		return nil, errors.New("could not login db error  " + err.Error())
	}
	auth.Token, err = createToken(auth.Email, auth.Fullname)
	if err != nil {
		panic("could not create token ")
	}

	statement, _ = db.Prepare("delete from auths where user_id = ?")
	_, _ = statement.Exec(auth.Id)

	statement, err = db.Prepare("INSERT INTO auths (token, email, fullname,user_id) VALUES (?,?,?,?)")
	if err != nil {
		return nil, errors.New("could not make auth record" + err.Error())
	}
	_, err = statement.Exec(auth.Token, auth.Email, auth.Fullname, auth.Id)
	if err != nil {
		return nil, errors.New("something wrong " + err.Error())
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

func (l *UserRepository) Register(email, fullname, address, password string) (*dtos.RegisterDTO, error) {

	db := core.GetDB()
	statement, err := db.Prepare("INSERT INTO users (email, fullname, address ,password) VALUES (?,?,?,?)")
	if err != nil {
		return nil, errors.New("invalid data  ")
	}

	_, err = statement.Exec(email, fullname, address, getPass(password))
	if err != nil {
		return nil, errors.New("invalid data  ")
	}

	return nil, nil
}

func (l *UserRepository) GetUsers(page, offset int) ([]*models.UserModel, error) {
	items := make([]*models.UserModel, 0)
	db := core.GetDB()
	if page < 0 {
		page = 0
	}
	pag := (page * offset) + offset
	statement, err := db.Prepare("SELECT id, email, fullname, address from users limit ? ")
	if err != nil {
		return nil, errors.New("something wrong")
	}

	var u models.UserModel
	res, _ := statement.Query(pag)

	for res.Next() {
		err = res.Scan(&u.Id, &u.Email, &u.FullName, &u.Address)
		if err == nil {
			items = append(items, &u)
		}
	}

	return items, nil
}
