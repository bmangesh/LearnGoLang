package main

import (
	"fmt"
)

type VM struct {
	MacID    string
	IP       string
	hostname string
	OsFamily string
	OsName   string
}

func main() {
	fmt.Println(VM{MacID: "28:cf:e9:00:55:88", hostname: "Docker"})
}
