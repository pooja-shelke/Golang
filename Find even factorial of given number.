# Golang
//find even factorial of given number.

package main

import "fmt"

func FactorialEvn(iValue int)int{

	var i int=0
	var iMult int=1

	if(iValue<0){
		iValue=-iValue
	}

	for i=1;i<=iValue;i++{

		if(i%2==0){
			iMult=iMult*i
		}
	}
	return iMult
}

func main(){

	var iValue int=0
	var iResult int=0

	fmt.Println("Enter number")
	fmt.Scanln(&iValue)

	iResult=FactorialEvn(iValue)

	fmt.Println("Factorial of a number is:",iResult)
}
