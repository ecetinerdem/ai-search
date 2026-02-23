package main

import (
	"log"
	"math"
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

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func euclideanDist(p, goal Point) float64 {
	return math.Sqrt(float64(p.Row-goal.Row)*float64(p.Row-goal.Row) + float64(p.Col-goal.Col)*float64(p.Col-goal.Col))
}
