//Nicholas Larsen
//March 5,2020
//Finds the min and max numbers in an array
package main

import "fmt"

func main() {

	var a = [10]int{22,32,1,65,8,90,2,11,19,7}
  //creates the array
	min, max := findMinAndMax(a)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
  //prints out the min and the max

}

func findMinAndMax(a [10]int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
      //uses range to find the min number and the max number
		}
	}
	return min, max
  //returns them so they can be used in the func main
}