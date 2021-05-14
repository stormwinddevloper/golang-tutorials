package main

import "fmt"

// func updateName(n string) {
// 	n = "wedge"
// }

func updateName(n *string) {
	*n = "wedge"
}

func main() {
	name := "tifa"

	// updateName(name)
	// fmt.Println(name)

	// & gets the memory address of the value (pointer)
	fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	updateName(&name) // using pointer as arg
	fmt.Println(name)

}

/*

|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|

*/

