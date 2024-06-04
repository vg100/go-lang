package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func main() {
	students := []Student{
		{
			Name:  "Vijay",
			Class: "5th",
		},
	}

	names := []interface{}{1, 2, 3, 4, 5}
	var temp int
	for i := 0; i < len(names); i++ {
		temp = temp + names[i]
		fmt.Println(names[i])
	}
	// arr := []interface{}{"name", 5, "uiu", 7, "jjj"}
	fmt.Println(students[0].Class)
	fmt.Println(names)
}
