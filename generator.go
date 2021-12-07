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

func CreateFiles(day string) {

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
