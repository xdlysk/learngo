package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)

var gol = "go"
const Pi = 3.14

func basicAll(){
	fmt.Println("=================basic begin===================")
	fmt.Println("My favorite number is",rand.Intn(10))
	fmt.Printf("Now you have %g problems",math.Sqrt(7))
	fmt.Println()
	fmt.Println(math.Pi)
	fmt.Println(add(111,222))
	a,b := swap("first","second")
	fmt.Println(a,b)

	fmt.Println(split(40))

	variables()

	variablesWithInitializers()

	shortVariableDeclarations()

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	typeConversions()

	fmt.Println("=================basic end===================")
}


func typeConversions(){
	var x,y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x,y,z)
}

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func shortVariableDeclarations(){
	var i,j int = 1,2
	k:=3
	c,python,java := true,false,"no!"
	fmt.Println(i,j,k,c,python,java)
}

func variablesWithInitializers(){
	var i,j int = 1,2
	var c,python,java = true,false,"no!"
	fmt.Println(i,j,c,python,java)
}


func variables(){
	var c,python,java bool
	var i int
	fmt.Println(i,c,python,java)
}

func add(x int,y int) int{
	return x+y
}

func addShort(x,y int) int{
	return x+y
}


//tuple
func swap(x,y string)(string,string){
	return y,x
}


func split(sum int)(x,y int){
	x = sum*4/9
	y = sum -x
	return 
}