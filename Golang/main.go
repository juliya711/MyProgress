//Data Types
package main
import "fmt"
func main(){
	fmt.Println("Hello!")
	var my_test bool=true //default value is false, boolean type
	fmt.Println(my_test)
	var my_number int=10 //integer type
	fmt.Println(my_number)
	var my_float float64=10.5 // float type
	fmt.Println(my_float)
	var my_string string = "Hello!" //string type
	fmt.Println(my_string)
	//string multiple values
	x,y := 10,20
	fmt.Println(x,y)
	// different datatypes & values
	var (
		a=10
		b=20.09
		z="haha"
	)
	fmt.Println(a,b,z)

	//Array- fixed size and can not be changed during run time
	var my_array[3]int //memory allocation
	fmt.Println(my_array)
	var my_array1=[3] int {10,20} // response will be [10 20 0]
	fmt.Println(my_array1)

	//Slice - Dynamically - sized
	var my_slice =[]int{50}
	var other_slice=[]int {10,20,30}
	// fmt.Println(my_slice)
	my_slice = append(my_slice,other_slice...)
	my_slice=append(my_slice,40) // appending another value
	fmt.Println(my_slice)

}


