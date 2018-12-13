package main

import (
	"fmt"
)

// import "strconv"

type X int

func (x *X) inc() {
	*x++
}

type userErr interface {
	error
	Message() string
}

type haha string

func (u haha) Error() string{
	return string(u)
}

func (u haha) Message() string{
	return string(u)
}


func main() {
	fmt.Println("hello world")

	//var x X
	//x.inc()
	//println(x)

	// a, _ := strconv.ParseInt("1100100",2,32)
	// b, _ := strconv.ParseInt("0144",8,32)
	// c, _ := strconv.ParseInt("64",16,32)

	// println(a,b,c)

}
