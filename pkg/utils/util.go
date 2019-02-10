package utils

import (
	"fmt"
	"os/user"
)

func GetUserHome() string {
	usr, err := user.Current()
	if err != nil {
		panic(fmt.Errorf("failed to get user home dir. Error %+v", err))
	}
	return usr.HomeDir
}
