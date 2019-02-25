package main

import (
	"fmt"

	"github.com/Alikibidan/GoLangBasic/002_packages/hello"
	"github.com/Alikibidan/GoLangBasic/002_packages/say"
)

func main() {
	fmt.Println(say.CallFromSay() + hello.CallFromHello())
}
