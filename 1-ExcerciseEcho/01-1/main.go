package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	user := "Aman"
	pass := "jaiho"
	id := fmt.Sprintf("%v:%v", user, pass)
	bs, _ := bcrypt.GenerateFromPassword([]byte(id), bcrypt.MinCost)
	fmt.Println(string(bs))

	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s = " - "
		s += sep + os.Args[i]
		sep = "hiiiiii"

	}
	fmt.Println("Lenght of Arguments", len(os.Args))
	fmt.Println("Capacity of Arguments", cap(os.Args))
	fmt.Println(os.Args[:])
	fmt.Println(s)
}
