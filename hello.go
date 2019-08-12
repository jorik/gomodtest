package hello

import "fmt"

func GetHello() string {
	return "Hello world, v1.1.1!"
}

func SayHello() {
	fmt.Println(GetHello())
}
