package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need arg, which is the password")
		os.Exit(3)
	}
	pwd := os.Args[1]
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(hash))
}
