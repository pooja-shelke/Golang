# Golang

//Write a program which accept number from user and return the count of even digits.

package main

import "fmt"

func CntEvn(iValue int)int{

	var iDigit int=0
	var iCnt int=0

	for iValue!=0{

		iDigit=iValue%10

		if(iDigit%2==0){
			iCnt++
		}
		iValue=iValue/10
	}
	return iCnt
}

func main(){

	var iNum int=0
	var iResult int=0

	fmt.Println("Enter Number")
	fmt.Scanln(&iNum)

	iResult=CntEvn(iNum)

	fmt.Println("Count of even dugit is:",iResult)
}
