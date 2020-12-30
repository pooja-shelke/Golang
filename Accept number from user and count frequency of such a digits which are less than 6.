# Golang

//Write a program which accept number from user and count frequency of such a
//digits which are less than 6.

package main

import "fmt"

func CntFrequency(iValue int)int{

	var iDigit int=0
	var iCnt int=0

	for iValue!=0{

	iDigit=iValue%10
	
	if(iDigit<6){
		iCnt++
	}
	iValue=iValue/10
  }
    return iCnt
}

func main(){

	var iValue int=0
	var iResult int=0

	fmt.Println("Enter number")
	fmt.Scanln(&iValue)

	iResult=CntFrequency(iValue)

	fmt.Println("Freuency is :",iResult)
}
