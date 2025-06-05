package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for: \"%s\"\n"
	errPwd   = "Invalid password for \"%s\"\n"
	accessOk = "Access granted to \"%s\"\n"

	user, user2         = "jack", "inanc"
	password, password2 = "1888", "1879"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user && u != user2 {
		fmt.Printf(errUser, u)
		return
	}

	if u == user && p == password {
		fmt.Printf(accessOk, u)
		return
	}

	if u == user2 && p == password2 {
		fmt.Printf(accessOk, u)
		return
	}

	fmt.Printf(errPwd, u)
}
