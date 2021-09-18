package main

import (
	"evpeople/toyLang/bot"
	"fmt"
)

func main() {
	a := bot.New("sd", make(map[string]string), 21)
	a.AddQa("ds", "wd")
	fmt.Println(a)

}
