package main

import (
	"fmt"
	"math"
)

var x float64

func main()  {
	fmt.Println("Данная программа предназначена для вычисления произвадной выражения:4x^3*sin(x)")
	fmt.Println("Введите x")
	_, _ = fmt.Scanf("%f", &x)
	decision(x)
}

func derivative(x float64) float64 {
	return 4*math.Pow(x,3)*math.Sin(x)
}

func decision(x float64)  {
	h:=0.001
	res1:=(derivative(x+h)-derivative(x))/h
	h/=2
	res2:=(derivative(x+h)-derivative(x))/h
	if math.Abs(res1-res2)<=0.01 {
		fmt.Println(res1)
	} else {
		fmt.Println("недостаточная точность вычислений")
	}
}