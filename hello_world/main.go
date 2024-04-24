package main

import (
	"flag"
	"fmt"
)

var krGreeting string

func main() {
	lang := flag.String("lang", "en", "greeting language")
	flag.Parse()

	krGreeting = "안녕, 친구"
	var enGreeting = "Hello"
	chinaGreeting := "你好朋友"

	if *lang == "en" {
		fmt.Println(enGreeting)
	} else if *lang == "kr" {
		fmt.Println(krGreeting)
	} else if *lang == "china" {
		fmt.Println(chinaGreeting)
	} else {
		fmt.Println("Language is not support.")
	}
}
