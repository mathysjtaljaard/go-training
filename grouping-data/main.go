package main

import "fmt"

func main() {
	printArray()
	slices()
	loopOverSlice(getSliceInt())
	sliceSlice(getSliceInt())
	appendToSlice(getSliceInt())
	appendSliceToSlice(getSliceInt(), []int{30, 20, 10})
	deleteFromSlice([]int{1, 2, 3, 4, 5, 6})
	makeSlice([]int{9, 8, 7, 6, 5, 4, 3})
	makeMultiSlice()
	makeMap()
	addToMap(map[int]int{1: 1, 2: 2, 3: 3})
	deleteFromMap(map[int]int{1: 1, 2: 2, 3: 3})
}

func printArray() {
	var x [10]string
	x[0] = "A"
	fmt.Println(x)

	xPtr := new([10]string)
	xPtr[1] = "b"
	fmt.Println(xPtr)
	fmt.Println(*xPtr)
	fmt.Println(&xPtr)

}

// SLice allows you to group together values of the same type
func slices() {
	//x := type{value} <- composite literal
	x := []int{4, 5, 6, 7}
	fmt.Println(x)

}

func getSliceInt() []int {
	x := []int{4, 5}
	return x
}

func loopOverSlice(slice []int) {
	fmt.Println(getSliceInt())
	for i, v := range getSliceInt() {
		fmt.Println(i, v)
	}
}

func sliceSlice(sliceToSlice []int) {
	if sliceToSlice != nil && len(sliceToSlice) > 0 {
		for i, _ := range sliceToSlice {
			fmt.Println(sliceToSlice[0 : i+1])
		}
	}
}

func appendToSlice(sliceToAppend []int) {

	for i := 0; i < 10; i++ {
		sliceToAppend = append(sliceToAppend, i+20)
	}
	fmt.Println(sliceToAppend[0:])
}

func appendSliceToSlice(slice1 []int, slice2 []int) {
	// ... aka unfurl
	newSlice := append(slice1, slice2...)
	fmt.Println(newSlice)
}

func deleteFromSlice(slice []int) {
	slice = (slice)[1:4]
	fmt.Println(slice)
}

func makeSlice(slice []int) {
	slice = make([]int, 10, 11)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 1, 2, 3)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func makeMultiSlice() {
	x := make([][][]int, 4)

	for i := 0; i < len(x); i++ {
		x[i] = make([][]int, 4)
		for j := 0; j < len(x[i]); j++ {
			x[i][j] = make([]int, 4)
		}
	}

	fmt.Println(x)
}

func makeMap() {
	mapStringString := map[string]string{
		"one": "1235",
		"two": "4321",
	}

	fmt.Println(mapStringString)
	fmt.Println(mapStringString["one"])
	fmt.Println(mapStringString["two"])

	mapIntString := map[int]string{
		1: "one",
		2: "two",
	}

	fmt.Println(mapIntString)
	fmt.Println(mapIntString[1])
	fmt.Println(mapIntString[1])

	v, ok := mapIntString[3]
	fmt.Println(v, ok)

	if v, ok := mapIntString[4]; ok {
		fmt.Println(v, ok)
	}
	// comma ok idiom
	if v, ok := mapIntString[2]; ok {
		fmt.Println(v, ok)
	}
}

func addToMap(valueMap map[int]int) {
	valueMap[4] = 4
	fmt.Println(valueMap)

	for i, j := range valueMap {
		fmt.Println(i, j)
	}
}

func deleteFromMap(mapValuesToDelete map[int]int) {

	fmt.Println(mapValuesToDelete)

	delete(mapValuesToDelete, 3)

	fmt.Println(mapValuesToDelete)
}
