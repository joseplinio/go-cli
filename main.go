package main

import (
	"fmt"
	"github.com/damienbutt/figlet"
)

func main() {
	text, err := figlet.Text("go cli")
	if err != nil {
		panic(err)
	}

	fmt.Println(text)
}
