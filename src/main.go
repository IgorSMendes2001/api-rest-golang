package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Igor!")
	fmt.Println(os.Getenv("TEST"))
}
