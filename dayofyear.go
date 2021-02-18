package main

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	dateArray := strings.Split(date, "-")
	year, _ := strconv.Atoi(dateArray[0])
	month, _ := strconv.Atoi(dateArray[1])
	day, _ := strconv.Atoi(dateArray[2])

	days := getAccumulatedMonthDays(year, month) + day

	return days

}

func isLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}

	return false
}

func getAccumulatedMonthDays(year, month int) int {
	days := 0
	for i := 1; i < month; i++ {
		days += getMonthDays(year, i)
	}

	return days
}

func getMonthDays(year, month int) int {
	everyMonthDays := map[int]int{
		1: 31, 3: 31, 5: 31, 7: 31, 8: 31, 10: 31, 12: 31,
		4: 30, 6: 30, 9: 30, 11: 30,
	}

	if isLeapYear(year) {
		everyMonthDays[2] = 29
	} else {
		everyMonthDays[2] = 28
	}

	return everyMonthDays[month]
}
