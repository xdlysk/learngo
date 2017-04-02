package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func flowControlAll(){
	fmt.Println("=================flowControl begin===================")

	forloop()

	forloop2()

	forWhileLike()

	fmt.Println(if1(2),if1(-4))

	ifelse(1)

	case1()

	case2()

	case3()

	defer1()
	
	deferStack()
	fmt.Println("=================flowControl begin===================")
}

func forloop(){
	sum := 0
	for i:=0;i<10;i++{
		sum+=i
	}
	fmt.Println(sum)
}

func forloop2(){
	sum:=1
	for ;sum<1000;{
		sum+=sum
	}
	fmt.Println(sum)
}

func forWhileLike(){
	sum:=1
	for sum<1000{
		sum+=sum
	}
	fmt.Println(sum)
}

//never call
func forever(){
	for{

	}
}

func if1(x float64) string{
	if x<0{
		return if1(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func ifelse(x int){
	if x >0{
		fmt.Println(x)
	}else{
		fmt.Println(x)
	}
}

func case1(){
	fmt.Print("Go runs on ")
	switch os:=runtime.GOOS; os{
		case "darwin":
		fmt.Println("OS X.")
		case "linux":
		fmt.Println("Linux.")
		default:
		fmt.Printf("%s\n\r",os)
	}
}

func case2(){
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday{
		case today +0:
			fmt.Println("Today.")
		case today+1:
			fmt.Println("Tomorrow.")
		case today+2:
			fmt.Println("In two days.")
		default:
			fmt.Println("Too far away.")
	}
}

func case3(){
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func defer1(){
	defer fmt.Println("defer")

	fmt.Println("hello")
}

func deferStack(){
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}