package main

import (
	"fmt"
	"time"
)

func main() {
	var setTime int64

	setTime = 1514029652

	then := time.Unix(setTime, 0)

	fmt.Println(then.Format("02 January, 2006"))

	fmt.Println(time.Now().Unix())
}
