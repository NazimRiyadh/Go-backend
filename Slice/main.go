//Slice works underneath with arrays

package main

// struct Slice {
//     pointer *T // points underlying array
//     length  int // current elements number
//     capacity int // maximum elements (until reallocation)
//     }

//arrays have fixed size
//Slice size is dyanamic

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	//creating a slice from an array
	slc := arr[2:4]
	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))

	//creating slice from a slice
	slc2 := slc[0:3]
	fmt.Println(slc2)
	fmt.Println(len(slc2))
	fmt.Println(cap(slc2))

	//slice literal-> creating a slice without an array
	yoyo := []int{1, 2, 3, 4, 5}
	fmt.Println("yoyo", yoyo)

	//creating slice using make()
	badshah := make([]int, 3)
	fmt.Println("badshah", badshah)
	//length 3 capacity 5
	eminem := make([]int, 3, 5)
	fmt.Println("eminem", eminem)

	//nil slice
	var nilSlice []int
	fmt.Println("nilslice", nilSlice, len(nilSlice), cap(nilSlice))

	nilSlice = append(nilSlice, 1)
	fmt.Println("nilslice after append", nilSlice, len(nilSlice), cap(nilSlice))

}
