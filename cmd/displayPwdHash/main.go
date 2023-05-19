package main

import (
	"fmt"
	"os"
	"github.com/Serhii1Epam/lets-go-chat/pkg/hasher"
)

func main() {
	fmt.Printf("displayPwdHash =>%v, Hash => %s", os.Args[1:], HashPassword(os.Args[1:]))
}
