package xyz

import (
	"fmt"
	"github.com/13ye/test-go-mod/cmd"
)

func GetMe() {
	fmt.Println("GetMe")
	cmd.GetCmd("2333")
}
