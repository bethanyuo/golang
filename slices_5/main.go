package main

import "fmt"

func main() {

	slice1 := []string{
		"Hello",
		"Golang",
		"World",
	}

	for index, value := range slice1 {
		fmt.Printf("%d : %q ", index, value)
	}
	fmt.Println() // 0 : "Hello" 1 : "Golang" 2 : "World"

	// Can only append to SLICES (dynamic)

	mySlice := []float64{1.2, 5.6} // TYPE == SLICE [2]float64{}

	mySlice[0] = float64(6)

	a := 10.
	mySlice[0] = a

	mySlice[1] = 10.10

	// mySlice = append(a)	ERROR: first argument to append must be slice; have float64

	mySlice = append(mySlice, a)

	fmt.Println(mySlice) // [10 10.1 10]

	nums := []float64{1.1, 2.2, 3.3}

	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println(nums) // [1.1 2.2 3.3 4.1 5.5 6.6]

	n := []float64{4.7, 8.8}

	nums = append(nums, n...) // ===> Append n to nums

	fmt.Println(nums) // [1.1 2.2 3.3 4.1 5.5 6.6 4.7 8.8]

	///// Iterate over the slice, ignoring the first and the last two elements && Print SUM \\\\\

	nums1 := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	sum := 0
	for i := 1; i < len(nums1)-2; i++ { // FOR LOOP without SLICE EXPRESSION
		fmt.Printf("%d ", nums1[i])
		sum += nums1[i]
	}
	fmt.Println()
	fmt.Printf("Sum: %d\n", sum)

	nums2 := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum2 := 0
	for _, v := range nums2[1 : len(nums2)-2] { // Using SLICE EXPRESSION ==> nums2[1 : lens(nums2)-2]
		fmt.Println(v)
		sum2 += v
	}
	fmt.Println("Sum:", sum2)

	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := make([]string, len(friends)) // making a NEW slice yourFriends with LEN of friends slice
	copy(yourFriends, friends)                  // copying the values of old slice to the new slice ==> [Marry John Paul Diana]

	yourFriends[0] = "Dan" // Changing element value ==> slices are not related

	fmt.Println(friends, yourFriends) // [Marry John Paul Diana] [Dan John Paul Diana]

	//myFriends := make([]string, len(friends)) // making a NEW slice yourFriends with LEN of friends slice
	myFriends := []string{}
	myFriends = append(myFriends, friends...)

	myFriends[0] = "Rosa"

	fmt.Println(friends, myFriends) // [Marry John Paul Diana] [Rosa John Paul Diana]

	///// Slice called newYears that contains the first 3 and the last 3 elements of the slice \\\\\

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}

	// newYears = append(newYears, years[:3], years[len(years)-3:]...)	// ERROR: too many arguments to append

	newYears = append(years[:3], years[len(years)-3:]...)

	fmt.Println(newYears)         // [2000 2001 2002 2008 2009 2010]
	fmt.Printf("%#v\n", newYears) // []int{2000, 2001, 2002, 2008, 2009, 2010}
}
