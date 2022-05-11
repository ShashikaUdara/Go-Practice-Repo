package main

import "fmt"

func main(){
	// its main
	var num1, num2, sum int;

	// checking sum before assigning values
	sum = num1 + num2;
	fmt.Println(sum);

	// checking values after the assignment
	num1 = 12,	num2 = 15;
	sum = num1 + num2;
	fmt.Println(sum);
}