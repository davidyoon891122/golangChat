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
	tmpInfo := Login{
		UserID:   "david",
		Password: "1234",
	}

	if login.UserID == tmpInfo.UserID {
		if login.Password == tmpInfo.Password {
			fmt.Println("Login Success!")
			return true, 00
		} else {
			fmt.Println("Check your password")
			return false, 02
		}
	} else {
		fmt.Println("Check your ID")
		return false, 01
	}

}

func (l *Login) GetUserID() string {
	return l.UserID
}

func (l *Login) GetPassword() string {
	return l.Password
}
