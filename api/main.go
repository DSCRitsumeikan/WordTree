package main

import (
	"WordTree/controller"
	"WordTree/infrastructure"
	"fmt"
)

func main() {
	_ = infrastructure.InitDB()
	controller.NewRouter()
	fmt.Println("hello")
}
