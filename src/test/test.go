package main

import (
	"fmt"
	"unicode/utf8"
)

func testMap() {
	map1 := make(map[string]int)
	map1["test"] = 1
	map1["tp"] = 2
	for key, value := range map1 {
		fmt.Println(key, ":", value)
	}
	values, ok := map1["ggg"]
	if ok == true {
		fmt.Println(values)
	} else {
		fmt.Println("name is not in map")
	}
	map2 := map[string]int{
		"test": 1,
		"tp":   2,
	}
	m := 0
	for key, value := range map1 {
		if map2[key] != value {
			m++
		}
	}
	if len(map1) != len(map2) {
		m++
	}
	if m == 0 {
		fmt.Println("map1 and map2 are same")
	} else {
		fmt.Println("map1 and map2 are not same")
	}
}

func testSlice() {
	a := [...]int{1, 2, 3, 4, 5}
	aslice := a[1:2]
	fmt.Println(aslice)
	for i := range aslice {
		aslice[i]++
	}
	fmt.Println("length %d cap %d", len(aslice), cap(aslice))
	aslice = aslice[:cap(aslice)]
	fmt.Println("length %d cap %d", len(aslice), cap(aslice))
	fmt.Println(aslice)
	var b []int
	fmt.Println(cap(b), len(b))
	b = append(b, 1)
	fmt.Println(cap(b), len(b))
}

func testString() {
	name := "Hello World"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c", name[i])
	}
	name2 := "Señor"
	for index, rune := range name2 {
		fmt.Printf("char %c is %d bytes\n", rune, index)
	}
	fmt.Println(len(name2))
	fmt.Println(utf8.RuneCountInString(name2))
	r_name2 := []rune(name2)
	r_name2[0] = 'a'
	fmt.Println(string(r_name2))
}

func zhizhen() {
	a := 112
	b := &a
	fmt.Println(*b)
}

func structTest() {
	type Place struct {
		city, state string
	}
	type Employee struct {
		name     string
		position string
		age      int
		Place
	}
	emp := Employee{"lisan", "bp", 13, Place{"guangzhou", "guangdong"}}
	fmt.Println(emp)
	fmt.Println(emp.age)
	var emp1 Employee
	emp1.age = 10
	fmt.Println(emp1)
	fmt.Println(emp.city)
}

func main() {
	testSlice()
}
