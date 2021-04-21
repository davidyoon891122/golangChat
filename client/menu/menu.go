package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../../tools"
)

func DisplayMenu() []byte {

	for {
		var menuSelect int

		fmt.Println("*****************")
		fmt.Println("1. Login ")
		fmt.Println("2. Exit ")
		fmt.Println("*****************")
		fmt.Println("Select menu : ")
		fmt.Scanf("%d", &menuSelect)

		if menuSelect == 1 {
			return tools.Pack(getUserInfo(), 1, 1, 0)
		} else if menuSelect == 2 {
			os.Exit(0)
		}
	}
}

func getUserInfo() *tools.Login {
	loginInfo := tools.InitLogin()
	fmt.Println("User ID : ")
	in := bufio.NewReader(os.Stdin)
	idData, err := in.ReadString('\n')
	idData = strings.TrimRight(idData, "\n")
	if err != nil {
		panic(err)
	}

	fmt.Println("User Password : ")
	pwdData, err := in.ReadString('\n')
	pwdData = strings.TrimRight(pwdData, "\n")
	fmt.Print(pwdData)
	if err != nil {
		panic(err)
	}

	loginInfo.LoginPacker(idData, pwdData)

	return loginInfo
}
