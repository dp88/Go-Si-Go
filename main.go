package main

/*
#pragma weak Finale
extern void Finale();
*/
import "C"

import "fmt"

//export HelloGo
func HelloGo() {
	fmt.Println("Hello from Go!")
	C.Finale()
}

func main() {
	
}
