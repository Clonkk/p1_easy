package fn2

import "fmt"

func Fn2(thrid int, payload chan string) {
	msg := <-payload
	fmt.Println(thrid, " -> ", msg)
}
