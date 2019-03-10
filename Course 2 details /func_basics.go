package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "roger okello"
	moduleName := "Python Beyond the basics"

	a, b := converter(name, moduleName)

	fmt.Println(a, b)
}

/*
The  function parameter list could have been writen as (name, moduleName string)
The return value parameter list cannot be the same as the input parameter list
ie. if the function parameter list looks like this (name string, moduleName string),
the return value parameter list cannot look in any way like this
(name string, moduleName string)
*/
func converter(name string, moduleName string) (returnName, returnModuleName string) {
	name = strings.Title(name)
	moduleName = strings.ToUpper(moduleName)

	return name, moduleName
}
