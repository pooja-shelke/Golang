# Golang

//Write a program which accept number from user and display its digits in reverse order.

package main

import "fmt"

func RevNum(iValue int)int{

	
	var iNum int=0
	var iRev int=0

	if(iValue<0){


		iValue=-iValue
	}

	for iValue!=0{

		iNum=iValue%10
		iRev=iRev*10+iNum
		iValue=iValue/10
	   }

	

return iRev
}

func main(){

	var iNum int=0
	var iResult int=0

	fmt.Println("Enter Number")
	fmt.Scanln(&iNum)

	iResult=RevNum(iNum)

	fmt.Println("Reverse number is:",iResult)
}
