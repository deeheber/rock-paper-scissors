package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
}
