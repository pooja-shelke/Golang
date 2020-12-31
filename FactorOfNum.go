# Golang

/*
Problem statement :
 Accept number from user and display factors of that number.
 
 Input : 12         Output :   1    2   3   4   6
*/
package main

import "fmt"

func Factors(iValue int){

	var i int=0
	var iAns int=0

	for i=1;i<iValue;i++{

		iAns=iValue%i

		if(iAns==0){
			fmt.Println(i)
		}
	}
}


func main(){
	var iValue int=0

	fmt.Println("Enter number")
	fmt.Scanln(&iValue)

	Factors(iValue)
}
