package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type storeTime struct {
	hours   []int
	minutes []int
	seconds []int
	message []string
}

func (s *storeTime) storeActions(hrs, min, sec int, y, cmd string) {
	store_time := *s
	store_time.message = append(store_time.message, y)
	store_time.hours = append(store_time.hours, hrs)
	store_time.minutes = append(store_time.minutes, min)
	store_time.seconds = append(store_time.seconds, sec)
	switch cmd {
	case "y":
		instanceover := 0
		for instanceover != 1 {
			Scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Type message you would like to see:")
			fmt.Println("")
			Scanner.Scan()
			result := Scanner.Text()
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
			fmt.Println("Do you want to time something else at the same time? (y/n)")
			Scanner.Scan()
			yesno := Scanner.Text()
			if yesno == "n" {
				instanceover = 1
			}
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
			*s = store_time
			store_time.storeActions(inthours, intminutes, intseconds, result, yesno)
		}

	case "n":
		for i := range store_time.message {
			go TimedAction(store_time.hours[i], store_time.minutes[i], store_time.seconds[i], store_time.message[i])
		}
		store_time = storeTime{}
		defer StartMessage()
	}
}

func StartMessage() {
	fmt.Println("==TIMER APP==")
	fmt.Println("(type 'q' to quit at any time)")
	fmt.Println("Anything already timed will be shown below...")
	fmt.Println("Type message you would like to see:")
	fmt.Println("")
}

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
	storedtime := storeTime{}
	gameover := 0
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("==TIMER APP==")
	fmt.Println("(type 'q' to quit at any time)")
	fmt.Println("Anything already timed will be shown below...")
	fmt.Println("Type message you would like to see:")
	fmt.Println("")
	for gameover != 1 {
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
		fmt.Println("Do you want to time something else at the same time? (y/n)")
		Scanner.Scan()
		yesno := Scanner.Text()
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
		storedtime.storeActions(inthours, intminutes, intseconds, result, yesno)

	}
}
