package main

import (
	"fmt"
	"reflect"
)

type Interface interface {
	hello(name string) string
}

type Base struct {
}

func (*Base) hello(name string) string{
	return "Hello " + name
}

type Sub struct {
	*Base
}

//overwrite the function of Base
func (*Sub) hello(name string) string {
	return "Hi " + name
}

func main() {
	mycar := &Sub{}
	//call Sub.hello
	fmt.Println(mycar.hello("dma"))
	//call Base.hello
	fmt.Println(mycar.Base.hello("dma"))
	
	mymap := map[string]Interface {
		"first": mycar,
	}
	fmt.Println(reflect.TypeOf(mymap["first"]))
}
