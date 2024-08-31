package main

import "flag"

var (
	name string
)

func main() {
	flag.StringVar(&name, "name", "world", "name to say hello to (default: world)")
	flag.Parse()

	println("Hello, " + name + "!")
}
