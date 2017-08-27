package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Age    int
	Height int
}

func main() {
	Students := []*Student{
		&Student{"Danny", 15, 165},
		&Student{"Jacky", 16, 180},
		&Student{"Alan", 17, 172},
		&Student{"Sandy", 18, 168},
	}
	result1 := Filter(Students, AgeGreatThanFunc(16))
	fmt.Println("result1")
	for _, s := range result1 {
		fmt.Println(s.Name)
	}
	result2 := Filter(Students, ComplexFunc(15, 170))
	fmt.Println("result2")
	for _, s := range result2 {
		fmt.Println(s.Name)
	}
}

func Filter(l []*Student, f func(*Student) bool) []*Student {
	result := []*Student{}
	for _, s := range l {
		if f(s) {
			result = append(result, s)
		}
	}
	return result
}

func AgeGreatThanFunc(age int) func(*Student) bool {
	return func(s *Student) bool {
		return s.Age > age
	}
}

func ComplexFunc(age int, height int) func(*Student) bool {
	return func(s *Student) bool {
		return (s.Age > age) && (s.Height > height)
	}
}
