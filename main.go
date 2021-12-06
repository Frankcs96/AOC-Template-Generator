package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) != 3 {

		fmt.Println("Please enter the correct parameters [daynumber] [Path] [Session cookie]")
		os.Exit(-1)
	}

	day := args[0]
	path := args[1]
  session := args[2]

	CreateFolder(day, path)






  drawTree()
  fmt.Println("Day " + day + " created succesfully good luck and merry xmas!!")


}

func CreateFolder(day string, path string) {
	err := os.Mkdir(path+"/"+day, 0755)
	if err != nil {
		log.Fatal(err)
	}

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
