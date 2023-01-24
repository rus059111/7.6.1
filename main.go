package main

//Задание 1. Зеркальные билеты

import (
	"fmt"
)

func main() {

	maxNumber := 999999

	count := 0

	for NumberTicket := 100000; NumberTicket < maxNumber; NumberTicket++ {

		a := NumberTicket / 100000
		b := NumberTicket / 10000 % 10
		c := NumberTicket / 1000 % 10
		d := NumberTicket / 100 % 10
		e := NumberTicket / 10 % 10
		f := NumberTicket % 10
		summ1 := a + b + c
		summ2 := d + e + f
		if summ1 == summ2 {

			count++
		} else {
			continue
		}
	}
	fmt.Println(count)

}
