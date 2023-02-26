package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score []int
}

func (s *Student) AddScore(score int) {
	s.Score = append(s.Score, score)
}

func (s *Student) GetAverageScore() float64 {
	var sum int
	for _, score := range s.Score {
		sum += score
	}
	avg := float64(sum) / float64(len(s.Score))
	return math.Round(avg)
}

func (s *Student) GetMinScore() (string, int) {
	var minScore int = math.MaxInt64
	var minName string
	for i, score := range s.Score {
		if score < minScore {
			minScore = score
			minName = s.Name + fmt.Sprintf(" (%d)", i+1)
		}
	}
	return minName, minScore
}

func (s *Student) GetMaxScore() (string, int) {
	var maxScore int
	var maxName string
	for i, score := range s.Score {
		if score > maxScore {
			maxScore = score
			maxName = s.Name + fmt.Sprintf(" (%d)", i+1)
		}
	}
	return maxName, maxScore
}

func main() {
	var students []Student

	for i := range students {
		if len(students[i].Score) < 5 {
			fmt.Printf("Input %d Student's Name ", i+1)
			fmt.Scanln(&students[i].Name)
			fmt.Printf("Input %d Student's Score ", i+1)
			var score int
			fmt.Scanln(&score)
			students[i].AddScore(score)
		}
	}

	var totalScore int
	for _, s := range students {
		totalScore += int(s.GetAverageScore())
		fmt.Printf("%s average score: %d\n", s.Name, int(s.GetAverageScore()))

		minName, minScore := s.GetMinScore()
		fmt.Printf("Min score of %s: %d\n", minName, minScore)

		maxName, maxScore := s.GetMaxScore()
		fmt.Printf("Max score of %s: %d\n", maxName, maxScore)
	}
	fmt.Printf("Overall average score: %d\n", totalScore/len(students))
}
