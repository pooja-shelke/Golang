# Golang
/*
Problem statement :
 Accept number from user and return the addition of factors of that number.
 
 Input : 12         Output :   16   (1  +  2 +  3 +  4  + 6)
  Input : 10         Output :   8   (1  +  2 +  5)
  Input : 11         Output :   1   (1)
*/

package main

import "fmt"

func FactAdd(iValue int)int{

	var i int=0
	var iAns int=0
	var iSum int=0

	for i=1;i<=iValue/2;i++{

		iAns=iValue%i

		if(iAns==0){
			iSum=iSum+i
		}
	}
	return iSum

}

func main(){

	var iValue int=0
	var iResult int=0

	fmt.Println("Enter Number")
	fmt.Scanln(&iValue)

	iResult=FactAdd(iValue)

	fmt.Println("Addition of a factor is:",iResult)

}
