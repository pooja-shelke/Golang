# Golang

//Write a program to find factorial of given number.

package main

import "fmt"

func Factorial(iValue int)int{

	var iMult int=1
	var i int

	if(iValue<0){

		iValue=-iValue
	}

	for i=iValue;i>=1;i--{

		iMult=iMult*i
	}

	return iMult
}

func main(){

	var iValue int=0
	var iResult int=0

	fmt.Println("Enter Number")
	fmt.Scanln(&iValue)

	iResult=Factorial(iValue)

	fmt.Println("Factorial of a number is:",iResult)
}
