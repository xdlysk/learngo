package main

import (
	"fmt"
	"strings"
)

func moretypesAll(){
	fmt.Println("=================moretypes begin===================")

	pointers()

	struct1()

	pointToStruct()

	struct2()

	array1()

	arraySlice()

	sliceLiterals()

	sliceDefaults()

	sliceLengthAndCapacity()

	nilSlices()

	makeArray()

	slicesOfSlices()

	fmt.Println("=================moretypes begin===================")
}

func slicesOfSlices(){
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func makeArray(){
	a := make([]int, 5)
	fmt.Printf(" len=%d cap=%d %v\n", len(a), cap(a), a)

	b := make([]int, 0, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	c := b[:2]
	fmt.Printf("len=%d cap=%d %v\n", len(c), cap(c), c)

	d := c[2:5]
	fmt.Printf("len=%d cap=%d %v\n", len(d), cap(d), d)
}

func nilSlices(){
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func sliceLengthAndCapacity(){
	s:=[]int{1,2,3,4,5,6}
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
	// Slice the slice to give it zero length.
	s = s[:0]
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
	// Extend its length.
	s = s[:4]
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
	// Drop its first two values.
	s = s[2:]

	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}

func sliceDefaults(){
	s:=[]int{1,2,3,4,5,6}

	s = s[1:4]

	fmt.Println(s)

	s=s[:2]

	fmt.Println(s)

	s=s[1:]
	fmt.Println(s)
}

func sliceLiterals(){
	q:= []int{1,2,3,4,5,6}
	fmt.Println(q)

	r:=[]bool{true,false,true,false,true,false}
	fmt.Println(r)

	s:=[]struct{
		i int
		b bool
	}{
		{i:2,b:true},
		{3,false},
		{4,true},
		{5,false},
	}

	fmt.Println(s)
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
