package tools

import (
	"fmt"
)

type Login struct {
	UserID   string
	Password string
}

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

func runLogin() {
	tmpInfo := Login{
		UserID:   "david",
		Password: "1234",
	}

	if login.UserID == tmpInfo.UserID {
		if login.Password == tmpInfo.Password {
			fmt.Println("Login Success!")
		} else {
			fmt.Println("Check your password")
		}
	} else {
		fmt.Println("Check your ID")
	}

}

func checkInfo() {

}
