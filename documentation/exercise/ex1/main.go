package main 

import (
	"main/documentation/exercise/ex1/dog"
	"fmt"
	"log"
)

func main()  {
	ha, err := dog.Years(3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ha)

	ha, err = dog.Years(0)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ha)
}