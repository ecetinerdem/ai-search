🧭 Maze Solver in Go

A command-line maze solver written in Go, implementing multiple classic search algorithms.
The program reads a maze from a text file, finds a path from start (A) to goal (B), and generates a visual PNG output of the solution. It can also create an animated visualization of the solving process.

🚀 Features

✅ Depth-First Search (DFS)

✅ Breadth-First Search (BFS)

✅ Dijkstra’s Algorithm

✅ Greedy Best-First Search (GBFS)

✅ A* Search

✅ PNG image output of solved maze

✅ Optional animated solving process (APNG)

✅ Debug mode for step-by-step tracing

✅ Water cells support (w) for special terrain

📂 Maze File Format

The maze must be provided as a .txt file.

Symbols:
Symbol	Meaning
A	Start position
B	Goal position
#	Wall (blocked cell)
(space)	Empty path
w	Water cell
Example (maze.txt)
##########
#A     w #
# ###### #
#      # #
# #### #B#
##########
⚙️ Installation

Make sure you have Go installed (Go 1.20+ recommended).

Clone the repository:

git clone https://github.com/yourusername/maze-solver.git
cd maze-solver

Install dependencies:

go mod tidy
▶️ Usage

Run the program using:

go run . -file=maze.txt -search=dfs
Available Flags
Flag	Default	Description
-file	maze.txt	Maze file to load
-search	dfs	Search algorithm
-debug	false	Enable debug output
-animate	false	Generate animation
🔎 Search Algorithms

You can choose between:

dfs – Depth-First Search

bfs – Breadth-First Search

dijkstra – Dijkstra's Algorithm

gbfs – Greedy Best-First Search

astar – A* Search

Example:

go run . -file=maze.txt -search=astar -animate=true
🖼 Output

After solving:

image.png → Final maze with solution path

animation.png → (if -animate=true) animated solving process

tmp/ → Temporary frames for animation

Color Legend
Color	Meaning
🟢 Green	Solution path
🟩 Dark Green	Start
🔴 Red	Goal
🟡 Yellow	Explored nodes
⚫ Black	Walls
🔵 Blue	Water cells
🟠 Orange	Current node
🧠 How It Works

The maze file is parsed into a 2D grid.

A search algorithm explores neighboring cells.

The solution path is reconstructed using parent pointers.

The maze is rendered as an image using Go’s image package.

Optional animation is generated using APNG.

📦 Dependencies

This project uses:

golang.org/x/image

github.com/StephaneBunel/bresenham

github.com/kmicki/apng

They will be automatically installed with:

go mod tidy
🛠 Project Structure
.
├── main.go
├── maze.go
├── search_*.go
├── image.go
├── utils.go
├── maze.txt
├── tmp/
└── README.md
📈 Example Output

After running:

Solution is 14 steps
Time to solve: 2.3ms
Explored 47 nodes

An image like this will be generated:

Solution path marked in green

Explored nodes in yellow

Grid layout preserved
