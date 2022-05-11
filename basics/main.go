package main

import "fmt"

func main(){
	// its main
	var num1, num2, sum int;
	var fl1, fl2, flSum float32;
	var character, level byte;

	// you can use direct declaration, the compilar will identify the correct data type itself.
	number1 = 4.25;
	number2 = 45.25;
	sumNum = 0;

	// checking sum before assigning values
	sum = num1 + num2;
	fmt.Println(sum);

	// checking values after the assignment
	// integers
	num1 = 12; num2 = 15;
	sum = num1 + num2;
	fmt.Println(sum);

	// floats
	fl1 = 4.12; 
	fl2 = 45.89; 
	flSum = fl1 + fl2;
	fmt.Println(flSum);


	// bytes
	character = 180;
	level = character - 10;
	fmt.Println(level);

	sumNum = number1 + number2
	fmt.Println(sumNum);
}