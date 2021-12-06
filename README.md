# AOC Template Generator 🎄

I'm tired of creating a folder for each day with the same file structure so why not having a tool to automate that proccess 🛠 !

In case you don't know about Advent of code, it's a website where you will have to face a new challenge from the first of December until the 25th
https://adventofcode.com/


# How to use it

You have to run the program with 2 parameters:
- Day of the problem
- Path

following command will generate a folder day12 with all the files in the current directory

```bash
go run 12 ./
```
I prefer just to add the executable to the path so I just do this 

```bash
aoc 12 ./
```

To generate the executable just run

```bash
go build
```
