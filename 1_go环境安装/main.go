package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("hello")
	fmt.Println("a random int : ", rand.Intn(50))
}

/*
go run main.go

// result:
hello
a random int :  31
*/
