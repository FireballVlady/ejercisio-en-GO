package main

import (
	"fmt"
	"time"
)

var entradav int = 0
var salidav int = 0

func salida() {
	for i := 0; i > 1000; i++ {
		salidav--
	}
	fmt.Println(salidav)

}
func entrada() {
	for i := 0; i < 1000; i++ {
		entradav++

	}
	fmt.Println(entradav)
}
func entrada2() {
	for i := 0; i < 1000; i++ {

		entradav++
	}
	fmt.Println(entradav)
}
func imprime(message string) {
	fmt.Println(message)
}

func main() {

	go salida()
	go entrada()
	go imprime("hola desde goroutine")
	time.Sleep(10 * time.Second)
}
