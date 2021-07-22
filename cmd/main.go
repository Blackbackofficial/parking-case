package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TimeCar struct {
	time int
	isArrival bool
}

func main() {
	var time string
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
		_, err := fmt.Fscan(os.Stdin, &time)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		times := strings.Split(time, ",")

		start, err := strconv.Atoi(times[0])
		if err != nil {
			fmt.Println("Not number")
		}
		arrival := TimeCar{start,true}
		TimeCars = append(TimeCars, arrival)

		end, err := strconv.Atoi(times[1])
		if err != nil {
			fmt.Println("Not number")
		}

		car := TimeCar{end,false}
		TimeCars = append(TimeCars, car)
		fmt.Println(TimeCars)
	}

	//fmt.Fscan(os.Stdin, &name)
	//fmt.Println(name)
}
