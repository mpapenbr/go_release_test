package main

import (
	"fmt"

	"github.com/mpapenbr/go_release_test/cmd"
)

var (
	version = "localdev"
)

func main() {

	fmt.Printf("This is Sample app LocalVersion: %s Cmd.Version: %s MyVersion: %s\n", version, cmd.GetVersion(), cmd.MyVersion)
	fmt.Printf("We can now print another line!")
	fmt.Printf("Try again!")
	fmt.Printf("And still another try!")
	fmt.Printf("Here we go again")
	fmt.Printf("Yeah!!")

}
