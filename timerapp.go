package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func TimedAction(hrs, min, sec int, y string) {
	if y == "" {
		return
	}
	fmt.Printf("%s will be alerted to you below in %d hours, %d minutes and %d seconds\n", y, hrs, min, sec)
	hrt := time.Duration(hrs * 3600)
	mint := time.Duration(min * 60)
	sect := time.Duration(sec * 1)
	time.Sleep(hrt * time.Second)
	time.Sleep(mint * time.Second)
	time.Sleep(sect * time.Second)
	fmt.Println(y)
}

func main() {
	gameover := 0
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("==TIMER APP==")
	fmt.Println("(type 'q' to quit at any time)")
	for gameover != 1 {
		fmt.Println("Anything already timed will be shown below...")
		fmt.Println("Type message you would like to see:")
		fmt.Println("")
		Scanner.Scan()
		result := Scanner.Text()
		if result == "q" {
			gameover = 1
			break
		}
		fmt.Println("2a) Now type in hours:")
		Scanner.Scan()
		hours := Scanner.Text()
		fmt.Println("2b) Now type in minutes:")
		Scanner.Scan()
		minutes := Scanner.Text()
		fmt.Println("2b) Now type in seconds:")
		Scanner.Scan()
		secondz := Scanner.Text()
		inthours, err := strconv.Atoi(hours)
		if err != nil {
			fmt.Println("Please type in numbers only...")
		}
		intminutes, err := strconv.Atoi(minutes)
		if err != nil {
			fmt.Println("Please type in numbers only...")
		}
		intseconds, err := strconv.Atoi(secondz)
		if err != nil {
			fmt.Println("Please type in numbers only...")
		}
		go TimedAction(inthours, intminutes, intseconds, result)

	}
}
