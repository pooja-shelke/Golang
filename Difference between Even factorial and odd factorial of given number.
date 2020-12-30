# Golang

/* Write a program which returns difference between Even factorial and odd factorial
of given number.

Input : 5
Output : -7 (8 - 15)

Input : -5
Output : -7 (8 - 15)

Input : 10
Output : 2895 (3840 - 945)
*/

package main

import "fmt"

func DifferenceEvnOdd(iValue int)int{

	var i int=0
	var iEvn int=1
	var iOdd int=1

	if(iValue<0){
		iValue=-iValue
	}
	
	for i=1;i<=iValue;i++{

		if(i%2==0){

			iEvn=iEvn*i
		}else{

			iOdd=iOdd*i
		}
	}

	return iEvn-iOdd
}

func main(){
	var iValue int=0
	var iResult int=0

	fmt.Println("Enter number")
	fmt.Scanln(&iValue)

	iResult=DifferenceEvnOdd(iValue)

	fmt.Println("Difference between Even and odd is:",iResult)
}
