package commands

import (
	"fmt"
	"time"

	"github.com/damienbutt/figlet"
)

func Status() {
	cliName := cliName()
	timeGreeting := greeting()

	fmt.Printf("%v -- %v\n", cliName, timeGreeting)
	fmt.Printf("Version: %0.1f\n", 0.1)
}

func cliName() string {
	text,
		err := figlet.Text("Eclipse Cli")
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
		return "Please go to sleep, You need!"
	}
}
func getLastCommit() {

}
