package main

import (
	"fmt"
	"time"
)

//model
var singers = []singer{}

type singer struct {
	name string
}

type song struct {
	name   string
	singer singer
	date   time.Time
}

func Scans(message string, a interface{}) {
	fmt.Println(message)
	fmt.Scan(a)
}

func main() {
	var first_name, last_name string
	i, err := fmt.Sscan("ou qingjia", &first_name, &last_name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Println(first_name + last_name)
}
