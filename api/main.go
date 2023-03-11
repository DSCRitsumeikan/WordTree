package main

import (
	"WordTree/infrastructure"
	"fmt"
)

func main() {
	_ = infrastructure.InitDB()
	infrastructure.NewRouter()
	fmt.Println("hello")
}
