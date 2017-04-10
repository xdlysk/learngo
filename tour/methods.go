package main

import (
	"fmt"
	"math"
)

func methodsAll(){
	fmt.Println("=================methods begin===================")

	v :=VertexM{3,4}
	fmt.Println(v.Abs())

	fmt.Println("=================methods end===================")
}

type VertexM struct{
	X,Y float64
}

//similar with .net extend method
func(v VertexM) Abs() float64{
	return math.Sqrt(v.X * v.X + v.Y*v.Y)
}