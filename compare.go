package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
		var take, err = exec.Command("sh", "-c", "sourcetarget").Output()
		if err != nil {
			fmt.Println("Modified")
			fmt.Printf("%s", err)
		}
		fmt.Println(take)
	} else {
		fmt.Println(d1)
		fmt.Println("Nothing Happend")
	}
}
