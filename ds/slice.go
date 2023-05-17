package ds

import "fmt"

func RunSliceExmaples() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	slice := numbers[1:4]

	fmt.Println("slice :: ", slice)
	fmt.Println("slice len ::", len(slice))
	fmt.Println("slice cap ::", cap(slice))

	slice2 := numbers[:3]
	fmt.Println("slice2 ::", slice2)

	slice3 := numbers[3:]

	fmt.Println("slice3 ::", slice3)

	// concat example
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := append(s1, s2...)
	fmt.Println("concatinated slice ::", s3)

	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var filtered []int
	var found int
	for _, num := range numbers {
		if num%2 == 0 {
			filtered = append(filtered, num)
		}
	}

	for _, num := range numbers {
		if num == 3 {
			found = num
			break
		}

	}

	fmt.Println("found number ::", found)
	fmt.Println("filtered array ::", filtered)

	lenght := len(numbers)
	//reverse an slice
	for i := 0; i < lenght/2; i++ {
		numbers[i], numbers[lenght-1-i] = numbers[lenght-1-i], numbers[i]
	}

	fmt.Println("reversed slice", numbers)
}
