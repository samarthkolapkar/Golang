//Write a Go program to find the sum of all elements in an integer array.
package main

import "fmt"

func main(){
	var n int
	fmt.Println("Enter the size of array:")
	fmt.Scan(&n)
	arr:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	var total int 
	for i:=0;i<n;i++{
		total=total+arr[i]
	}
	fmt.Printf("Sum of element in array is:%d\n",total)
}