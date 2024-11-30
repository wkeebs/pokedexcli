package main

import (
	"bufio"
	"fmt"
	"os"
)

func Reader() string {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Enter your command: ")

	key, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }

    return key
}