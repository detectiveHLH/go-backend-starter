package authentication

import "fmt"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	userName := a.Username
	passWord := a.Password
	fmt.Println(userName, passWord)
	return true, nil
}
