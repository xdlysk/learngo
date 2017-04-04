package main

import (
	"fmt"
)

func moretypesAll(){
	fmt.Println("=================moretypes begin===================")

	pointers()

	struct1()

	pointToStruct()

	struct2()

	array1()

	arraySlice()

	fmt.Println("=================moretypes begin===================")
}

func arraySlice(){
	primes := [6]int{1,2,3,4,5,6}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func array1(){
	var a [2]string
	a[0] = "hello"
	a[1] = "go"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes := [6]int{1,2,3,4,5,6}
	fmt.Println(primes)
}

func pointers(){
	i,j := 42,2701;

	p:= &i //point to i 
	fmt.Println(*p) //read i through the pointer

	*p = 21 //set i through the pointer
	fmt.Println(i) //print i

	p = &j //set p point to j
	*p = *p /37 
	fmt.Println(j)

}


func struct1(){
	s := Vertex{1,2}
	fmt.Println(s.X,s.Y)
}

func pointToStruct(){
	v := Vertex{1,2}
	p:= &v
	p.X = 1e9
	fmt.Println(v)
}

func struct2(){
	var (
		v1 = Vertex{1,2}
		v2 = Vertex{X : 1}
		v3 = Vertex{}
		p = &Vertex{1,2}
	)

	fmt.Println(v1,v2,v3,p)
}

//Vertex some struct
type Vertex struct{
	X int
	Y int
}
