package main

import (
	"fmt"
)

type verse struct {
	subject string
	verb    string
	action  string
}

var lyrics = [...]verse{
	{"wheels", "go", "round and round"},
	{"wipers", "go", "Swish, swish, swish"},
	{"horn", "goes", "Beep, beep, beep"},
	{"people", "go", "Up and down"},
	{"driver", "says", "move on back"},
	{"babies", "say", "Wah, wah, wah"},
	{"dadies", "say", "Ssh, ssh, ssh"},
	{"wheels", "go", "round and round"},
}

func main() {
	for _, e := range lyrics {
		fmt.Printf("The %s on the bus %s %s,\n", e.subject, e.verb, e.action)
		fmt.Printf("%s; %s.\n", e.action, e.action)
		fmt.Printf("The %s on the bus %s %s,\n", e.subject, e.verb, e.action)
		fmt.Printf("all through the town\n")
		fmt.Printf("\n")
	}
}
