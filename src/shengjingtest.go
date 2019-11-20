package main

import "fmt"

//import "fmt"

type person struct {
	name string
}
func main01() {
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}
