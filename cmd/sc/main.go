package main

import (
	"fmt"
	sshcommand "github.com/squarescale/sshcommand"
)

func main() {
	sc, err := sshcommand.New([]string{"ssh", "user@host"})
	if err != nil {
		fmt.Printf("%v\n", sc)
	}
	fmt.Printf("%+v\n", sc)
	hostname := sc.Hostname()
	fmt.Printf("%+v\n", hostname)
}
