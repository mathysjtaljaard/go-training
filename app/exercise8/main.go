package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Name       string `json:"UserName"`
	Occupation string `json:"JobType"`
	password   []byte
	Age        int `json:"Age"`
}

func main() {
	users := []user{
		{Name: "jim", Occupation: "strangerDanger", password: encrypt("12345"), Age: 22},
		{Name: "ali", Occupation: "boxer", password: encrypt("45678"), Age: 34},
		{Name: "cat", Occupation: "latenightEntertainer", password: encrypt("33333"), Age: 19},
	}

	problem1(users)
}

type SortByAge []user

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []user

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func (p user) Println() {

	fmt.Printf("UserName: %s\nOccupation: %s\nAge: %v\n", p.Name, p.Occupation, p.Age)
}

func problem1(users []user) {
	fmt.Println()
	fmt.Println("Problem 1")
	value, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	for _, u := range users {
		u.Println()
	}
	fmt.Println(string(value))
	problem2(value)
	problem3(users)
	problem4(users)
	problem5(users)
}

func problem2(valueToUnmarshal []byte) {
	fmt.Println()
	fmt.Println("Problem 2")
	users := []user{}
	json.Unmarshal(valueToUnmarshal, &users)

	for _, u := range users {
		u.Println()
	}
}

func problem3(users []user) {
	fmt.Println()
	fmt.Println("Problem 3")
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("error for New Encoder is", err)
		return
	}

	// fmt.Println("Problem 3 Decode")
	// users := json.NewDecoder(strings.NewReader(string(decode)))
	// fmt.Println(users)
	// formatPrintString(users)
}

func problem4(users []user) {
	fmt.Println()
	fmt.Println("Problem 4")
	sort.Sort(SortByAge(users))
	fmt.Println(users)
}

func problem5(users []user) {
	fmt.Println()
	fmt.Println("Problem 4")
	sort.Sort(ByName(users))
	fmt.Println(users)
}

func encrypt(pass string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return []byte(fmt.Sprint(pass, "couldNotHash"))
	}
	return hash
}

func formatPrintString(values ...interface{}) {
	fmt.Printf("\n%v\n", values)
}
