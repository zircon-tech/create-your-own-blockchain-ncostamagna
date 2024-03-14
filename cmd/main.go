package main

import (
	"fmt"
	"github.com/zircon-tech/create-your-own-blockchain-ncostamagna/pkg/crypto"
	"time"
)

func main() {
	b1 := crypto.NewBlock([]byte("Test123"), time.Now())
	b2 := crypto.NewBlock([]byte("Test456"), time.Now())
	b3 := crypto.NewBlock([]byte("Test231"), time.Now())
	bc := crypto.NewBC()

	if valid := bc.Add(b1).Add(b2).Add(b3).IsValid(); valid {
		fmt.Println("Is OK")
	} else {
		fmt.Println("Bad")
	}

}
