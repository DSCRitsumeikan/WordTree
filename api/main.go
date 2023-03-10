package main

import (
	"WordTree/infrastructure"
	"fmt"
)

func main() {
	_ = infrastructure.InitDB()
	fmt.Println("hello")
}
