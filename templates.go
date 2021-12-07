package main

func getTestData(day string) string {

	data := `package day` + day + `_test

import (
	"testing"
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day` + day + `"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution("input-example.txt")
	expected := -1

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionPartTwo(t *testing.T) {

  t.SkipNow()
	solution := SolutionPartTwo("input-example.txt")
	expected := -1

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolution(t *testing.T) {

	solution := Solution("input.txt")
	expected := -1

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolutionPartTwo(t *testing.T) {

  t.SkipNow()
	solution := SolutionPartTwo("input.txt")
	expected := -1

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
`

	return data
}
func getMainFileData(day string) string {

	data := `package day` + day + `

import (
	"bufio"
	"log"
	"os"
)

func Solution(filename string) int {

  return 1
}

func SolutionPartTwo(filename string) int {

  return 1
}

func GetInput(filename string) []string {

	input := []string{}
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := scanner.Text()
		input = append(input, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}
`
	return data
}
