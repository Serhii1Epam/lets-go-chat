package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("displayPwdHash =>%v, Hash => %s", os.Args[1:], HashPassword(os.Args[1:]))
}
