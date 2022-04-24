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

}
