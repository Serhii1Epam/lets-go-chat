/* Cmd util for printing hash
* Use: displayPwdHash <password>
 */
package main

import (
	"fmt"
	"os"
	"github.com/Serhii1Epam/lets-go-chat/pkg/hasher"
)

func main() {
	var arg string

	if numArgs := len(os.Args); numArgs > 1 {
		arg = os.Args[1]
	} else {
		fmt.Println("Hashing errors: epmpty password entered")
		return
	}

	hash, err := hasher.HashPassword(arg)
	if err != nil {
		fmt.Println("Hashing errors: arg => ", arg)
	}

	fmt.Printf("displayPwdHash: Password => %s, Hash => %s", arg, hash)
}
