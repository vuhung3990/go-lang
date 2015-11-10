package main
import "fmt"

func main() {
	// function
	fmt.Println(plus(3, 4, 5))
	
	// multiple return
	fmt.Println("=============================")
	a, b := multiple()
	// If you only want a subset of the returned values, use the blank identifier _
	_, c := multiple()
	fmt.Println(a, b)
	fmt.Println(c)

	// variadic functions (should use this method instead of limit param functions)
	fmt.Println("=============================")
	fmt.Println(sum(2,3,4,7))
	fmt.Println(sum(2,2,1))
	// variadic array
	numbers := []int{1,4,5,6}
	fmt.Println(sum(numbers...))

	
}

// plus 3 number
func plus(a, b, c int) int {
	return a + b + c
}

// The (int, int) in this function signature shows that the function returns 2 ints.
func multiple() (int, int) {
	return 3,5
}

// param is unlimited
func sum(nums ...int) int {
	total := 0
	for _, num:= range nums{
		total+= num
	}
	return total
}