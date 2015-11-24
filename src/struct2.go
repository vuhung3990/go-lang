package main
import "fmt"

type People struct {
	name string
	age int
}

// get name of object people
func (p People) getName() string {
	return p.name
}

// set name of object
func (p *People) setName(name string)  {
	p.name = name
}

// example: addplus(4) = 1+2+3+4 = 10
func addplus(n uint) uint {
	if(n == 0){
		return 0
	}
	return n + addplus(n-1)
}

func main() {
	hung := People{name:"hung", age:25}
	fmt.Println(hung.getName())

	hung.setName("Vu Van Hung")
	fmt.Println(hung.getName())

	fmt.Println(addplus(9))
}
