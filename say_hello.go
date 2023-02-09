package golatmodule

import "fmt"

func SayHello() string {
	return "Hello World"
}

func Welcome(name string) string {
	return fmt.Sprintf("welcome, %s", name)
}
