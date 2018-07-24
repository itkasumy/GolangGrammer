package main

import (
	"os"
	"fmt"
)

func main()  {
	list := os.Args
	n := len(list)
	fmt.Println("n = ", n)

	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}

	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}
