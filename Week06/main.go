package main

import (
	"fmt"
	"time"

	counter "github.com/elliottlin2020/go-000/Week05/pkg"
)

func main() {
	counter := counter.New(time.Second, 10)
	fmt.Println(counter.Add(3))
	fmt.Println(counter.Add(3))
	time.Sleep(time.Second)
	fmt.Println(counter.Add(400))
	time.Sleep(time.Second)
	fmt.Println(counter.Add(100))
	fmt.Println(counter.Add(10))
}
