package main

import (
	"fmt"
	"time"
)

func main() {
	op := func(i int) {
		fmt.Println("op", i, time.Now().Format("5"))
		time.Sleep(1 * time.Second)
		fmt.Println("op", i, time.Now().Format("5"))
		time.Sleep(1 * time.Millisecond)
		fmt.Println("temporizador ", i, time.Now().Format("5"))

	}

	comienzo := time.Now()

	for i := 12; i > 5; i-- {
		op(i)
	}

	fmt.Println("total", time.Since(comienzo))

}
