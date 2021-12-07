package main_test

import (
	"testing"

	. "github.com/Frankcs96/AOC-Template-Generator"
)

func TestArgumentsNumber(t *testing.T) {

	ValidArguments := []string{"03", "239424AB"}
	invalidArguments := []string{"03", "239424AB", "sad"}

	err := CheckArguments(ValidArguments)
	err2 := CheckArguments(invalidArguments)

	if err != nil {

		t.Errorf("Expecting 2 arguments got %v arguments", len(ValidArguments))
	}

	if err2 == nil {

		t.Errorf("got %v arguments so it should not throw an error", len(invalidArguments))
	}
}
