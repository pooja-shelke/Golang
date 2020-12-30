# Golang

/*Write a program which accept number from user and print even factors of that
number.
*/

package main

import "fmt"
func EvnFct(iValue int){

	var i int=0
	var iAns int=0

	for i=1;i<=iValue/2;i++{

		iAns=iValue%i

		if((iAns==0) && (i%2==0)){

			fmt.Println(i)
		}
	}



}

func main(){

	var iNum int=0

	fmt.Println("Enter Number")
	fmt.Scanln(&iNum)

	EvnFct(iNum)
}
