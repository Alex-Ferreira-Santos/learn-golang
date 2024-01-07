package main

import (
	"fmt"
)

type Student struct {
	name   string
	grades []int
	age    int
}

func (student Student) calculateAverageGrade() float32 {
	var sum int
	for _, value := range student.grades {
		sum += value
	}
	return float32(sum) / float32((len(student.grades)))
}

func (student Student) getMaxGrade() int {
	var currentMaxValue int
	for _, value := range student.grades {
		if value > currentMaxValue { currentMaxValue = value}
	}

	return currentMaxValue
}

func main() {
	student := Student{"Alex", []int{77, 83, 94}, 20}
	fmt.Printf("average grade: %f\n", student.calculateAverageGrade())
	fmt.Printf("max grade: %d", student.getMaxGrade())
}
