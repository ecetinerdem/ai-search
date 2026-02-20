package main

import (
	"log"
	"os"
)

func inExplored(needle Point, haystack []Point) bool {

	for _, x := range haystack {
		if needle.Row == x.Row && needle.Col == x.Col {
			return true
		}
	}
	return false
}

func emptyTmp() {
	directory := "./tmp/"

	dir, err := os.Open(directory)
	if err != nil {
		log.Println(err)
	}

	filesToDelete, err := dir.ReadDir(0)

	if err != nil {
		log.Println(err)
	}
	for index := range filesToDelete {
		f := filesToDelete[index]
		fullPath := directory + f.Name()

		err := os.Remove(fullPath)
		if err != nil {
			log.Println(err)
		}
	}
}
