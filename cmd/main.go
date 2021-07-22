package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TimeCar struct {
	time int
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

	for i := 0; i < countCar; i++ {
		var time string
		_, err := fmt.Fscan(os.Stdin, &time)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		times := strings.Split(time, ",")

		start, err := strconv.Atoi(times[0])
		if err != nil {
			fmt.Println(err)
		}
		arrival := TimeCar{start,true}
		TimeCars = append(TimeCars, arrival)

		end, err := strconv.Atoi(times[1])
		if err != nil {
			fmt.Println(err)
		}
		departure := TimeCar{end,false}
		TimeCars = append(TimeCars, departure)
	}

	fmt.Println(TimeCars)

	sort.SliceStable(TimeCars, func(i, j int) bool {
		return TimeCars[i].time < TimeCars[j].time
	})

	fmt.Println(TimeCars)

	var max, count int

	for _, time := range TimeCars{
		if time.isArrival {
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
