package main

import (
	"fmt"
	"time"
)

// go> time package

func main() {
	fmt.Println("my time file ")

	presentTime := time.Now()
	presentNanoTime := time.Now().Nanosecond()

	fmt.Println("present time: ", presentTime)
	fmt.Println("present nano time: ", presentNanoTime)

	fmt.Println("", presentTime.Format("01-02-2006  15:04:05 Monday"))

	createdDate := time.Date(2020, time.April, 04, 11, 54, 0, 0, time.UTC)
	fmt.Println("Created date :", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

// go env in terminal
// GOOS =WINDOWS
// CHANGE IT if you want to get linux executable file GOOS ="linux" go build

// Memory Management
// Memory allocation and Deallocation happens Automatically.

// ----- 2 methods:-----
// 1. new():
//  allocate memory but no INIT
// 	you will get memory address
//  zeroed storage

// 2. make():

// allocate memory and INIT
// you will get memory address
// non- zeroed storage

// GC happens automatically
// out of scope or nil
//  https://pkg.go.dev/runtime
