package main
import "fmt"
func main() {
	next := closuresDemo()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

// must have func()
func closuresDemo() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}