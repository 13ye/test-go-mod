package xyz

import (
	"fmt"
	"xyz/test-go-mod/cmd"
)

func GetMe() {
	fmt.Println("GetMe")
	cmd.GetCmd("2333")
}
