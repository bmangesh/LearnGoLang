package main

import (
	"fmt"
)

// Define a struct for a Virtual Machine to store it's attributes
type VM struct {
	MacID    string
	IP       string
	hostname string
	OsFamily string
	OsName   string
}

func main() {

	//Declaring and Initializing a struct using struct literal
	vm1 := VM{
		"28:cf:e9:00:91:8c",
		"172.17.254.80",
		"openstack",
		"Linux",
		"Centos-7",
	}

	// Print the value of vm1 variable.
	fmt.Println(vm1)
}
