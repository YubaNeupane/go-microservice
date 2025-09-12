package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello from driver service")

	for {
		time.Sleep(time.Second * 1)
	}
}
