package hello

import "fmt"

func GetHello(ignored int) string {
	return "Hello world, v2.0.0!"
}

func SayHello() {
	fmt.Println(GetHello(1))
}
