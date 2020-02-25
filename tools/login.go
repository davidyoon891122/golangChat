package tools

import (
	"fmt"
)

type Login struct {
	UserID   string
	Password string
}

var LoginCode map[int]string = map[int]string{
	00: "Login Success",
	01: "Check Your ID",
	02: "Check Your Password",
}

type Accounts []Login

func InitLogin() *Login {
	login := &Login{}

	return login
}

func (l *Login) LoginPacker(userID string, password string) {
	l.putUserID(userID)
	l.putPassword(password)
}

func (l *Login) putUserID(userID string) {
	l.UserID = userID
}

func (l *Login) putPassword(password string) {
	l.Password = password
}

func readLogin() {
	login = InitLogin()
	login.UserID = readString()
	fmt.Println(login.UserID)
	login.Password = readString()
	fmt.Println(login.Password)
}

func runLogin() (bool, int) {
	accounts := Accounts{
		{
			UserID:   "david",
			Password: "1234",
		},
		{
			UserID:   "qqpo12",
			Password: "1234",
		},
	}

	//for loop for checking registed User.
	fmt.Println(accounts)
	var res bool
	var code int

	for _, account := range accounts {
		if login.UserID == account.UserID {
			if login.Password == account.Password {
				fmt.Println("Login Success!")
				res = true
				code = 00
				break
			} else {
				fmt.Println("Check your password")
				res = false
				code = 02
				break
			}
		} else {
			fmt.Println("Check your ID")
			res = false
			code = 01
		}
	}
	return res, code
}

func (l *Login) GetUserID() string {
	return l.UserID
}

func (l *Login) GetPassword() string {
	return l.Password
}
