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

func open_file(name string) *os.File {
	// FIX: I don't know what this "0666" mean
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
	return file
}
