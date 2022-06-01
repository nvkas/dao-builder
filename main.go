package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		//恢复程序的控制权
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("hello word!")
}
