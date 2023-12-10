package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	path := getBasePath() + "/../../../resources/day1/input"
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var acc int

	rows := strings.Split(string(dat), "\n")
	for _, row := range rows {
		if len(row) > 0 {
			acc += getNumber(row)
		}
	}

	fmt.Println(acc)
}

func getBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}

var stringDigits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func validDigit(charsFound string) int {
	if len(charsFound) >= 3 {
		for i, digit := range stringDigits {
			if strings.Contains(charsFound, digit) {
				return i + 1
			}
		}
	}
	return -1
}

func getFistUnit(row string) int {
	var charsFound string
	for i := 0; i < len(row); i++ {
		number, err := strconv.Atoi(string(row[i]))
		if err == nil {
			return number
		} else {
			charsFound += string(row[i])
			if digit := validDigit(charsFound); digit != -1 {
				return digit
			}
		}
	}
	return -1
}

func getLastUnit(row string) int {
	var charsFound string
	for i := len(row) - 1; i >= 0; i-- {
		number, err := strconv.Atoi(string(row[i]))
		if err == nil {
			return number
		} else {
			charsFound = string(row[i]) + charsFound
			if digit := validDigit(charsFound); digit != -1 {
				return digit
			}
		}
	}
	return -1
}

func getNumber(row string) int {
	firstUnit := getFistUnit(row)
	lastUnit := getLastUnit(row)

	if firstUnit == -1 {
		return 0
	} else if lastUnit == -1 {
		return firstUnit
	}

	combined := strconv.Itoa(firstUnit) + strconv.Itoa(lastUnit)

	number, err := strconv.Atoi(combined)
	if err != nil {
		panic(err)
	}

	return number
}
