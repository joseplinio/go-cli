package main

import (
	"fmt"
	"github.com/damienbutt/figlet"
	"time"
)

func main() {
	// NOTE: Showing up that the cli is working
	// FIX: The variables's name are so bad, you need to improve this
	// if you want, start to reread the book Clean Code

	greeting_message := greeting()
	cli_message := firts_message()
	fmt.Printf("%v -- %v\n", cli_message, greeting_message)
}
func firts_message() string {
	text,
		err := figlet.Text("go cli")
	if err != nil {
		panic(err)
	}
	return text
}
func greeting() string {
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
