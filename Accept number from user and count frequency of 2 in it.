# Golang

//Write a program which accept number from user and count frequency of 2 in it.

package main
import "fmt"

func CntFreq(iValue int)int{

	var iCnt int=0
	var iDigit int=0

	for iValue!=0{

		iDigit=iValue%10
		if(iDigit==2){

		iCnt++
	}
		iValue=iValue/10
	}
	return iCnt
}

 func main(){

 	var iNum int=0
 	var iResult int=0

 	fmt.Println("Enter number")
 	fmt.Scanln(&iNum)

 	iResult=CntFreq(iNum)

 	fmt.Println("Frequency is:",iResult)

 }
