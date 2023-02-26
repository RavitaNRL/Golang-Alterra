package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score int
}

type Students []Student

func (z Students) AverageScore() float64 {
	var sum int
	for _, student := range z {
		sum += student.Score
	}
	return float64(sum) / float64(len(z))
}

func (x Students) MinScore() (string, int) {
	var minScore int = math.MaxInt32
	var minName string

	for _, student := range x {
		if student.Score < minScore {
			minScore = student.Score
			minName = student.Name
		}
	}
	return minName, minScore
}

func (s Students) MaksScore() (string, int) {
	var maksScore int = math.MinInt32
	var maksName string

	for _, student := range s {
		if student.Score > maksScore {
			maksScore = student.Score
			maksName = student.Name
		}
	}
	return maksName, maksScore
}

func main() {
	var students Students

	for i := 0; i < 5; i++ {
		var name string
		var score int
		fmt.Printf("Input %d student name: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Input %d student score: ", i+1)
		fmt.Scan(&score)

		student := Student{
			Name:  name,
			Score: score,
		}
		students = append(students, student)
	}

	averageScore := students.AverageScore()
	minName, minScore := students.MinScore()
	maksName, maksScore := students.MaksScore()

	fmt.Printf("Average score: %.2f\n", averageScore)
	fmt.Printf("Min score of student: %s %d\n", minName, minScore)
	fmt.Printf("Max score of student: %s %d\n", maksName, maksScore)
}
