package main

import (
	"log"
	"os"
)

func main() {
	host_name := get_hostname()
	user_home_dir := get_user_home_dir()

	userData := [2]string{host_name, user_home_dir}
	file := open_file("user-activity.txt")
	for x := range userData {
		write_file(file, userData[x])
	}
}

