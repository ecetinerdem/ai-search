package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var m Maze

	var maze, searchType string

	flag.StringVar(&maze, "file", "maze.txt", "maze file")
	flag.StringVar(&searchType, "search", "dfs", "search type")
	flag.Parse()

	err := m.Load(maze)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("maze height/width:", m.Height, m.Width)
}

// Set IDs for each type
const (
	DFS      = iota //Depth First Search
	BFS             //Breath First Search
	GBFS            //Greedy Best First Search
	ASTAR           //A* Search
	DIJKSTRA        //Dijkstra's algorithm
)

// Point struct is for xy coordinates
type Point struct {
	Row int
	Col int
}

// Wall is potential node to be area that cannot be explored
type Wall struct {
	State Point
	wall  bool
}

type Node struct {
	index  int
	State  Point
	Parent *Node
	Action string
}

type Solution struct {
	Action []string
	Cells  []Point
}

// Maze is our game . It keeps every information for the game
type Maze struct {
	Height      int      // How tall the maze
	Width       int      // How wide the maze
	Start       Point    // Starting point to the maze
	Goal        Point    // Ending point to the maze
	WallS       [][]Wall // Walls that cannot be explored
	CurrentNode *Node
	Solution    Solution
	Explored    []Point
	Steps       int
	NumExplored int
	Debug       bool
	SearchType  int
}

func (g *Maze) Load(fileName string) error {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("error opening %s: %s\n", fileName, err)
	}

	defer f.Close()

	var fileContents []string

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.New(fmt.Sprintf("cannot open file %s: %s", fileName, err))
		}
		fileContents = append(fileContents, line)
	}

	foundStart, foundEnd := false, false

	for _, line := range fileContents {
		if strings.Contains(line, "A") {
			foundStart = true
		}

		if strings.Contains(line, "B") {
			foundEnd = true
		}
	}

	if !foundStart {
		return errors.New("starting location not found")
	}

	if !foundEnd {
		return errors.New("ending location not found")
	}

	g.Height = len(fileContents)
	g.Width = len(fileContents[0])

	var rows [][]Wall

	for i, row := range fileContents {
		var cols []Wall
		for j, col := range row {
			currLetter := fmt.Sprintf("%c", col)
			var wall Wall
			switch currLetter {
			case "A":
				g.Start = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false

			case "B":
				g.Goal = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false
			case " ":
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false
			case "#":
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = true
			default:
				continue
			}
			cols = append(cols, wall)
		}
		rows = append(rows, cols)
	}
	g.WallS = rows
	return nil
}
