//Addition of two numbers

package main                                               

import "fmt"                                         //import package fmt for Println and Scanln

func Addition(iValue1 int,iValue2 int) int {         //Function Defination
	var iResult int=0

	if(iValue1<0 || iValue2<0){                       //If value is negative use updater

		iValue1=-iValue1;
		iValue2=-iValue2;

	}
	iResult=iValue1+iValue2

	return iResult
}

func main() {                                  //Entry point Function

	fmt.Println("Inside main")       

	var iValue1 int=0                              //Local variables
	var iValue2 int=0
	var iResult int=0

	fmt.Println("Enter first Number")               //Accept input from user
	fmt.Scanln(&iValue1)

	fmt.Println("Enter Second Number")
	fmt.Scanln(&iValue2)

	iResult=Addition(iValue1,iValue2)             //Call Addition function

	fmt.Println("Adddition is:",iResult)           //Print output of the Addition Function
}
