package main

import (
	"fmt"
	"github.com/damienbutt/figlet"
	"time"
)

func main() {
	// NOTE: Showing up that the cli is working
	welcome_message := welcomeMessage()
	cli_message := toShowInitMessage()

	fmt.Printf("%v -- %v\n", cli_message, welcome_message)
}
func toShowInitMessage() string {
	text,
		err := figlet.Text("go cli")
	if err != nil {
		panic(err)
	}
	return text
}
func welcomeMessage() string {
	real_time := time.Now()
	hours, _, _ := real_time.Clock()

	if hours >= 06 && hours < 12 {
		return "Good morning."
	} else if hours >= 12 && hours < 18 {
		return "Good afternoon."
	} else if hours >= 18 && hours < 23 {
		return "Good nigth."
	} else {
		return "Please go to sleep bro you need!"
	}

}
