package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("names.txt")
	er(err)

	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	er(err)
	fmt.Println(string(bs))
}

func er(er error) {
	if er != nil {
		fmt.Println(er)
		log.Println(er)
		log.Fatalln(er)
	}
}
