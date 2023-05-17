package ds

import "fmt"


func RunSliceExmaples(){
	numbers := []int{1,2,3,4,5,6}
	slice :=numbers[1:4]

	fmt.Println("slice :: ",slice)
	fmt.Println("slice len ::", len(slice))
	fmt.Println("slice cap ::", cap(slice))

	slice2 := numbers[:3]
	fmt.Println("slice2 ::", slice2)

	slice3 :=numbers[3:]

	fmt.Println("slice3 ::", slice3)
}