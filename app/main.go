package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type person []struct {
	First string   `json:"firstName"`
	Last  string   `json:"lastName"`
	Age   int      `json:"age"`
	Foods []string `json:"favorite foods"`
}

func main() {
	p := person{
		{
			First: "welcome",
			Last:  "brad",
			Age:   21,
			Foods: []string{"cake", "biltong", "pizza"},
		},
		{
			First: "crap",
			Last:  "shoot",
			Age:   44,
			Foods: []string{"poo", "nuts", "brownie"},
		},
	}
	jValue := p.marshal()
	fmt.Println("marshalled", jValue)

	var people person
	unmarshal([]byte(jValue), &people)
	fmt.Printf("unmarshalled %+v\n", people)

	startingCode()

	customSort()

	passwordGen("bank1234")
}

func (p *person) marshal() string {
	s, err := json.MarshalIndent(p, " ", "    ")
	if err != nil {
		return err.Error()
	}

	return string(s)
}

func unmarshal(value []byte, p *person) {
	err := json.Unmarshal(value, &p)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func startingCode() {
	sliceInt := []int{4, 5, 7, 22, 3, 1, 65, 99, 80, 1000}
	sliceString := []string{"Jabba", "Ann", "Burt", "Alphalpa", "Wonderboy"}

	fmt.Printf("%T\n", sliceInt)
	fmt.Printf("%T\n", sliceString)
	fmt.Println(sliceInt, sliceString)

	sort.Ints(sliceInt)
	sort.Strings(sliceString)

	fmt.Println(sliceInt, sliceString)

}

type personSort struct {
	first string
	age   int
}

type SortByName []personSort
type SortByAge []personSort

func (a personSort) sum() int {
	sum := 0
	for i := 0; i < len(a.first); i++ {
		sum += int(a.first[i])
	}
	return sum
}

func (a SortByName) Len() int      { return len(a) }
func (a SortByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByName) Less(i, j int) bool {
	return a[i].first < a[j].first
}

func (a SortByAge) Len() int      { return len(a) }
func (a SortByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func customSort() {
	p1 := personSort{"hames", 32}
	p2 := personSort{"blbpenny", 21}
	p3 := personSort{"No", 12}
	p4 := personSort{"crow", 52}
	p5 := personSort{"crop", 52}

	people := []personSort{
		p1, p2, p3, p4, p5,
	}

	fmt.Println(people)

	sort.Stable(SortByName(people))
	fmt.Println(people)

	fmt.Println(p1.sum())

	fmt.Println(p1.first < p2.first)
}

func passwordGen(p string) {
	hs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("P =", p, "EncP = ", string(hs))
}
