package main

import (
	"clonkk/mthreads/fn1"
	"clonkk/mthreads/fn2"
	"clonkk/mthreads/tools"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Let's try some go routines")
	payload := make(chan string)
	go tools.ToolHelper()
	go fn1.Fn1(11, payload)
	go fn2.Fn2(22, payload)
	payload <- "MSG1"
	fmt.Println("WAITING...")
	time.Sleep(5 * time.Second)
	fmt.Println("WAITING DONE")
	payload <- "MSG2"

	fmt.Println("Bye-bye !")
}
