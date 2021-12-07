package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]
	err := CheckArguments(args)
	if err != nil {
		log.Fatal(err)
	}

	day := args[0]
	session := args[1]

	serverResponse := GetProblemInput(day, session)
	err = CheckStatusCode(serverResponse.StatusCode)

	if err != nil {
		drawError()
		log.Fatal(err)
	}

	CreateFolder(day)

	CreateFiles(day, session)
	drawTree()
	fmt.Println("Day " + day + " created succesfully good luck and merry xmas!!")

}

func CheckArguments(args []string) error {

	if len(args) != 2 {
		return errors.New("Incorrect number of arguments")
	}
	return nil
}
