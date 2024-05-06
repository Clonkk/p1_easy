package fn1

import (
	"fmt"
)

func Fn1(thrid int, payload chan string) {
	msg := <-payload
	fmt.Println(thrid, " -> ", msg)
}
