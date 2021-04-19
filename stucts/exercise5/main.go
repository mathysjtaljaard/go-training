package main

import "fmt"

func main() {
	num1()
	num2()
	num3()
	num4()
}

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

type vehicle struct {
	numDoors int
	color    string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

var person1 person = person{
	firstName:       "Jon",
	lastName:        "Snow",
	iceCreamFlavors: []string{"mint", "chocolate"},
}

var person2 person = person{
	firstName:       "Jenna",
	lastName:        "Jameson",
	iceCreamFlavors: []string{"salty", "balls"},
}

func num1() {

	persons := []person{person1, person2}

	for _, person := range persons {
		printPerson(person)
	}
}

func num2() {

	personsMap := map[string]person{
		person1.lastName: person1,
		person2.lastName: person2,
	}

	for k, v := range personsMap {
		fmt.Printf("Key (Last Name): %v \n", k)
		printPerson(v)
	}

}

func num3() {

	vehTruck := truck{
		vehicle: vehicle{
			numDoors: 2,
			color:    "red",
		},
		fourWheel: false,
	}

	vehSedan := sedan{
		vehicle: vehicle{
			numDoors: 4,
			color:    "blue",
		},
		luxury: true,
	}

	fmt.Printf(`
	Truck Specs:
		color: %v
		number of doors: %v,
		four wheel drive: %v
`,
		vehTruck.color, vehTruck.numDoors, vehTruck.fourWheel)

	fmt.Printf(`
Sedan Specs:
	color: %v
	number of doors: %v,
	four wheel drive: %v
`,
		vehSedan.color, vehSedan.numDoors, vehSedan.luxury)
}

func num4() {

	whatAmI := struct {
		door    int
		color   string
		isMango bool
	}{
		door:    2,
		color:   "blue",
		isMango: false,
	}

	fmt.Printf(`
	What am I?
		doors? %v
		color? %v
		is a mango? %v
	`, whatAmI.door, whatAmI.color, whatAmI.isMango)
}

func printPerson(person person) {
	flavors := ""
	for _, iceCream := range person.iceCreamFlavors {
		flavors += fmt.Sprintf("\t- %v\n", iceCream)
	}
	fmt.Printf(`
First name: %v
Last name: %v
Favorite iceCream: 
%v
			`, person.firstName, person.lastName, flavors)
}
