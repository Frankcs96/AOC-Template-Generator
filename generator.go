package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateFolder(day string) {
	err := os.Mkdir("day"+day, 0755)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateFiles(day string, session string) {

	input := GetProblemInput(day, session)

	filename_main := "day" + day + ".go"
	filename_test := "day" + day + "_test.go"
	filename_input := "input.txt"
	filename_example_input := "example_input.txt"
	dir := "day" + day
	file, err := os.Create(filepath.Join(dir, filepath.Base(filename_main)))
	file.WriteString(getMainFileData(day))
	file, err = os.Create(filepath.Join(dir, filepath.Base(filename_test)))
	file.WriteString(getTestData(day))
	file, err = os.Create(filepath.Join(dir, filepath.Base(filename_input)))
	file.WriteString(input.Payload)
	file, err = os.Create(filepath.Join(dir, filepath.Base(filename_example_input)))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}

func drawTree() {

	tree := `  
             /\
            <  >
             \/
             /\
            /  \
           /++++\
          /  ()  \
          /      \
         /~*~*~*~*\
        /  ()  ()  \
        /          \
       /*&*&*&*&*&*&\
      /  ()  ()  ()  \
      /              \
     /++++++++++++++++\
    /  ()  ()  ()  ()  \
    /                  \
   /~*~*~*~*~*~*~*~*~*~*\
  /  ()  ()  ()  ()  ()  \
  /*&*&*&*&*&*&*&*&*&*&*&\
 /                        \
/,.,.,.,.,.,.,.,.,.,.,.,.,.\
           |   |
        __/_____\__
        \_________/
`

	fmt.Println("\033[32m" + tree)
}

func drawError() {

	error := `

  .d88b. 888d888888d888 .d88b. 888d888 
d8P  Y8b888P"  888P"  d88""88b888P"   
88888888888    888    888  888888     
Y8b.    888    888    Y88..88P888     
 "Y8888 888    888     "Y88P" 888

  `

	fmt.Println("\u001b[31m" + error)
}
