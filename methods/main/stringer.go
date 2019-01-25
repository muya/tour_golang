pack

age main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Stringer Bell", 43}
	z := Person{"Avon Barksdale", 234}
	fmt.Println(a, z)
}
