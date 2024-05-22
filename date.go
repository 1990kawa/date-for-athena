package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("start")

	flag.Parse()
	args := flag.Args()

	startStringDate := args[0]
	endStringDate := args[1]

	layout := "2006-01-02 15:04:05"

	startDate, _ := time.Parse(layout, startStringDate)
	endDate, _ := time.Parse(layout, endStringDate)

	for {
		if startDate.After(endDate) {
			break
		} else {
			fmt.Printf("dt = '%v-%v-%02d-%02d' or \n", startDate.Year(), getMonth(startDate.Month().String()), startDate.Day(), startDate.Hour())
			startDate = startDate.Add(time.Hour * 1)
		}
	}
}

func getMonth(month string) string {
	months := map[string]string{
		"January":   "01",
		"February":  "02",
		"March":     "03",
		"April":     "04",
		"May":       "05",
		"June":      "06",
		"July":      "07",
		"August":    "08",
		"September": "09",
		"October":   "10",
		"November":  "11",
		"December":  "12",
	}
	return months[month]
}
