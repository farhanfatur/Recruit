package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func argument2(index int) string {
	p := ""
	if len(os.Args) > 1 {
		p = os.Args[index]
	}
	return p
}

func inspectFiles2(location string) map[string]string {
	m := make(map[string]string)

	filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		pathWithoutBaseArray := strings.Split(path, string(os.PathSeparator))
		pathModified := strings.Join(pathWithoutBaseArray[1:], string(os.PathSeparator))

		buf, _ := ioutil.ReadFile(path)
		m[strings.ToLower(pathModified)] = string(buf)
		return nil
	})

	return m
}

func main() {
	sourcePath := argument2(1)
	if sourcePath == "" {
		log.Fatal("Source path cannot be empty")
	}

	targetPath := argument2(2)
	if targetPath == "" {
		log.Fatal("Target path cannot be empty")
	}

	sourceMap := inspectFiles2(sourcePath)
	targetMap := inspectFiles2(targetPath)

	for sourceEach, sourceVal := range sourceMap {
		if targetVal, ok := targetMap[sourceEach]; ok {
			if sourceVal == targetVal {
				continue
			} else {
				fmt.Println(sourceEach, "MODIFIED")
			}
		} else {
			fmt.Println(sourceEach, "NEW")
		}
	}

	for targetEach := range targetMap {
		if _, ok := sourceMap[targetEach]; ok {
			continue
		} else {
			fmt.Println(targetEach, "DELETED")
		}
	}
}
