package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
	ex10()
}

func ex1() {

	//composite value
	//x := []int{}
	var arrayInt [5]int //NOT WHAT WAS ASKED
	for i := 0; i < len(arrayInt); i++ {
		arrayInt[i] = i
	}

	fmt.Println(arrayInt)

	for _, i := range arrayInt {
		fmt.Println(i)
	}

	fmt.Printf("%T\n", arrayInt)
	println()
}

func ex2() {
	//Create a Slice using composite values
	arrayInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range arrayInt {
		fmt.Println(i)
	}
	fmt.Printf("%T\n", arrayInt)
	println()

}

func ex3() {
	arrayInt := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(arrayInt[:5])
	fmt.Println(arrayInt[5:])
	fmt.Println(arrayInt[2 : len(arrayInt)-3])
	fmt.Println(arrayInt[1 : len(arrayInt)-4])
	println()
}

func ex4() {
	arrayInt := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	arrayInt = append(arrayInt, 52)
	fmt.Println(arrayInt)

	arrayInt = append(arrayInt, 52, 54, 55)
	fmt.Println(arrayInt)

	x := []int{56, 57, 58, 59, 60}
	arrayInt = append(arrayInt, x...)
	fmt.Println(arrayInt)
	println()
}

func ex5() {
	arrayInt := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	temp2 := append(arrayInt[:3], arrayInt[6:]...)
	fmt.Println(temp2)
	println()
}

func ex6() {
	states := make([]string, 50)
	fmt.Println(states)
	for i := 65; i < 115; i++ {
		// states[i-65] = fmt.Sprintf("%#U", i)
		states[i-65] = string(rune(i))
	}
	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Printf("states[%v]=%v\n", i+1, states[i])
	}
	fmt.Println(states)
	println()
}

func ex7() {
	value := [][]string{{"James", "Bond", "Shaken not Stirred"}, {"Miss", "Moneypenny", "Helloooooo, James"}}
	fmt.Println(value)

	for _, i := range value {
		for _, j := range i {
			fmt.Println(j)
		}
		fmt.Println(i)
	}
	println()
}

func ex8() {
	mapString := map[string][]string{
		`bond_james`:      {"shaken", "martinis", `women`},
		`moneypenny_miss`: {"James", `literature`, `cs`},
		`no_dr`:           {"evil", "ice", "sunsets"},
	}
	fmt.Println(mapString)

	for idx, keys := range mapString {
		fmt.Printf("index %v \n", idx)
		for i := 0; i < len(keys); i++ {
			fmt.Printf("\tindex %v \t == %v\n", i, keys[i])
		}
	}
	println()
}

func ex9() {
	mapString := map[string][]string{
		`bond_james`:      {"shaken", "martinis", `women`},
		`moneypenny_miss`: {"James", `literature`, `cs`},
		`no_dr`:           {"evil", "ice", "sunsets"},
	}

	mapString["ranger_power"] = []string{"monsters", "puns", "kung-fu"}

	for idx, keys := range mapString {
		fmt.Printf("index %v \n", idx)
		for idxInn, val := range keys {
			fmt.Printf("\tindex %v \t == %v\n", idxInn, val)
		}
	}
	println()
}

func ex10() {

	mapString := map[string][]string{
		`bond_james`:      {"shaken", "martinis", `women`},
		`moneypenny_miss`: {"James", `literature`, `cs`},
		`no_dr`:           {"evil", "ice", "sunsets"},
	}

	delete(mapString, "bond_james")
	for idx, keys := range mapString {
		fmt.Printf("index %v \n", idx)
		for idxInn, val := range keys {
			fmt.Printf("\tindex %v \t == %v\n", idxInn, val)
		}
	}
	println()
}
