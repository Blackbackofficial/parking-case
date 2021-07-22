package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type TimeCar struct {
	timeA time.Time
	isArrival bool
}

func main() {
	var countCar int
	var TimeCars []TimeCar

	fmt.Println("Enter car's count")
	_, err := fmt.Fscan(os.Stdin, &countCar)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Enter arrival && departure:")

	var t string
	for i := 0; i < countCar; i++ {
		_, err := fmt.Fscan(os.Stdin, &t)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		times := strings.Split(t, ",")

		// if you want to use with date substitute 2006-01-02T15:04:05 to layout
		tArrival, _ := time.Parse("15:04", times[0])
		tDeparture, _ := time.Parse("15:04", times[1])

		arrival := TimeCar{tArrival,true}
		TimeCars = append(TimeCars, arrival)

		departure := TimeCar{tDeparture,false}
		TimeCars = append(TimeCars, departure)
	}

	sort.SliceStable(TimeCars, func(i, j int) bool {
		return TimeCars[i].timeA.Before(TimeCars[j].timeA)// not including scope, example: 13:30 and 13:30
	})

	var max, count int

	for _, timeCar := range TimeCars {
		if timeCar.isArrival {
			count++
		} else {
			count--
		}
		if max < count {
			max = count
		}
	}
	fmt.Println(max)
}