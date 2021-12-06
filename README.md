# AOC Template Generator 🎄

# WORK IN PROGRESS ⚙️

I'm tired of creating a folder for each day with the same file structure so why not having a tool to automate that proccess 🛠 !

In case you don't know about Advent of code, it's a website where you will have to face a new challenge from the first of December until the 25th
https://adventofcode.com/


# Folder and file structure

Example:
- Day06
  - day06.go
  - day06_test.go
  - input.txt
  - example_input.txt

all the files have a basic structure to have a quick start of the problems.

the example_input.txt will be empty because this is not your input. This file will be here for adding the example in the problem description.
I find this really good since you have the solution for that test case so you can test your code in a easy way.

input.txt will be autofilled with your current input of the problem. This will be done using your session cookie and a http request to the advent of code website.

# How to use it

You have to run the program with 2 parameters:
- Day of the problem
- Path
- Session cookie

following command will generate a folder day12 with all the files in the current directory

```bash
go run 12 ./ 23420324abc
```
I prefer just to add the executable to the path so I just do this 

```bash
aoc 12 ./ 234234jsd24
```

To generate the executable just run

```bash
go build
```

