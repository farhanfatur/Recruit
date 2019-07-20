package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	sourcePath := argument(1)
	if sourcePath == "" {
		log.Fatal("Source path cannot be empty")
	}

	targetPath := argument(2)

	sourceMap := inspectFiles(sourcePath)
	targetMap := inspectFiles(targetPath)

	for sourceEach := range sourceMap {
		if _, ok := targetMap[sourceEach]; ok {
			continue
		} else {
			fmt.Println(sourceEach, "New")
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

func argument(index int) string {
	p := ""
	if len(os.Args) > 1 {
		p = os.Args[index]
	}
	return p
}

func inspectFiles(location string) map[string]bool {
	m := make(map[string]bool)

	filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		pathWithoutBaseArray := strings.Split(path, string(os.PathSeparator))
		pathModified := strings.Join(pathWithoutBaseArray[1:], string(os.PathSeparator))

		m[strings.ToLower(pathModified)] = false
		return nil
	})
	return m
}
