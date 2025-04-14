package main

import "fmt"

func commandJump() error {
	fmt.Println()
	fmt.Println("I'm a jumper!")
	fmt.Println("Jumping")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
