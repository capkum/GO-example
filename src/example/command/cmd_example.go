package main

import (
	"fmt"
	"os/exec"
)

func execute() {
	cmd := "netstat -an | grep 'LISTEN'"

	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(out)
	}
	fmt.Println(string(out))
}

func main() {
	fmt.Println("simple Shell")
	fmt.Println("-----------------------------")
	execute()
}
