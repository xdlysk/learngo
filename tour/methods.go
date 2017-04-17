package main

import (
	"fmt"
	"math"
)

func methodsAll(){
	fmt.Println("=================methods begin===================")

	v :=VertexM{3,4}
	fmt.Println(v.Abs())

	v.Scale(10);
	fmt.Println(v.Abs())

	f:= MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	var a Abser
	a = f
	a = &v
	//a = v

	fmt.Println(a.Abs())

	fmt.Println("=================methods end===================")
}

type VertexM struct{
	X,Y float64
}

//similar with .net extend method
func(v VertexM) Abs() float64{
	return math.Sqrt(v.X * v.X + v.Y*v.Y)
}

func(v *VertexM) Scale(f float64){
	v.X = v.X*f
	v.Y = v.Y*f
}

type MyFloat float64

func (f MyFloat) Abs() float64{
	if f<0{
		return float64(-f)
	}
	return float64(f)
}

type Abser interface{
	Abs() float64
}