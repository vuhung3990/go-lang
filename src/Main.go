package main

import (
	"fmt"
)

func main() {
	// print hello world
	fmt.Println("hello world");

	// variables
	var name string = "fjgdsjhgsdjfs"
	fmt.Println(name)
	var isTrue bool = false
	fmt.Println(isTrue)

	// syntax is shorthand for declaring and initializing a variable
	age := 12
	fmt.Println(age);

	// constance
	const HERO string = "hung";

	// loop
	fmt.Println("===============================");
	for i := 0; i < 10; i++ {
		fmt.Printf("number: %d\n", i);
	}

	// conditional
	fmt.Println("===============================");
	if isTrue {
		fmt.Println("The True King");
	} else {
		fmt.Println("well it's not true");
	}

	// switch - case
	fmt.Println("===============================");
	switch age {
	case 10:
		fmt.Println("10");
	case 11:
		fmt.Println("11");
	default:
		fmt.Println("other");
	}

	// array
	fmt.Println("===============================");
	// must have length
	var names [9]int;
	for i := 0; i < 3; i++ {
		names[i] = i
	}
	fmt.Println(names);
	// short hand
	diemdanh := []bool{true, false, true, true}
	fmt.Println(diemdanh);
	//array length
	fmt.Println(len(names));

	// slices
	fmt.Println("===============================");
	s := make([]string, 3);
	// modify (3)
	s[0] = "ada";
	s[1] = "dsa";
	s[2] = "gdsu"
	// append will resize s
	s = append(s, "pougf")
	s = append(s, "eyui")
	fmt.Println(s);
	// print all element from 2->5
	fmt.Println("s[2:5] ", s[2:5]);
	// print all element from 2 to end
	fmt.Println("s[2:] ", s[2:]);
	// print all element from start to 2
	fmt.Println("s[:2] ", s[:2]);

	// map
	fmt.Println("===============================");
	m := make(map[string]int)
	m["me"] = 25;
	m["duyhung"] = 27;
	m["trien"] = 26;
	fmt.Println("map: ", m);
	// delete element in map
	delete(m, "me")
	fmt.Println("map after delete: ", m);

	// range
	fmt.Println("===============================");
	for i, v := range s {
		fmt.Println("range s: ", i, v);
	}

	for k, v:= range m{
		fmt.Println("range map: ",k, v);
	}
}
