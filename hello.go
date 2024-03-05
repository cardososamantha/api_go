package main

import "fmt"

const sufix = "Hello, "

func Hello(name string) string {
	return sufix + name;
}

func main() {
	fmt.Println(Hello("Isa"))
}
