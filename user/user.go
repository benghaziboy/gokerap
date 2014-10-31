package user

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/md5"
	"fmt"
	"log"
	"time"
	"gokerap/db"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	IsAdmin  bool   `json:"bool"`
	Token    string `json:"token,omitempty"`
}

func New(email, password string) (*User, error) {
	u := &User{
		Email: email,
	}

	err := u.SetPassword(password)
	if err != nil {
		return nil, err
	}

	err = db.Conn.QueryRow(createUser, u.Email, u.Password).Scan(&u.Id)
	if err != nil {
		return nil, err
	}

	err = u.generateToken()
	if err != nil {
		return nil, err
	}

	return u, err
}

func Find(id int) (*User, error) {
	u := User{}
	err := db.Conn.QueryRow(getUser, id).Scan(
		&u.Id, &u.Email, &u.Password, &u.IsAdmin)

	return &u, err
}

func (u *User) SetPassword(password string) error {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = hpass

	return nil
}

func (u *User) generateToken() error {
	token := fmt.Sprintf("%s%s%v", u.Email, time.Now().UnixNano(), u.Password)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(token)))

	_, err := db.Conn.Exec(createUserToken, u.Id, hash)
	if err != nil {
		return err
	}

	u.Token = hash
	return nil
}

func AuthenticatePassword(email, password string) (*User, error) {
	u := User{}

	err := db.Conn.QueryRow(getUserByEmail, email).Scan(&u.Id, &u.Email, &u.Password, &u.Token)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		return nil, err
	}

	return &u, err
}

func AuthenticateToken(token string) (*User, error) {
	u := User{}
	err := db.Conn.QueryRow(getUserByToken, token).Scan(
		&u.Id, &u.Email)

	return &u, err
}

func init() {
	_, err := db.Conn.Exec(createUserTable)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Conn.Exec(createUserTokenTable)
	if err != nil {
		log.Println(err)
	}
}
