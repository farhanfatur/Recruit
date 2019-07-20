package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

func ReadDir() []string {
	dirname := "."

	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	var data []string
	for _, file := range files {

		fmt.Println(file.Name())
		data = append(data, file.Name())
	}
	return data
}

func main() {
	var d1 = ReadDir()
	fmt.Println(d1)
	fmt.Scanln()
	time.Sleep(time.Second * 5)
	var d2 = ReadDir()
	if !reflect.DeepEqual(d1, d2) {
		fmt.Println("Modified")
	} else {
		fmt.Println(d1)
	}
}
