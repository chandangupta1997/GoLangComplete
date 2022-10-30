package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(" welcome to time management class ")

	presentTime := time.Now()

	fmt.Println(" this is present time ", presentTime)

	//2022-10-14 20:04:21.1729609 +0530 +0530 m=+0.004358201

	formtedDate := presentTime.Format("01-02-2006")
	fmt.Println(" this  is Date  ", formtedDate) //10-14-2022

	formatedTime := presentTime.Format("15:04:05")

	fmt.Println(" this  is time   ", formatedTime) //20:11:36

	formatedDay := presentTime.Format("Monday")

	fmt.Println(" this  is the day    ", formatedDay) //Friday

	// another way of creating a date

	createdDate := time.Date(2020, time.August, 10, 23, 60, 0, 0, time.UTC)
	//    2020-08-11 00:00:00 +0000 UTC

	fmt.Println(" this is the date created from time.date and by providing all arguments ", createdDate)

}
